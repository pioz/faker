package factory

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

func (f providerFunc) Error() string {
	return "providerFunc error"
}

var providers = make(map[string]providerFunc)

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
