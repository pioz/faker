package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestParamsToMinMaxIntInvaliParams1(t *testing.T) {
	factory.SetSeed(700)
	s := &struct {
		IntField int `factory:"IntInRange"`
	}{}
	err := factory.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters", err.Error())
}

func TestParamsToMinMaxIntInvaliParams2(t *testing.T) {
	factory.SetSeed(700)
	s := &struct {
		IntField int `factory:"IntInRange(1)"`
	}{}
	err := factory.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters", err.Error())
}

func TestParamsToMinMaxIntInvaliParams3(t *testing.T) {
	factory.SetSeed(700)
	s := &struct {
		IntField int `factory:"IntInRange(a,b)"`
	}{}
	err := factory.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())
}

func TestParamsToMinMaxIntInvaliParams4(t *testing.T) {
	factory.SetSeed(700)
	s := &struct {
		IntField int `factory:"IntInRange(1,b)"`
	}{}
	err := factory.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters: strconv.Atoi: parsing \"b\": invalid syntax", err.Error())
}
