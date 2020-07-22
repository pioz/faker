package factory

import (
	"fmt"
	"sync"

	"github.com/pioz/factory/data"
)

var poolData = data.PoolData{
	"timezone": data.Timezone,
}

var dataMutex = &sync.Mutex{}

func GetData(namespace, group string) (interface{}, error) {
	dataMutex.Lock()
	defer dataMutex.Unlock()
	var (
		poolGroup data.PoolGroup
		pool      data.Pool
		found     bool
	)
	poolGroup, found = poolData[namespace]
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
	dataMutex.Lock()
	defer dataMutex.Unlock()
	var (
		poolGroup data.PoolGroup
		found     bool
	)
	_, found = poolData[namespace]
	if !found {
		poolData[namespace] = make(data.PoolGroup)
	}
	poolGroup = poolData[namespace]
	poolGroup[group] = pool
}

func SetPoolGroup(namespace string, poolGroup data.PoolGroup) {
	dataMutex.Lock()
	defer dataMutex.Unlock()
	poolData[namespace] = poolGroup
}
