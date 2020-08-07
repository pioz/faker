package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleCountryName() {
	faker.SetSeed(1000)
	fmt.Println(faker.CountryName())
	// Output: Rwanda
}

func ExampleCountryAlpha2() {
	faker.SetSeed(1001)
	fmt.Println(faker.CountryAlpha2())
	// Output: ML
}

func ExampleCountryAlpha3() {
	faker.SetSeed(1002)
	fmt.Println(faker.CountryAlpha3())
	// Output: TKM
}

func ExampleCountryNationality() {
	faker.SetSeed(1002)
	fmt.Println(faker.CountryNationality())
	// Output: Venezuelan
}

func ExampleCountryFlag() {
	faker.SetSeed(1003)
	fmt.Println(faker.CountryFlag())
	// Output: 🇳🇷
}

func TestCountryBuild(t *testing.T) {
	faker.SetSeed(1010)
	s := &struct {
		Field1 string `faker:"CountryName"`
		Field2 string `faker:"CountryAlpha2"`
		Field3 string `faker:"CountryAlpha3"`
		Field4 string `faker:"CountryNationality"`
		Field5 string `faker:"CountryFlag"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Iceland", s.Field1)
	assert.Equal(t, "IN", s.Field2)
	assert.Equal(t, "RUS", s.Field3)
	assert.Equal(t, "Cocos Islander", s.Field4)
	assert.Equal(t, "🇲🇹", s.Field5)
}
