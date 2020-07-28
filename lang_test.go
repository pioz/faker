package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestLangName(t *testing.T) {
	factory.SetSeed(1100)
	value := factory.LangName()
	t.Log(value)
	assert.Equal(t, "Kirghiz", value)
}

func TestLangCode(t *testing.T) {
	factory.SetSeed(1101)
	value := factory.LangCode()
	t.Log(value)
	assert.Equal(t, "sk", value)
}

func TestLangBuild(t *testing.T) {
	factory.SetSeed(1110)
	s := &struct {
		Field1 string `factory:"LangName"`
		Field2 string `factory:"LangCode"`
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Samoan", s.Field1)
	assert.Equal(t, "fr", s.Field2)
}
