package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestUsername(t *testing.T) {
	faker.SetSeed(1700)
	value := faker.Username()
	t.Log(value)
	assert.Equal(t, "polychasium", value)
}

func TestDomain(t *testing.T) {
	faker.SetSeed(1701)
	value := faker.Domain()
	t.Log(value)
	assert.Equal(t, "lorusso.name", value)
}

func TestEmail(t *testing.T) {
	faker.SetSeed(1702)
	value := faker.Email()
	t.Log(value)
	assert.Equal(t, "homocercal@fulmer.net", value)
}

func TestFreeEmail(t *testing.T) {
	faker.SetSeed(1703)
	value := faker.FreeEmail()
	t.Log(value)
	assert.Equal(t, "atlas@gmail.com", value)
}

func TestSafeEmail(t *testing.T) {
	faker.SetSeed(1704)
	value := faker.SafeEmail()
	t.Log(value)
	assert.Equal(t, "disbelieve@example.com", value)
}

func TestSlug(t *testing.T) {
	faker.SetSeed(1705)
	value := faker.Slug()
	t.Log(value)
	assert.Equal(t, "a-reliable-seal-s-bee-comes-with-it-the-thought-that-the-adventurous-giraffe-is-an-alligator", value)
}

func TestUrl(t *testing.T) {
	faker.SetSeed(1706)
	value := faker.Url()
	t.Log(value)
	assert.Equal(t, "https://www.mcmillon.info/this-could-be-or-perhaps-their-alligator-was-in-this-moment-an-eager-spider", value)
}

func TestInternetBuild(t *testing.T) {
	faker.SetSeed(1720)
	s := &struct {
		Field1 string `faker:"Username"`
		Field2 string `faker:"Domain"`
		Field3 string `faker:"Email"`
		Field4 string `faker:"SafeEmail"`
		Field5 string `faker:"FreeEmail"`
		Field6 string `faker:"Slug"`
		Field7 string `faker:"Url"`
	}{}
	err := faker.Build(&s)
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
