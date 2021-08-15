package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleColorName() {
	faker.SetSeed(1900)
	fmt.Println(faker.ColorName())
	// Output: apricot
}

func ExampleColorHex() {
	faker.SetSeed(1901)
	fmt.Println(faker.ColorHex())
	// Output: #89f01b
}

func ExampleColorRGB() {
	faker.SetSeed(1902)
	fmt.Println(faker.ColorRGB())
	// Output: [25 8 176]
}

func ExampleColorHSL() {
	faker.SetSeed(1903)
	fmt.Println(faker.ColorHSL())
	// Output: [135 5 41]
}

func TestColorBuild(t *testing.T) {
	faker.SetSeed(1904)
	s := &struct {
		Field1 string `faker:"ColorName"`
		Field2 string `faker:"ColorHex"`
		Field3 [3]int `faker:"ColorRGB"`
		Field4 [3]int `faker:"ColorHSL"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "lilac", s.Field1)
	assert.Equal(t, "#c3e065", s.Field2)
	assert.Equal(t, [3]int{46, 59, 194}, s.Field3)
	assert.Equal(t, [3]int{56, 22, 19}, s.Field4)
}
