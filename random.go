package faker

import (
	"math/rand"
	"sync"
	"time"
)

var (
	random      = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomMutex = &sync.Mutex{}
)

// SetRand set a new source of random numbers (see rand.Rand). The default
// source is:
//
//   rand.New(rand.NewSource(time.Now().UnixNano()))
//
func SetRand(r *rand.Rand) {
	randomMutex.Lock()
	random = r
	randomMutex.Unlock()
}

// SetSeed uses the provided seed value to initialize the generator to a
// deterministic state (see rand.Seed).
func SetSeed(seed int64) {
	randomMutex.Lock()
	random.Seed(seed)
	randomMutex.Unlock()
}
