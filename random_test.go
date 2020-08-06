package faker_test

import (
	"math/rand"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestSetRand(t *testing.T) {
	random := rand.New(rand.NewSource(21))
	assert.NotPanics(t, func() {
		faker.SetRand(random)
	})
}

func TestSetSeed(t *testing.T) {
	assert.NotPanics(t, func() {
		faker.SetSeed(21)
	})
}
