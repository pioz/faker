package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyName(t *testing.T) {
	factory.SetSeed(1200)
	value := factory.CurrencyName()
	t.Log(value)
	assert.Equal(t, "Sierra Leonean Leone", value)
}

func TestCurrencyCode(t *testing.T) {
	factory.SetSeed(1201)
	value := factory.CurrencyCode()
	t.Log(value)
	assert.Equal(t, "XOF", value)
}

func TestCurrencySymbol(t *testing.T) {
	factory.SetSeed(1201)
	value := factory.CurrencySymbol()
	t.Log(value)
	assert.Equal(t, "MK", value)
}

func TestCurrencyBuild(t *testing.T) {
	factory.SetSeed(1210)
	s := &struct {
		Field1 string `factory:"CurrencyName"`
		Field2 string `factory:"CurrencyCode"`
		Field3 string `factory:"CurrencySymbol"`
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Bermudian Dollar", s.Field1)
	assert.Equal(t, "XCD", s.Field2)
	assert.Equal(t, "K", s.Field3)
}
