package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	faker.SetSeed(102)
	var value bool
	value = faker.Bool()
	t.Log(value)
	assert.True(t, value)
	value = faker.Bool()
	t.Log(value)
	assert.False(t, value)
}

func TestPhoneNumber(t *testing.T) {
	faker.SetSeed(103)
	value := faker.PhoneNumber()
	t.Log(value)
	assert.Equal(t, "152.380.7298", value)
}

func TestUuid(t *testing.T) {
	faker.SetSeed(104)
	value := faker.Uuid()
	t.Log(value)
	assert.Equal(t, "40abb44c-895e-45b8-9f67-cc02a811744a", value)
}

func TestMiscBuild(t *testing.T) {
	faker.SetSeed(500)
	s := &struct {
		BoolField        bool `faker:"bool"`
		DefaultBoolField bool
		Field1           string `faker:"PhoneNumber"`
		Field2           string `faker:"Uuid"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.True(t, s.BoolField)
	assert.True(t, s.DefaultBoolField)
	assert.Equal(t, "1-740-515-9178", s.Field1)
	assert.Equal(t, "05e3b503-b43d-4d23-bca4-224b2e3e12f3", s.Field2)
}
