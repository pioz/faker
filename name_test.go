package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestMaleFirstName(t *testing.T) {
	faker.SetSeed(1400)
	value := faker.MaleFirstName()
	t.Log(value)
	assert.Equal(t, "Lucien", value)
}

func TestFemaleFirstName(t *testing.T) {
	faker.SetSeed(1401)
	value := faker.FemaleFirstName()
	t.Log(value)
	assert.Equal(t, "Sharell", value)
}

func TestNeutralFirstName(t *testing.T) {
	faker.SetSeed(1402)
	value := faker.NeutralFirstName()
	t.Log(value)
	assert.Equal(t, "Hayden", value)
}

func TestFirstName(t *testing.T) {
	faker.SetSeed(1403)
	value := faker.FirstName()
	t.Log(value)
	assert.Equal(t, "Barney", value)
}

func TestLastName(t *testing.T) {
	faker.SetSeed(1404)
	value := faker.LastName()
	t.Log(value)
	assert.Equal(t, "Ankunding", value)
}

func TestNamePrefix(t *testing.T) {
	faker.SetSeed(1405)
	value := faker.NamePrefix()
	t.Log(value)
	assert.Equal(t, "Msgr.", value)
}

func TestNameSuffix(t *testing.T) {
	faker.SetSeed(1406)
	value := faker.NameSuffix()
	t.Log(value)
	assert.Equal(t, "PhD", value)
}

func TestFullName(t *testing.T) {
	faker.SetSeed(1407)
	value := faker.FullName()
	t.Log(value)
	assert.Equal(t, "Sawyer Littel", value)
}

func TestNameInitials(t *testing.T) {
	faker.SetSeed(1408)
	value := faker.NameInitials()
	t.Log(value)
	assert.Equal(t, "NP", value)
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
