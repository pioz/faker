package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyName(t *testing.T) {
	faker.SetSeed(1200)
	value := faker.CurrencyName()
	t.Log(value)
	assert.Equal(t, "Sierra Leonean Leone", value)
}

func TestCurrencyCode(t *testing.T) {
	faker.SetSeed(1201)
	value := faker.CurrencyCode()
	t.Log(value)
	assert.Equal(t, "XOF", value)
}

func TestCurrencySymbol(t *testing.T) {
	faker.SetSeed(1201)
	value := faker.CurrencySymbol()
	t.Log(value)
	assert.Equal(t, "MK", value)
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
