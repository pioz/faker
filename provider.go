package factory

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var providers = map[string]providerFunc{
	// misc
	providerKey("Bool", "bool"): boolFn,
	providerKey("", "bool"):     boolFn,
	// number
	providerKey("IntInRange", "int"):         intInRangeFn,
	providerKey("Int", "int"):                intFn,
	providerKey("", "int"):                   intFn,
	providerKey("Int64InRange", "int64"):     int64InRangeFn,
	providerKey("Int64", "int64"):            int64Fn,
	providerKey("", "int64"):                 int64Fn,
	providerKey("Int32InRange", "int32"):     int32InRangeFn,
	providerKey("Int32", "int32"):            int32Fn,
	providerKey("", "int32"):                 int32Fn,
	providerKey("Int16InRange", "int16"):     int16InRangeFn,
	providerKey("Int16", "int16"):            int16Fn,
	providerKey("", "int16"):                 int16Fn,
	providerKey("Int8InRange", "int8"):       int8InRangeFn,
	providerKey("Int8", "int8"):              int8Fn,
	providerKey("", "int8"):                  int8Fn,
	providerKey("UintInRange", "uint"):       uintInRangeFn,
	providerKey("Uint", "uint"):              uintFn,
	providerKey("", "uint"):                  uintFn,
	providerKey("Uint64InRange", "uint64"):   uint64InRangeFn,
	providerKey("Uint64", "uint64"):          uint64Fn,
	providerKey("", "uint64"):                uint64Fn,
	providerKey("Uint32InRange", "uint32"):   uint32InRangeFn,
	providerKey("Uint32", "uint32"):          uint32Fn,
	providerKey("", "uint32"):                uint32Fn,
	providerKey("Uint16InRange", "uint16"):   uint16InRangeFn,
	providerKey("Uint16", "uint16"):          uint16Fn,
	providerKey("", "uint16"):                uint16Fn,
	providerKey("Uint8InRange", "uint8"):     uint8InRangeFn,
	providerKey("Uint8", "uint8"):            uint8Fn,
	providerKey("", "uint8"):                 uint8Fn,
	providerKey("Float64InRange", "float64"): float64InRangeFn,
	providerKey("Float64", "float64"):        float64Fn,
	providerKey("", "float64"):               float64Fn,
	providerKey("Float32InRange", "float32"): float32InRangeFn,
	providerKey("Float32", "float32"):        float32Fn,
	providerKey("", "float32"):               float32Fn,
	// string
	providerKey("StringWithSize", "string"):  stringWithSizeFn,
	providerKey("String", "string"):          stringFn,
	providerKey("", "string"):                stringFn,
	providerKey("DigitsWithSize", "string"):  digitsWithSizeFn,
	providerKey("Digits", "string"):          digitsFn,
	providerKey("LettersWithSize", "string"): lettersWithSizeFn,
	providerKey("Letters", "string"):         lettersFn,
	providerKey("Lexify", "string"):          lexifyFn,
	providerKey("Numerify", "string"):        numerifyFn,
	// time
	providerKey("DurationInRange", "time.Duration"): durationInRangeFn,
	providerKey("Duration", "time.Duration"):        durationFn,
	providerKey("", "time.Duration"):                durationFn,
	providerKey("Time", "time.Time"):                timeFn,
	providerKey("", "time.Time"):                    timeFn,
	providerKey("NanoSecond", "int"):                nanoSecondFn,
	providerKey("Second", "int"):                    secondFn,
	providerKey("Minute", "int"):                    minuteFn,
	providerKey("Hour", "int"):                      hourFn,
	providerKey("Day", "int"):                       dayFn,
	providerKey("WeekDay", "string"):                weekDayFn,
	providerKey("Month", "string"):                  monthFn,
	providerKey("Year", "int"):                      yearFn,
	providerKey("TimeZone", "string"):               timeZoneFn,
	providerKey("TimeZoneAbbr", "string"):           timeZoneAbbrFn,
	providerKey("TimeZoneFull", "string"):           timeZoneFullFn,
	providerKey("TimeZoneOffset", "float32"):        timeZoneOffsetFn,
	providerKey("TimeZoneRegion", "string"):         timeZoneRegionFn,
}

const skipTag = "-"
const uniqueTag = "unique"

var (
	tagFnCallRegexp = regexp.MustCompile(`(.+?)\((.+?)\)`)
	tagLenRegexp    = regexp.MustCompile(`len=(\d+)`)
)

type fakerTag struct {
	funcName string
	unique   bool
	length   int
	params   []string
}

func (tag *fakerTag) mustSkip() bool {
	return tag.funcName == skipTag
}

func decodeTag(tagString string) *fakerTag {
	tag := &fakerTag{}
	for _, token := range strings.Split(tagString, ";") {
		if token == skipTag {
			tag.funcName = skipTag
			return tag
		}
		if token == uniqueTag {
			tag.unique = true
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

type providerFunc func(...string) (interface{}, error)

func providerKey(providerName, providerType string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", providerName, providerType))
}

func RegisterProvider(providerName, providerType string, fn providerFunc) error {
	key := providerKey(providerName, providerType)
	if _, ok := providers[key]; ok {
		return errors.New("Provider already registered")
	}
	providers[key] = fn
	return nil
}

func UnregisterProvider(providerName, providerType string) error {
	key := providerKey(providerName, providerType)
	if _, ok := providers[key]; !ok {
		return errors.New("Provider not registered")
	}
	delete(providers, key)
	return nil
}
