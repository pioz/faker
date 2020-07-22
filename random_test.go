package factory_test

import (
	"math/rand"
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestSetRand(t *testing.T) {
	random := rand.New(rand.NewSource(21))
	assert.NotPanics(t, func() {
		factory.SetRand(random)
	})
}

func TestSetSeed(t *testing.T) {
	assert.NotPanics(t, func() {
		factory.SetSeed(21)
	})
}
