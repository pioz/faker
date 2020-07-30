package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestGender(t *testing.T) {
	factory.SetSeed(1500)
	value := factory.Gender()
	t.Log(value)
	assert.Equal(t, "Cisgender Male", value)
}

func TestBinaryGender(t *testing.T) {
	factory.SetSeed(1501)
	value := factory.BinaryGender()
	t.Log(value)
	assert.Equal(t, "Male", value)
}

func TestShortBinaryGender(t *testing.T) {
	factory.SetSeed(1502)
	value := factory.ShortBinaryGender()
	t.Log(value)
	assert.Equal(t, "m", value)
}

func TestGenderBuild(t *testing.T) {
	factory.SetSeed(1520)
	s := &struct {
		Field1 string `factory:"Gender"`
		Field2 string `factory:"BinaryGender"`
		Field3 string `factory:"ShortBinaryGender"`
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Cis", s.Field1)
	assert.Equal(t, "Female", s.Field2)
	assert.Equal(t, "m", s.Field3)
}
