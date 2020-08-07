package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleUsername() {
	faker.SetSeed(1700)
	fmt.Println(faker.Username())
	// Output: polychasium
}

func ExampleDomain() {
	faker.SetSeed(1701)
	fmt.Println(faker.Domain())
	// Output: lorusso.name
}

func ExampleEmail() {
	faker.SetSeed(1702)
	fmt.Println(faker.Email())
	// Output: homocercal@fulmer.net
}

func ExampleFreeEmail() {
	faker.SetSeed(1703)
	fmt.Println(faker.FreeEmail())
	// Output: atlas@gmail.com
}

func ExampleSafeEmail() {
	faker.SetSeed(1704)
	fmt.Println(faker.SafeEmail())
	// Output: disbelieve@example.com
}

func ExampleSlug() {
	faker.SetSeed(1705)
	fmt.Println(faker.Slug())
	// Output: a-reliable-seal-s-bee-comes-with-it-the-thought-that-the-adventurous-giraffe-is-an-alligator
}

func ExampleURL() {
	faker.SetSeed(1706)
	fmt.Println(faker.URL())
	// Output: https://www.mcmillon.info/this-could-be-or-perhaps-their-alligator-was-in-this-moment-an-eager-spider
}

func TestInternetBuild(t *testing.T) {
	faker.SetSeed(1720)
	s := &struct {
		Field1 string `faker:"Username"`
		Field2 string `faker:"Domain"`
		Field3 string `faker:"Email"`
		Field4 string `faker:"SafeEmail"`
		Field5 string `faker:"FreeEmail"`
		Field6 string `faker:"Slug"`
		Field7 string `faker:"Url"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "tight", s.Field1)
	assert.Equal(t, "extortionate.info", s.Field2)
	assert.Equal(t, "euchology@farris.info", s.Field3)
	assert.Equal(t, "andino@example.com", s.Field4)
	assert.Equal(t, "inflorescence@gmail.com", s.Field5)
	assert.Equal(t, "a-friendly-blackberry-is-an-eagle-of-the-mind", s.Field6)
	assert.Equal(t, "https://www.condescendence.info/nowhere-is-it-disputed-that-their-blueberry-was-in-this-moment-a-plausible-plum", s.Field7)
}
