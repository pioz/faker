package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestAddressCity(t *testing.T) {
	factory.SetSeed(1300)
	value := factory.AddressCity()
	t.Log(value)
	assert.Equal(t, "Ntoroko", value)
}

func TestAddressState(t *testing.T) {
	factory.SetSeed(1301)
	value := factory.AddressState()
	t.Log(value)
	assert.Equal(t, "Indiana", value)
}

func TestAddressStateCode(t *testing.T) {
	factory.SetSeed(1302)
	value := factory.AddressStateCode()
	t.Log(value)
	assert.Equal(t, "DC", value)
}

func TestAddressStreetName(t *testing.T) {
	factory.SetSeed(1303)
	value := factory.AddressStreetName()
	t.Log(value)
	assert.Equal(t, "Hopton Street", value)
}

func TestAddressStreetNumber(t *testing.T) {
	factory.SetSeed(1304)
	value := factory.AddressStreetNumber()
	t.Log(value)
	assert.Equal(t, "81-680", value)
}

func TestAddressSecondaryAddress(t *testing.T) {
	factory.SetSeed(1305)
	value := factory.AddressSecondaryAddress()
	t.Log(value)
	assert.Equal(t, "Suite 208", value)
}

func TestAddressZip(t *testing.T) {
	factory.SetSeed(1306)
	value := factory.AddressZip()
	t.Log(value)
	assert.Equal(t, "36168", value)
}

func TestAddressFull(t *testing.T) {
	factory.SetSeed(1307)
	value := factory.AddressFull()
	t.Log(value)
	assert.Equal(t, "John Snow\nApt. 248\n943 Wager Street\nBerezniki PR 52209\nSaudi Arabia", value)
}

func TestAddressBuild(t *testing.T) {
	factory.SetSeed(1008)
	s := &struct {
		Field1 string `factory:"AddressCity"`
		Field2 string `factory:"AddressState"`
		Field3 string `factory:"AddressStateCode"`
		Field4 string `factory:"AddressStreetName"`
		Field5 string `factory:"AddressStreetNumber"`
		Field6 string `factory:"AddressSecondaryAddress"`
		Field7 string `factory:"AddressZip"`
		Field8 string `factory:"AddressFull"`
	}{}
	err := factory.Build(&s)
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
