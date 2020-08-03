package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestUsername(t *testing.T) {
	factory.SetSeed(1700)
	value := factory.Username()
	t.Log(value)
	assert.Equal(t, "polychasium", value)
}

func TestDomain(t *testing.T) {
	factory.SetSeed(1701)
	value := factory.Domain()
	t.Log(value)
	assert.Equal(t, "lorusso.name", value)
}

func TestEmail(t *testing.T) {
	factory.SetSeed(1702)
	value := factory.Email()
	t.Log(value)
	assert.Equal(t, "homocercal@fulmer.net", value)
}

func TestFreeEmail(t *testing.T) {
	factory.SetSeed(1703)
	value := factory.FreeEmail()
	t.Log(value)
	assert.Equal(t, "atlas@gmail.com", value)
}

func TestSafeEmail(t *testing.T) {
	factory.SetSeed(1704)
	value := factory.SafeEmail()
	t.Log(value)
	assert.Equal(t, "disbelieve@example.com", value)
}

func TestSlug(t *testing.T) {
	factory.SetSeed(1705)
	value := factory.Slug()
	t.Log(value)
	assert.Equal(t, "a-reliable-seal-s-bee-comes-with-it-the-thought-that-the-adventurous-giraffe-is-an-alligator", value)
}

func TestUrl(t *testing.T) {
	factory.SetSeed(1706)
	value := factory.Url()
	t.Log(value)
	assert.Equal(t, "https://www.mcmillon.info/this-could-be-or-perhaps-their-alligator-was-in-this-moment-an-eager-spider", value)
}

func TestInternetBuild(t *testing.T) {
	factory.SetSeed(1720)
	s := &struct {
		Field1 string `factory:"Username"`
		Field2 string `factory:"Domain"`
		Field3 string `factory:"Email"`
		Field4 string `factory:"SafeEmail"`
		Field5 string `factory:"FreeEmail"`
		Field6 string `factory:"Slug"`
		Field7 string `factory:"Url"`
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "tight", s.Field1)
	assert.Equal(t, "extortionate.info", s.Field2)
	assert.Equal(t, "euchology@farris.info", s.Field3)
	assert.Equal(t, "andino@example.com", s.Field4)
	assert.Equal(t, "inflorescence@gmail.com", s.Field5)
	assert.Equal(t, "a-friendly-blackberry-is-an-eagle-of-the-mind", s.Field6)
	assert.Equal(t, "https://www.condescendence.info/nowhere-is-it-disputed-that-their-blueberry-was-in-this-moment-a-plausible-plum", s.Field7)
}
