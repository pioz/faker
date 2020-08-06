package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestParamsToMinMaxIntInvaliParams1(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters", err.Error())
}

func TestParamsToMinMaxIntInvaliParams2(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange(1)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters", err.Error())
}

func TestParamsToMinMaxIntInvaliParams3(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange(a,b)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())
}

func TestParamsToMinMaxIntInvaliParams4(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange(1,b)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid parameters: strconv.Atoi: parsing \"b\": invalid syntax", err.Error())
}
