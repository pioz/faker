package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestLangName(t *testing.T) {
	faker.SetSeed(1100)
	value := faker.LangName()
	t.Log(value)
	assert.Equal(t, "Kirghiz", value)
}

func TestLangCode(t *testing.T) {
	faker.SetSeed(1101)
	value := faker.LangCode()
	t.Log(value)
	assert.Equal(t, "sk", value)
}

func TestLangBuild(t *testing.T) {
	faker.SetSeed(1110)
	s := &struct {
		Field1 string `faker:"LangName"`
		Field2 string `faker:"LangCode"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Samoan", s.Field1)
	assert.Equal(t, "fr", s.Field2)
}
