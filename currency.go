package factory

func CurrencyName() string {
	value, err := GetData("currency", "name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func CurrencyCode() string {
	value, err := GetData("currency", "code")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func CurrencySymbol() string {
	value, err := GetData("currency", "symbol")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Provider functions

func currencyNameProvider(params ...string) (interface{}, error) {
	return CurrencyName(), nil
}

func currencyCodeProvider(params ...string) (interface{}, error) {
	return CurrencyCode(), nil
}

func currencySymbolProvider(params ...string) (interface{}, error) {
	return CurrencySymbol(), nil
}
