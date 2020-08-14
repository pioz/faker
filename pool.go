package faker

import (
	"fmt"
	"sync"

	"github.com/pioz/faker/data"
)

var dbMutex = &sync.Mutex{}

// GetData return a random value of the Pool present in the group group with
// namespace namespace or error if the pool does not exist. Faker organize
// fake data in a map of string and map of string and array of interface. The
// keys of the first level map are called namespaces, the keys of the second
// level map are called groups.
func GetData(namespace, group string) (interface{}, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var (
		poolGroup data.PoolGroup
		pool      data.Pool
		found     bool
	)
	poolGroup, found = data.DB[namespace]
	if !found {
		return "", fmt.Errorf("the namespace '%s' does not exist", namespace)
	}

	pool, found = poolGroup[group]
	if !found {
		return "", fmt.Errorf("the group '%s' in namespace '%s' does not exist", group, namespace)
	}

	i := IntInRange(0, len(pool)-1)
	return pool[i], nil
}

// SetPool add a new Pool under the group group with namespace namespace (see
// GetData).
func SetPool(namespace, group string, pool data.Pool) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	var (
		poolGroup data.PoolGroup
		found     bool
	)
	_, found = data.DB[namespace]
	if !found {
		data.DB[namespace] = make(data.PoolGroup)
	}
	poolGroup = data.DB[namespace]
	poolGroup[group] = pool
}

// SetPoolGroup add a new PoolGroup under the namespace namespace (see
// GetData).
func SetPoolGroup(namespace string, poolGroup data.PoolGroup) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	data.DB[namespace] = poolGroup
}
