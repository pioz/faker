package factory

import (
	"errors"
	"sync"
)

const uniqueDefaultMaxRetry = 10000

var uniqueCacheMutex = &sync.Mutex{}
var uniqueCache = make(map[string]map[interface{}]struct{})

func Uniq(group string, maxRetry int, fn func() (interface{}, error)) (interface{}, error) {
	if maxRetry == 0 {
		maxRetry = uniqueDefaultMaxRetry
	}
	uniqueCacheMutex.Lock()
	defer uniqueCacheMutex.Unlock()
	if _, ok := uniqueCache[group]; !ok {
		uniqueCache[group] = make(map[interface{}]struct{})
	}
	var value interface{}
	var err error
	for i := 0; i < maxRetry; i++ {
		value, err = fn()
		if err != nil {
			return nil, err
		}
		if _, ok := uniqueCache[group][value]; !ok {
			uniqueCache[group][value] = struct{}{}
			return value, nil
		}
	}
	return value, errors.New("Failed to generate a unique value")
}

func UniqSlice(size int, group string, maxRetry int, fn func() (interface{}, error)) ([]interface{}, error) {
	var value interface{}
	var err error
	slice := Slice(size, func() interface{} {
		value, err = Uniq(group, maxRetry, fn)
		return value
	})
	if err != nil {
		return nil, err
	}
	return slice, nil
}

func ClearUniqCache(key string) {
	uniqueCacheMutex.Lock()
	delete(uniqueCache, key)
	uniqueCacheMutex.Unlock()
}

func ClearAllUniqCache() {
	uniqueCacheMutex.Lock()
	uniqueCache = make(map[string]map[interface{}]struct{})
	uniqueCacheMutex.Unlock()
}
