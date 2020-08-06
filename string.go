package faker

import (
	"regexp"
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

func Parameterize(str string) string {
	notAlphaNumRegExp := regexp.MustCompile("[^A-Za-z0-9]+")
	firstLastDashRegexp := regexp.MustCompile("^-|-$")

	parameterizeString := notAlphaNumRegExp.ReplaceAllString(str, "-")
	parameterizeString = firstLastDashRegexp.ReplaceAllString(parameterizeString, "")
	return strings.ToLower(parameterizeString)
}

func Pick(pool ...string) string {
	if len(pool) == 0 {
		return ""
	}
	i := IntInRange(0, len(pool)-1)
	return pool[i]
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

// Builder functions

func stringWithSizeBuilder(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return StringWithSize(size), nil
}

func stringBuilder(params ...string) (interface{}, error) {
	return String(), nil
}

func digitsWithSizeBuilder(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return DigitsWithSize(size), nil
}

func digitsBuilder(params ...string) (interface{}, error) {
	return Digits(), nil
}

func lettersWithSizeBuilder(params ...string) (interface{}, error) {
	size, err := paramsToInt(params...)
	if err != nil {
		return nil, err
	}
	return LettersWithSize(size), nil
}

func lettersBuilder(params ...string) (interface{}, error) {
	return Letters(), nil
}

func lexifyBuilder(params ...string) (interface{}, error) {
	if len(params) != 1 {
		return nil, parametersError(nil)
	}
	return Lexify(params[0]), nil
}

func numerifyBuilder(params ...string) (interface{}, error) {
	if len(params) != 1 {
		return nil, parametersError(nil)
	}
	return Numerify(params[0]), nil
}

func parameterizeBuilder(params ...string) (interface{}, error) {
	if len(params) != 1 {
		return nil, parametersError(nil)
	}
	return Parameterize(params[0]), nil
}

func pickBuilder(params ...string) (interface{}, error) {
	return Pick(params...), nil
}
