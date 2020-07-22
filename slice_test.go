package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	factory.SetSeed(200)
	slice := factory.Slice(3, func() interface{} { return factory.IntInRange(0, 10) })
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 6, slice[0])
	assert.Equal(t, 9, slice[1])
	assert.Equal(t, 9, slice[2])
}
