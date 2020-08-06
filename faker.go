package faker

import (
	"errors"
	"fmt"
	"reflect"
)

const tagName = "faker"

func Build(input interface{}) error {
	inputReflectType := reflect.TypeOf(input)
	if inputReflectType == nil {
		return errors.New("faker.Build input interface{} not allowed")
	}
	if inputReflectType.Kind() != reflect.Ptr {
		return errors.New("faker.Build input is not a pointer")
	}
	inputReflectValue := reflect.ValueOf(input)
	err := build(inputReflectValue, &fakerTag{})
	if err != nil {
		return err
	}
	return nil
}

func build(inputReflectValue reflect.Value, tag *fakerTag) error {
	inputReflectType := inputReflectValue.Type()
	kind := inputReflectType.Kind()
	// fmt.Println(">", kind.String(), inputReflectType.String(), tag)

	var (
		fn    builderFunc
		found bool
		key   string
	)

	key = builderKey(tag.funcName, inputReflectType.String())
	fn, found = builders[key]

	if found {
		var value interface{}
		var err error

		if tag.unique {
			value, err = Uniq(key, 3, func() (interface{}, error) {
				return fn(tag.params...)
			})
		} else {
			value, err = fn(tag.params...)
		}
		if err != nil {
			return err
		}
		if inputReflectValue.IsZero() {
			inputReflectValue.Set(reflect.ValueOf(value))
		}
		return nil
	}
	// if tag.funcName != "" {
	// 	return fmt.Errorf("Invalid faker function '%s' for type '%s'", tag.funcName, inputReflectType.String())
	// }

	switch kind {
	case reflect.Ptr:
		if inputReflectValue.IsNil() {
			newVar := reflect.New(inputReflectType.Elem())
			inputReflectValue.Set(newVar)
		}
		return build(inputReflectValue.Elem(), tag)
	case reflect.Struct:
		for i := 0; i < inputReflectValue.NumField(); i++ {
			fieldTag := decodeTag(inputReflectType.Field(i).Tag.Get(tagName))
			if fieldTag.mustSkip() {
				continue
			}
			if !inputReflectValue.Field(i).CanSet() {
				continue // to avoid panic to set on unexported field in struct
			}
			if inputReflectValue.Field(i).Kind() == reflect.Ptr && inputReflectValue.Field(i).Type().Elem() == inputReflectType {
				continue // do not enter in an infinite loop on recursive type
			}
			err := build(inputReflectValue.Field(i), fieldTag)
			if err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		if inputReflectValue.IsNil() {
			var sliceLen int
			if tag != nil && tag.length != 0 {
				sliceLen = tag.length
			} else {
				sliceLen = IntInRange(0, 9)
			}
			newSlice := reflect.MakeSlice(inputReflectType, sliceLen, sliceLen)
			inputReflectValue.Set(newSlice)
			for i := 0; i < newSlice.Len(); i++ {
				err := build(newSlice.Index(i), tag)
				if err != nil {
					return err
				}
			}
		}
	case reflect.Map:
		if inputReflectValue.IsNil() {
			var mapLen int
			if tag != nil && tag.length != 0 {
				mapLen = tag.length
			} else {
				mapLen = IntInRange(0, 9)
			}
			keyReflectType := inputReflectType.Key()
			elemReflectType := inputReflectType.Elem()
			newMap := reflect.MakeMap(inputReflectType)
			var (
				key  reflect.Value
				elem reflect.Value
				err  error
			)
			for i := 0; i < mapLen; i++ {
				key = reflect.New(keyReflectType).Elem()
				elem = reflect.New(elemReflectType).Elem()
				err = build(key, tag)
				if err != nil {
					return err
				}
				err = build(elem, tag)
				if err != nil {
					return err
				}
				newMap.SetMapIndex(key, elem)
			}
			inputReflectValue.Set(newMap)
		}
	default:
		if tag.funcName != "" {
			return fmt.Errorf("Invalid faker function '%s' for type '%s'", tag.funcName, inputReflectType.String())
		}
		return nil
	}
	return nil
}
