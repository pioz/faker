package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	faker.SetSeed(200)
	slice := faker.Slice(3, func() interface{} { return faker.IntInRange(0, 10) })
	assert.Equal(t, 3, len(slice))
	assert.Equal(t, 6, slice[0])
	assert.Equal(t, 9, slice[1])
	assert.Equal(t, 9, slice[2])
}
