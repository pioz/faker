package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestMaleFirstName(t *testing.T) {
	factory.SetSeed(1400)
	value := factory.MaleFirstName()
	t.Log(value)
	assert.Equal(t, "Lucien", value)
}

func TestFemaleFirstName(t *testing.T) {
	factory.SetSeed(1401)
	value := factory.FemaleFirstName()
	t.Log(value)
	assert.Equal(t, "Sharell", value)
}

func TestNeutralFirstName(t *testing.T) {
	factory.SetSeed(1402)
	value := factory.NeutralFirstName()
	t.Log(value)
	assert.Equal(t, "Hayden", value)
}

func TestFirstName(t *testing.T) {
	factory.SetSeed(1403)
	value := factory.FirstName()
	t.Log(value)
	assert.Equal(t, "Barney", value)
}

func TestLastName(t *testing.T) {
	factory.SetSeed(1404)
	value := factory.LastName()
	t.Log(value)
	assert.Equal(t, "Ankunding", value)
}

func TestNamePrefix(t *testing.T) {
	factory.SetSeed(1405)
	value := factory.NamePrefix()
	t.Log(value)
	assert.Equal(t, "Msgr.", value)
}

func TestNameSuffix(t *testing.T) {
	factory.SetSeed(1406)
	value := factory.NameSuffix()
	t.Log(value)
	assert.Equal(t, "PhD", value)
}

func TestFullName(t *testing.T) {
	factory.SetSeed(1407)
	value := factory.FullName()
	t.Log(value)
	assert.Equal(t, "Sawyer Littel", value)
}

func TestNameInitials(t *testing.T) {
	factory.SetSeed(1408)
	value := factory.NameInitials()
	t.Log(value)
	assert.Equal(t, "NP", value)
}

func TestNameBuild(t *testing.T) {
	factory.SetSeed(1008)
	s := &struct {
		Field1 string `factory:"MaleFirstName"`
		Field2 string `factory:"FemaleFirstName"`
		Field3 string `factory:"NeutralFirstName"`
		Field4 string `factory:"FirstName"`
		Field5 string `factory:"LastName"`
		Field6 string `factory:"NamePrefix"`
		Field7 string `factory:"NameSuffix"`
		Field8 string `factory:"FullName"`
		Field9 string `factory:"NameInitials"`
	}{}
	err := factory.Build(&s)
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
