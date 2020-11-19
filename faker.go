// Package faker is a random data generator and struct fake data generator.
package faker

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	tagName   = "faker"
	skipTag   = "-"
	uniqueTag = "unique"
)

var (
	tagFnCallRegexp = regexp.MustCompile(`(.+?)\((.+?)\)`)
	tagLenRegexp    = regexp.MustCompile(`len=(\d+)`)
	tagSkipIfRegexp = regexp.MustCompile(`skip\s+if\s+(\w+)`)
)

type fakerTag struct {
	funcName    string
	skipIf      string
	uniqueGroup string
	length      int
	params      []string
}

func (tag *fakerTag) mustSkip() bool {
	return tag.funcName == skipTag
}

// func decodeTag(tagString string) *fakerTag {
func decodeTag(structReflectType reflect.Type, fieldIndex int) *fakerTag {
	fieldReflectType := structReflectType.Field(fieldIndex)
	tagString := fieldReflectType.Tag.Get(tagName)
	tag := &fakerTag{}
	for _, token := range strings.Split(tagString, ";") {
		if token == skipTag {
			tag.funcName = skipTag
			return tag
		}
		if m := tagSkipIfRegexp.FindStringSubmatch(token); len(m) == 2 {
			tag.skipIf = m[1]
			continue
		}
		if token == uniqueTag {
			tag.uniqueGroup = fmt.Sprintf("%s-%s", structReflectType.Name(), fieldReflectType.Name)
			continue
		}
		if m := tagLenRegexp.FindStringSubmatch(token); len(m) == 2 {
			tag.length, _ = strconv.Atoi(m[1])
			continue
		}
		if tag.funcName == "" {
			if m := tagFnCallRegexp.FindStringSubmatch(token); len(m) == 3 {
				tag.funcName = m[1]
				tag.params = strings.Split(m[2], ",")
				continue
			}
			tag.funcName = token
		}
	}
	return tag
}

// Build fills in exported elements of a struct with random data based on the
// value of `faker` tag of exported elements. The faker tag value can be any
// available function (case insensitive).
//
// Use `faker:"-"` to explicitly skip an element.
//
// Use `faker:"skip if FieldName"` to explicitly skip this field if another
// field (FieldName) is not empty.
//
// Use `faker:"unique"` to guarantee a unique value.
//
// Use `faker:"len=x"` to specify the length of a slice or the size of a map
// (if ommitted, will be generated a slice or map with random size between 1
// and 8).
//
// Built-in supported types are: bool, int, int8, int16, int32, int64, uint,
// uint8, uint16, uint32, uint64, float32, float64, string. Other standard
// library supported types are time.Time and time.Duration. But is really easy
// to extend faker to add other builders to support other types and or
// customize faker's behavior (see RegisterBuilder function).
func Build(input interface{}) error {
	inputReflectType := reflect.TypeOf(input)
	if inputReflectType == nil {
		return errors.New("faker.Build input interface{} not allowed")
	}
	if inputReflectType.Kind() != reflect.Ptr {
		return errors.New("faker.Build input is not a pointer")
	}

	parentReflectTypes := make(map[reflect.Type]struct{})
	inputReflectValue := reflect.ValueOf(input)
	err := build(inputReflectValue, parentReflectTypes, &fakerTag{})
	if err != nil {
		return err
	}
	return nil
}

func build(inputReflectValue reflect.Value, parentReflectTypes map[reflect.Type]struct{}, tag *fakerTag) error {
	inputReflectType := inputReflectValue.Type()
	kind := inputReflectType.Kind()

	var (
		fn    builderFunc
		found bool
		key   string
	)

	key = builderKey(tag.funcName, inputReflectType.String())
	fn, found = builders[key]

	if found {
		if !inputReflectValue.IsZero() {
			return nil
		}
		var (
			value interface{}
			err   error
		)
		if tag.uniqueGroup != "" {
			value, err = Uniq(tag.uniqueGroup, 0, func() (interface{}, error) {
				return fn(tag.params...)
			})
		} else {
			value, err = fn(tag.params...)
		}
		if err != nil {
			return err
		}
		inputReflectValue.Set(reflect.ValueOf(value))
		return nil
	}

	switch kind {
	case reflect.Ptr:
		return buildPtr(inputReflectValue, inputReflectType, parentReflectTypes, tag)
	case reflect.Struct:
		return buildStruct(inputReflectValue, inputReflectType, parentReflectTypes, tag)
	case reflect.Slice:
		return buildSlice(inputReflectValue, inputReflectType, parentReflectTypes, tag)
	case reflect.Map:
		return buildMap(inputReflectValue, inputReflectType, parentReflectTypes, tag)
	default:
		if tag.funcName != "" {
			return fmt.Errorf("invalid faker function '%s' for type '%s'", tag.funcName, inputReflectType.String())
		}
		return nil
	}
}

func buildPtr(inputReflectValue reflect.Value, inputReflectType reflect.Type, parentReflectTypes map[reflect.Type]struct{}, tag *fakerTag) error {
	if detectInfiniteLoopRecursion(parentReflectTypes, inputReflectType) {
		return nil
	}
	if inputReflectValue.IsNil() {
		newVar := reflect.New(inputReflectType.Elem())
		inputReflectValue.Set(newVar)
	}
	return build(inputReflectValue.Elem(), parentReflectTypes, tag)
}

func buildStruct(inputReflectValue reflect.Value, inputReflectType reflect.Type, parentReflectTypes map[reflect.Type]struct{}, tag *fakerTag) error {
	if detectInfiniteLoopRecursion(parentReflectTypes, inputReflectType) {
		return nil
	}
	parentReflectTypesCopy := make(map[reflect.Type]struct{})
	for k, v := range parentReflectTypes {
		parentReflectTypesCopy[k] = v
	}
	parentReflectTypesCopy[inputReflectType] = struct{}{}
	for i := 0; i < inputReflectValue.NumField(); i++ {
		fieldTag := decodeTag(inputReflectType, i)
		if fieldTag.mustSkip() {
			continue
		}
		if fieldTag.skipIf != "" {
			skipIfReflectValue := inputReflectValue.FieldByName(fieldTag.skipIf)
			if (skipIfReflectValue != reflect.Value{}) && !skipIfReflectValue.IsZero() {
				continue
			}
		}
		if !inputReflectValue.Field(i).CanSet() {
			continue // to avoid panic to set on unexported field in struct
		}
		err := build(inputReflectValue.Field(i), parentReflectTypesCopy, fieldTag)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildSlice(inputReflectValue reflect.Value, inputReflectType reflect.Type, parentReflectTypes map[reflect.Type]struct{}, tag *fakerTag) error {
	if detectInfiniteLoopRecursion(parentReflectTypes, inputReflectType.Elem()) {
		return nil
	}
	if inputReflectValue.IsNil() {
		var sliceLen int
		if tag != nil && tag.length != 0 {
			sliceLen = tag.length
		} else {
			sliceLen = IntInRange(1, 8)
		}
		newSlice := reflect.MakeSlice(inputReflectType, sliceLen, sliceLen)
		inputReflectValue.Set(newSlice)
		return buildSliceElems(inputReflectValue, parentReflectTypes, tag)
	}
	if len(parentReflectTypes) == 0 {
		return buildSliceElems(inputReflectValue, parentReflectTypes, tag)
	}
	return nil
}

func buildMap(inputReflectValue reflect.Value, inputReflectType reflect.Type, parentReflectTypes map[reflect.Type]struct{}, tag *fakerTag) error {
	keyReflectType := inputReflectType.Key()
	if detectInfiniteLoopRecursion(parentReflectTypes, keyReflectType) {
		return nil
	}
	elemReflectType := inputReflectType.Elem()
	if detectInfiniteLoopRecursion(parentReflectTypes, elemReflectType) {
		return nil
	}
	if inputReflectValue.IsNil() {
		var (
			mapLen int
			key    reflect.Value
			elem   reflect.Value
			err    error
		)
		if tag != nil && tag.length != 0 {
			mapLen = tag.length
		} else {
			mapLen = IntInRange(1, 8)
		}
		newMap := reflect.MakeMap(inputReflectType)
		for i := 0; i < mapLen; i++ {
			key = reflect.New(keyReflectType).Elem()
			elem = reflect.New(elemReflectType).Elem()
			err = build(key, parentReflectTypes, tag)
			if err != nil {
				return err
			}
			err = build(elem, parentReflectTypes, tag)
			if err != nil {
				return err
			}
			newMap.SetMapIndex(key, elem)
		}
		inputReflectValue.Set(newMap)
	}
	return nil
}

func buildSliceElems(inputReflectValue reflect.Value, parentReflectTypes map[reflect.Type]struct{}, tag *fakerTag) error {
	for i := 0; i < inputReflectValue.Len(); i++ {
		err := build(inputReflectValue.Index(i), parentReflectTypes, tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func detectInfiniteLoopRecursion(parentReflectTypes map[reflect.Type]struct{}, reflectType reflect.Type) bool {
	for {
		if reflectType.Kind() == reflect.Ptr {
			reflectType = reflectType.Elem()
		} else {
			break
		}
	}
	_, found := parentReflectTypes[reflectType]
	return found
}
