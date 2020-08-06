package faker

import (
	"fmt"
	"sync"

	"github.com/pioz/faker/data"
)

var dbMutex = &sync.Mutex{}

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
		return "", fmt.Errorf("The namespace '%s' does not exist", namespace)
	}

	pool, found = poolGroup[group]
	if !found {
		return "", fmt.Errorf("The group '%s' in namespace '%s' does not exist", group, namespace)
	}

	i := IntInRange(0, len(pool)-1)
	return pool[i], nil
}

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

func SetPoolGroup(namespace string, poolGroup data.PoolGroup) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	data.DB[namespace] = poolGroup
}
