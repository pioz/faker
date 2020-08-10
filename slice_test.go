package faker_test

import (
	"fmt"

	"github.com/pioz/faker"
)

func ExampleSlice() {
	faker.SetSeed(200)
	fmt.Println(faker.Slice(3, func() interface{} { return faker.IntInRange(0, 10) }))
	// Output: [6 9 9]
}
