package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestCountryName(t *testing.T) {
	faker.SetSeed(1000)
	value := faker.CountryName()
	t.Log(value)
	assert.Equal(t, "Rwanda", value)
}

func TestCountryAlpha2(t *testing.T) {
	faker.SetSeed(1001)
	value := faker.CountryAlpha2()
	t.Log(value)
	assert.Equal(t, "ML", value)
}

func TestCountryAlpha3(t *testing.T) {
	faker.SetSeed(1002)
	value := faker.CountryAlpha3()
	t.Log(value)
	assert.Equal(t, "TKM", value)
}

func TestCountryNationality(t *testing.T) {
	faker.SetSeed(1002)
	value := faker.CountryNationality()
	t.Log(value)
	assert.Equal(t, "Venezuelan", value)
}

func TestCountryFlag(t *testing.T) {
	faker.SetSeed(1003)
	value := faker.CountryFlag()
	t.Log(value)
	assert.Equal(t, "🇳🇷", value)
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
