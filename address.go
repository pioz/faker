package faker

import (
	"fmt"
)

func AddressCity() string {
	value, err := GetData("address", "city")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func AddressState() string {
	value, err := GetData("address", "state")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func AddressStateCode() string {
	value, err := GetData("address", "state_code")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func AddressStreetName() string {
	value, err := GetData("address", "street")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func AddressStreetNumber() string {
	value, err := GetData("address", "number")
	if err != nil {
		panic(err)
	}
	return Numerify(value.(string))
}

func AddressSecondaryAddress() string {
	value, err := GetData("address", "secondary_address")
	if err != nil {
		panic(err)
	}
	return Numerify(value.(string))
}

func AddressZip() string {
	value, err := GetData("address", "zip")
	if err != nil {
		panic(err)
	}
	return Numerify(value.(string))
}

func AddressFull() string {
	return fmt.Sprintf("%s\n%s\n%s %s\n%s %s %s\n%s", "John Snow", AddressSecondaryAddress(), AddressStreetNumber(), AddressStreetName(), AddressCity(), AddressStateCode(), AddressZip(), CountryName())
}

// Builder functions

func addressCityBuilder(params ...string) (interface{}, error) {
	return AddressCity(), nil
}

func addressStateBuilder(params ...string) (interface{}, error) {
	return AddressState(), nil
}

func addressStateCodeBuilder(params ...string) (interface{}, error) {
	return AddressStateCode(), nil
}

func addressStreetNameBuilder(params ...string) (interface{}, error) {
	return AddressStreetName(), nil
}

func addressStreetNumberBuilder(params ...string) (interface{}, error) {
	return AddressStreetNumber(), nil
}

func addressSecondaryAddressBuilder(params ...string) (interface{}, error) {
	return AddressSecondaryAddress(), nil
}

func addressZipBuilder(params ...string) (interface{}, error) {
	return AddressZip(), nil
}

func addressFullBuilder(params ...string) (interface{}, error) {
	return AddressFull(), nil
}
