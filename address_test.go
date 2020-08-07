package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleAddressCity() {
	faker.SetSeed(1300)
	fmt.Println(faker.AddressCity())
	// Output: Ntoroko
}

func ExampleAddressState() {
	faker.SetSeed(1301)
	fmt.Println(faker.AddressState())
	// Output: Indiana
}

func ExampleAddressStateCode() {
	faker.SetSeed(1302)
	fmt.Println(faker.AddressStateCode())
	// Output: DC
}

func ExampleAddressStreetName() {
	faker.SetSeed(1303)
	fmt.Println(faker.AddressStreetName())
	// Output: Hopton Street
}

func ExampleAddressStreetNumber() {
	faker.SetSeed(1304)
	fmt.Println(faker.AddressStreetNumber())
	// Output: 81-680
}

func ExampleAddressSecondaryAddress() {
	faker.SetSeed(1305)
	fmt.Println(faker.AddressSecondaryAddress())
	// Output: Suite 208
}

func ExampleAddressZip() {
	faker.SetSeed(1306)
	fmt.Println(faker.AddressZip())
	// Output: 36168
}

func ExampleAddressFull() {
	faker.SetSeed(1307)
	fmt.Println(faker.AddressFull())
	// Output: John Snow
	// Apt. 248
	// 943 Wager Street
	// Berezniki PR 52209
	// Saudi Arabia
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
