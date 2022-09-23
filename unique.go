package faker

import (
	"fmt"
	"sync"
)

const uniqueDefaultMaxRetry = 10000

var uniqueCacheMutex = &sync.Mutex{}
var uniqueCache = make(map[string]map[interface{}]int)

// Uniq run max maxRetry times fn function until fn returns a unique value for
// the group of runs group. Returns error if the number of runs reach
// maxRetry. If maxRetry is zero use default value that is 10000.
func Uniq(group string, maxRetry int, fn func() (interface{}, error)) (interface{}, error) {
	if maxRetry == 0 {
		maxRetry = uniqueDefaultMaxRetry
	}
	uniqueCacheMutex.Lock()
	defer uniqueCacheMutex.Unlock()
	if _, found := uniqueCache[group]; !found {
		uniqueCache[group] = make(map[interface{}]int)
	}
	var value interface{}
	var err error
	for i := 0; i < maxRetry; i++ {
		value, err = fn()
		if err != nil {
			return nil, err
		}
		if _, found := uniqueCache[group][value]; !found {
			uniqueCache[group][value] = 0
			return value, nil
		}
	}
	// If the value is a string avoid to return an error, and generate a unique
	// value adding an index after the string.
	valueAsString, ok := value.(string)
	if ok {
		uniqValue := fmt.Sprintf("%s %d", valueAsString, uniqueCache[group][value])
		uniqueCache[group][value]++
		return uniqValue, nil
	}
	return value, fmt.Errorf("failed to generate a unique value for group '%s'", group)
}

// UniqSlice run max maxRetry times fn function until fn returns a unique
// slice for the group of runs group. Returns error if the number of runs
// reach maxRetry. If maxRetry is zero use default value that is 10000.
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

// ClearUniqCache delete all results for the group group.
func ClearUniqCache(group string) {
	uniqueCacheMutex.Lock()
	delete(uniqueCache, group)
	uniqueCacheMutex.Unlock()
}

// ClearAllUniqCache delete all results for all groups of run.
func ClearAllUniqCache() {
	uniqueCacheMutex.Lock()
	uniqueCache = make(map[string]map[interface{}]int)
	uniqueCacheMutex.Unlock()
}
