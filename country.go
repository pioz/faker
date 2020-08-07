package faker

// CountryName will build a random country name string.
func CountryName() string {
	value, err := GetData("country", "name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// CountryAlpha2 will build a random 2 characters country code string.
func CountryAlpha2() string {
	value, err := GetData("country", "alpha2")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// CountryAlpha3 will build a random 3 characters country code string.
func CountryAlpha3() string {
	value, err := GetData("country", "alpha3")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// CountryNationality will build a random nationality string.
func CountryNationality() string {
	value, err := GetData("country", "nationality")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// CountryFlag will build a random emoji flag string.
func CountryFlag() string {
	value, err := GetData("country", "flag")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Builder functions

func countryNameBuilder(params ...string) (interface{}, error) {
	return CountryName(), nil
}

func countryAlpha2Builder(params ...string) (interface{}, error) {
	return CountryAlpha2(), nil
}

func countryAlpha3Builder(params ...string) (interface{}, error) {
	return CountryAlpha3(), nil
}

func countryNationalityBuilder(params ...string) (interface{}, error) {
	return CountryNationality(), nil
}

func countryFlagBuilder(params ...string) (interface{}, error) {
	return CountryFlag(), nil
}
