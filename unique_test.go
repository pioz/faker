package factory_test

import (
	"errors"
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	factory.SetSeed(400)
	values := make([]int, 0)
	for i := 0; i < 11; i++ {
		value, err := factory.Uniq("test", 0, func() (interface{}, error) { return factory.IntInRange(0, 10), nil })
		assert.Nil(t, err)
		values = append(values, value.(int))
	}

	factory.ClearUniqCache("notexistingkey")
	_, err := factory.Uniq("test", 0, func() (interface{}, error) { return factory.IntInRange(0, 10), nil })
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to generate a unique value", err.Error())

	factory.ClearUniqCache("test")
	_, err = factory.Uniq("test", 0, func() (interface{}, error) { return factory.IntInRange(0, 0), nil })
	assert.Nil(t, err)

	_, err = factory.Uniq("test", 0, func() (interface{}, error) { return factory.IntInRange(0, 0), nil })
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to generate a unique value", err.Error())

	factory.ClearAllUniqCache()
	_, err = factory.Uniq("test", 0, func() (interface{}, error) { return factory.IntInRange(0, 0), nil })
	assert.Nil(t, err)
}

func TestUniqError(t *testing.T) {
	_, err := factory.Uniq("test", 0, func() (interface{}, error) { return nil, errors.New("This is an error") })
	assert.NotNil(t, err)
	assert.Equal(t, "This is an error", err.Error())
}

func TestUniqSlice(t *testing.T) {
	factory.SetSeed(401)
	randIntSlice, err := factory.UniqSlice(100, "testSlice", 0, func() (interface{}, error) {
		return factory.IntInRange(1, 100), nil
	})
	t.Log(randIntSlice)
	assert.Nil(t, err)
	assert.Equal(t, 100, len(randIntSlice))
	for _, randInt := range randIntSlice {
		assert.NotEmpty(t, randInt)
	}
}

func TestUniqSliceError(t *testing.T) {
	factory.SetSeed(402)
	randIntSlice, err := factory.UniqSlice(10, "testSlice", 0, func() (interface{}, error) {
		return factory.IntInRange(1, 9), nil
	})
	t.Log(randIntSlice)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to generate a unique value", err.Error())
}
