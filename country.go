package faker

func CountryName() string {
	value, err := GetData("country", "name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func CountryAlpha2() string {
	value, err := GetData("country", "alpha2")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func CountryAlpha3() string {
	value, err := GetData("country", "alpha3")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func CountryNationality() string {
	value, err := GetData("country", "nationality")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func CountryFlag() string {
	value, err := GetData("country", "flag")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Provider functions

func countryNameProvider(params ...string) (interface{}, error) {
	return CountryName(), nil
}

func countryAlpha2Provider(params ...string) (interface{}, error) {
	return CountryAlpha2(), nil
}

func countryAlpha3Provider(params ...string) (interface{}, error) {
	return CountryAlpha3(), nil
}

func countryNationalityProvider(params ...string) (interface{}, error) {
	return CountryNationality(), nil
}

func countryFlagProvider(params ...string) (interface{}, error) {
	return CountryFlag(), nil
}
