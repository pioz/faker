package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestAddressCity(t *testing.T) {
	faker.SetSeed(1300)
	value := faker.AddressCity()
	t.Log(value)
	assert.Equal(t, "Ntoroko", value)
}

func TestAddressState(t *testing.T) {
	faker.SetSeed(1301)
	value := faker.AddressState()
	t.Log(value)
	assert.Equal(t, "Indiana", value)
}

func TestAddressStateCode(t *testing.T) {
	faker.SetSeed(1302)
	value := faker.AddressStateCode()
	t.Log(value)
	assert.Equal(t, "DC", value)
}

func TestAddressStreetName(t *testing.T) {
	faker.SetSeed(1303)
	value := faker.AddressStreetName()
	t.Log(value)
	assert.Equal(t, "Hopton Street", value)
}

func TestAddressStreetNumber(t *testing.T) {
	faker.SetSeed(1304)
	value := faker.AddressStreetNumber()
	t.Log(value)
	assert.Equal(t, "81-680", value)
}

func TestAddressSecondaryAddress(t *testing.T) {
	faker.SetSeed(1305)
	value := faker.AddressSecondaryAddress()
	t.Log(value)
	assert.Equal(t, "Suite 208", value)
}

func TestAddressZip(t *testing.T) {
	faker.SetSeed(1306)
	value := faker.AddressZip()
	t.Log(value)
	assert.Equal(t, "36168", value)
}

func TestAddressFull(t *testing.T) {
	faker.SetSeed(1307)
	value := faker.AddressFull()
	t.Log(value)
	assert.Equal(t, "John Snow\nApt. 248\n943 Wager Street\nBerezniki PR 52209\nSaudi Arabia", value)
}

func TestAddressBuild(t *testing.T) {
	faker.SetSeed(1008)
	s := &struct {
		Field1 string `faker:"AddressCity"`
		Field2 string `faker:"AddressState"`
		Field3 string `faker:"AddressStateCode"`
		Field4 string `faker:"AddressStreetName"`
		Field5 string `faker:"AddressStreetNumber"`
		Field6 string `faker:"AddressSecondaryAddress"`
		Field7 string `faker:"AddressZip"`
		Field8 string `faker:"AddressFull"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Pryor Creek", s.Field1)
	assert.Equal(t, "Ohio", s.Field2)
	assert.Equal(t, "AP", s.Field3)
	assert.Equal(t, "Shuttle Street", s.Field4)
	assert.Equal(t, "72", s.Field5)
	assert.Equal(t, "Suite 516", s.Field6)
	assert.Equal(t, "05094", s.Field7)
	assert.Equal(t, "John Snow\nApt. 485\n632 Windsor Walk\nVidalia AL 71829-6715\nAndorra", s.Field8)
}
