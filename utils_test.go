package faker_test

import (
	"testing"
	"time"

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
	assert.Equal(t, "invalid parameters", err.Error())
}

func TestParamsToMinMaxIntInvaliParams2(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange(1)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())
}

func TestParamsToMinMaxIntInvaliParams3(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange(a,b)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())
}

func TestParamsToMinMaxIntInvaliParams4(t *testing.T) {
	faker.SetSeed(700)
	s := &struct {
		IntField int `faker:"IntInRange(1,b)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.Atoi: parsing \"b\": invalid syntax", err.Error())
}

func TestParamsToMinMaxFloat64InvaliParams1(t *testing.T) {
	s := &struct {
		Field float64 `faker:"Float64InRange"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())
}

func TestParamsToMinMaxFloat64InvaliParams2(t *testing.T) {
	s := &struct {
		Field float64 `faker:"Float64InRange(1.0,b)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.ParseFloat: parsing \"b\": invalid syntax", err.Error())
}

func TestParamsToMinMaxFloat64InvaliParams3(t *testing.T) {
	s := &struct {
		Field float64 `faker:"Float64InRange(a,2)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.ParseFloat: parsing \"a\": invalid syntax", err.Error())
}

func TestParamsToMinMaxDurationInvalidParams1(t *testing.T) {
	s := &struct {
		Field time.Duration `faker:"DurationInRange"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())
}

func TestParamsToMinMaxDurationInvalidParams2(t *testing.T) {
	s := &struct {
		Field time.Duration `faker:"DurationInRange(2s,b)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: time: invalid duration \"b\"", err.Error())
}

func TestParamsToMinMaxDurationInvalidParams3(t *testing.T) {
	s := &struct {
		Field time.Duration `faker:"DurationInRange(a,2ms)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: time: invalid duration \"a\"", err.Error())
}

func TestParamsToIntInvalidParams1(t *testing.T) {
	s := &struct {
		Field string `faker:"StringWithSize"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())
}

func TestParamsToIntInvalidParams2(t *testing.T) {
	s := &struct {
		Field string `faker:"StringWithSize(a)"`
	}{}
	err := faker.Build(&s)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())
}
