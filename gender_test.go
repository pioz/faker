package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleGender() {
	faker.SetSeed(1500)
	fmt.Println(faker.Gender())
	// Output: Cisgender Male
}

func ExampleBinaryGender() {
	faker.SetSeed(1501)
	fmt.Println(faker.BinaryGender())
	// Output: Male
}

func ExampleShortBinaryGender() {
	faker.SetSeed(1502)
	fmt.Println(faker.ShortBinaryGender())
	// Output: m
}

func TestGenderBuild(t *testing.T) {
	faker.SetSeed(1520)
	s := &struct {
		Field1 string `faker:"Gender"`
		Field2 string `faker:"BinaryGender"`
		Field3 string `faker:"ShortBinaryGender"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Cis", s.Field1)
	assert.Equal(t, "Female", s.Field2)
	assert.Equal(t, "m", s.Field3)
}
