package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleCurrencyName() {
	faker.SetSeed(1200)
	fmt.Println(faker.CurrencyName())
	// Output: Sierra Leonean Leone
}

func ExampleCurrencyCode() {
	faker.SetSeed(1201)
	fmt.Println(faker.CurrencyCode())
	// Output: XOF
}

func ExampleCurrencySymbol() {
	faker.SetSeed(1202)
	fmt.Println(faker.CurrencySymbol())
	// Output: ₮
}

func TestCurrencyBuild(t *testing.T) {
	faker.SetSeed(1210)
	s := &struct {
		Field1 string `faker:"CurrencyName"`
		Field2 string `faker:"CurrencyCode"`
		Field3 string `faker:"CurrencySymbol"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Bermudian Dollar", s.Field1)
	assert.Equal(t, "XCD", s.Field2)
	assert.Equal(t, "K", s.Field3)
}
