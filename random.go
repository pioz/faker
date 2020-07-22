package factory

import (
	"math/rand"
	"sync"
	"time"
)

var (
	random      = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomMutex = &sync.Mutex{}
)

func SetRand(r *rand.Rand) {
	randomMutex.Lock()
	random = r
	randomMutex.Unlock()
}

func SetSeed(seed int64) {
	randomMutex.Lock()
	random.Seed(seed)
	randomMutex.Unlock()
}
