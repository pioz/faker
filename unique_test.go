package faker_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleUniq() {
	faker.SetSeed(403)
	generator := func() (interface{}, error) { return faker.IntInRange(0, 1), nil }
	value1, _ := faker.Uniq("test", 0, generator)
	fmt.Println(value1)
	value2, _ := faker.Uniq("test", 0, generator)
	fmt.Println(value2)
	faker.ClearUniqCache("test")
	value3, _ := faker.Uniq("test", 0, generator)
	fmt.Println(value3)
	// Output: 1
	// 0
	// 0
}

func TestUniq(t *testing.T) {
	faker.SetSeed(400)
	values := make([]int, 0)
	for i := 0; i < 11; i++ {
		value, err := faker.Uniq("test", 0, func() (interface{}, error) { return faker.IntInRange(0, 10), nil })
		assert.Nil(t, err)
		values = append(values, value.(int))
	}
	assert.Equal(t, 11, len(values))

	faker.ClearUniqCache("notexistingkey")
	_, err := faker.Uniq("test", 0, func() (interface{}, error) { return faker.IntInRange(0, 10), nil })
	assert.NotNil(t, err)
	assert.Equal(t, "failed to generate a unique value for group 'test'", err.Error())

	faker.ClearUniqCache("test")
	_, err = faker.Uniq("test", 0, func() (interface{}, error) { return faker.IntInRange(0, 0), nil })
	assert.Nil(t, err)

	_, err = faker.Uniq("test", 0, func() (interface{}, error) { return faker.IntInRange(0, 0), nil })
	assert.NotNil(t, err)
	assert.Equal(t, "failed to generate a unique value for group 'test'", err.Error())

	faker.ClearAllUniqCache()
	_, err = faker.Uniq("test", 0, func() (interface{}, error) { return faker.IntInRange(0, 0), nil })
	assert.Nil(t, err)
}

func TestUniqError(t *testing.T) {
	_, err := faker.Uniq("test", 0, func() (interface{}, error) { return nil, errors.New("this is an error") })
	assert.NotNil(t, err)
	assert.Equal(t, "this is an error", err.Error())
}

func TestUniqSlice(t *testing.T) {
	faker.SetSeed(401)
	randIntSlice, err := faker.UniqSlice(100, "testSlice", 0, func() (interface{}, error) {
		return faker.IntInRange(1, 100), nil
	})
	t.Log(randIntSlice)
	assert.Nil(t, err)
	assert.Equal(t, 100, len(randIntSlice))
	for _, randInt := range randIntSlice {
		assert.NotEmpty(t, randInt)
	}
}

func TestUniqSliceError(t *testing.T) {
	faker.SetSeed(402)
	randIntSlice, err := faker.UniqSlice(10, "testSlice", 0, func() (interface{}, error) {
		return faker.IntInRange(1, 9), nil
	})
	t.Log(randIntSlice)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to generate a unique value for group 'testSlice'", err.Error())
}
