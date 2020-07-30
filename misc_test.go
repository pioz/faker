package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	factory.SetSeed(102)
	var value bool
	value = factory.Bool()
	t.Log(value)
	assert.True(t, value)
	value = factory.Bool()
	t.Log(value)
	assert.False(t, value)
}

func TestPhoneNumber(t *testing.T) {
	factory.SetSeed(103)
	value := factory.PhoneNumber()
	t.Log(value)
	assert.Equal(t, "152.380.7298", value)
}

func TestUuid(t *testing.T) {
	factory.SetSeed(104)
	value := factory.Uuid()
	t.Log(value)
	assert.Equal(t, "40abb44c-895e-45b8-9f67-cc02a811744a", value)
}

func TestMiscBuild(t *testing.T) {
	factory.SetSeed(500)
	s := &struct {
		BoolField        bool `factory:"bool"`
		DefaultBoolField bool
		Field1           string `factory:"PhoneNumber"`
		Field2           string `factory:"Uuid"`
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.True(t, s.BoolField)
	assert.True(t, s.DefaultBoolField)
	assert.Equal(t, "1-740-515-9178", s.Field1)
	assert.Equal(t, "05e3b503-b43d-4d23-bca4-224b2e3e12f3", s.Field2)
}
