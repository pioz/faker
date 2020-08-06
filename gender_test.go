package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	faker.SetSeed(1500)
	value := faker.Gender()
	t.Log(value)
	assert.Equal(t, "Cisgender Male", value)
}

func TestBinaryGender(t *testing.T) {
	faker.SetSeed(1501)
	value := faker.BinaryGender()
	t.Log(value)
	assert.Equal(t, "Male", value)
}

func TestShortBinaryGender(t *testing.T) {
	faker.SetSeed(1502)
	value := faker.ShortBinaryGender()
	t.Log(value)
	assert.Equal(t, "m", value)
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
