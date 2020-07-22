package factory

import (
	"strings"
)

const digits = "0123456789"
const lowerLetters = "abcdefghijklmnopqrstuvwxyz"
const upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charset = digits + lowerLetters + upperLetters

func StringWithSize(size int) string {
	return stringWithSize(size, charset)
}

func String() string {
	return StringWithSize(IntInRange(1, 255))
}

func DigitsWithSize(size int) string {
	return stringWithSize(size, digits)
}

func Digits() string {
	return DigitsWithSize(IntInRange(1, 255))
}

func LettersWithSize(size int) string {
	return stringWithSize(size, lowerLetters+upperLetters)
}

func Letters() string {
	return LettersWithSize(IntInRange(1, 255))
}

func Lexify(str string) string {
	return replaceChar(str, "?", func() string { return LettersWithSize(1) })
}

func Numerify(str string) string {
	return replaceChar(str, "?", func() string { return DigitsWithSize(1) })
}

func stringWithSize(size int, charset string) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}

func replaceChar(str, chr string, fn func() string) string {
	r := str
	count := strings.Count(str, chr)
	for i := 0; i < count; i++ {
		r = strings.Replace(r, chr, fn(), 1)
	}
	return r
}

// Provider functions

func stringWithSizeFn(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return StringWithSize(size), nil
}

func stringFn(params ...string) (interface{}, error) {
	return String(), nil
}

func digitsWithSizeFn(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return DigitsWithSize(size), nil
}

func digitsFn(params ...string) (interface{}, error) {
	return Digits(), nil
}

func lettersWithSizeFn(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return LettersWithSize(size), nil
}

func lettersFn(params ...string) (interface{}, error) {
	return Letters(), nil
}

func lexifyFn(params ...string) (interface{}, error) {
	if len(params) != 1 {
		return nil, parametersError(nil)
	}
	return Lexify(params[0]), nil
}

func numerifyFn(params ...string) (interface{}, error) {
	if len(params) != 1 {
		return nil, parametersError(nil)
	}
	return Numerify(params[0]), nil
}
