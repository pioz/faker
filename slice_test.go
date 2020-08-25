package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleSlice() {
	faker.SetSeed(200)
	fmt.Println(faker.Slice(3, func() interface{} { return faker.IntInRange(0, 10) }))
	// Output: [6 9 9]
}

func ExampleSample() {
	faker.SetSeed(201)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(faker.Sample(slice))
	// Output: 5
}

func TestSampleWithNoSliceArg(t *testing.T) {
	assert.PanicsWithValue(t, "faker.Sample param must be a slice", func() {
		faker.Sample(2)
	})
}

func TestSampleWithNilSlice(t *testing.T) {
	var slice []int
	assert.PanicsWithValue(t, "faker.Sample param is nil", func() {
		faker.Sample(slice)
	})
}

func TestSampleWithEmptySlice(t *testing.T) {
	slice := make([]int, 0)
	assert.PanicsWithValue(t, "faker.Sample param is empty", func() {
		faker.Sample(slice)
	})
}
