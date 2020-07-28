package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestCountryName(t *testing.T) {
	factory.SetSeed(1000)
	value := factory.CountryName()
	t.Log(value)
	assert.Equal(t, "Rwanda", value)
}

func TestCountryAlpha2(t *testing.T) {
	factory.SetSeed(1001)
	value := factory.CountryAlpha2()
	t.Log(value)
	assert.Equal(t, "ML", value)
}

func TestCountryAlpha3(t *testing.T) {
	factory.SetSeed(1002)
	value := factory.CountryAlpha3()
	t.Log(value)
	assert.Equal(t, "TKM", value)
}

func TestCountryNationality(t *testing.T) {
	factory.SetSeed(1002)
	value := factory.CountryNationality()
	t.Log(value)
	assert.Equal(t, "Venezuelan", value)
}

func TestCountryFlag(t *testing.T) {
	factory.SetSeed(1003)
	value := factory.CountryFlag()
	t.Log(value)
	assert.Equal(t, "🇳🇷", value)
}

func TestCountryBuild(t *testing.T) {
	factory.SetSeed(1010)
	s := &struct {
		Field1 string `factory:"CountryName"`
		Field2 string `factory:"CountryAlpha2"`
		Field3 string `factory:"CountryAlpha3"`
		Field4 string `factory:"CountryNationality"`
		Field5 string `factory:"CountryFlag"`
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Iceland", s.Field1)
	assert.Equal(t, "IN", s.Field2)
	assert.Equal(t, "RUS", s.Field3)
	assert.Equal(t, "Cocos Islander", s.Field4)
	assert.Equal(t, "🇲🇹", s.Field5)
}
