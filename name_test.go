package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleMaleFirstName() {
	faker.SetSeed(1400)
	fmt.Println(faker.MaleFirstName())
	// Output: Lucien
}

func ExampleFemaleFirstName() {
	faker.SetSeed(1401)
	fmt.Println(faker.FemaleFirstName())
	// Output: Sharell
}

func ExampleNeutralFirstName() {
	faker.SetSeed(1402)
	fmt.Println(faker.NeutralFirstName())
	// Output: Hayden
}

func ExampleFirstName() {
	faker.SetSeed(1403)
	fmt.Println(faker.FirstName())
	// Output: Barney
}

func ExampleLastName() {
	faker.SetSeed(1404)
	fmt.Println(faker.LastName())
	// Output: Ankunding
}

func ExampleNamePrefix() {
	faker.SetSeed(1405)
	fmt.Println(faker.NamePrefix())
	// Output: Msgr.
}

func ExampleNameSuffix() {
	faker.SetSeed(1406)
	fmt.Println(faker.NameSuffix())
	// Output: PhD
}

func ExampleFullName() {
	faker.SetSeed(1407)
	fmt.Println(faker.FullName())
	// Output: Sawyer Littel
}

func ExampleNameInitials() {
	faker.SetSeed(1408)
	fmt.Println(faker.NameInitials())
	// Output: NP
}

func TestNameBuild(t *testing.T) {
	faker.SetSeed(1008)
	s := &struct {
		Field1 string `faker:"MaleFirstName"`
		Field2 string `faker:"FemaleFirstName"`
		Field3 string `faker:"NeutralFirstName"`
		Field4 string `faker:"FirstName"`
		Field5 string `faker:"LastName"`
		Field6 string `faker:"NamePrefix"`
		Field7 string `faker:"NameSuffix"`
		Field8 string `faker:"FullName"`
		Field9 string `faker:"NameInitials"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Bennett", s.Field1)
	assert.Equal(t, "Christi", s.Field2)
	assert.Equal(t, "Emerson", s.Field3)
	assert.Equal(t, "Lowell", s.Field4)
	assert.Equal(t, "Grady", s.Field5)
	assert.Equal(t, "Pres.", s.Field6)
	assert.Equal(t, "LLD", s.Field7)
	assert.Equal(t, "Frederick Bauch", s.Field8)
	assert.Equal(t, "QN", s.Field9)
}
