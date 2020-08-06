package faker

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

// Builder functions

func currencyNameBuilder(params ...string) (interface{}, error) {
	return CurrencyName(), nil
}

func currencyCodeBuilder(params ...string) (interface{}, error) {
	return CurrencyCode(), nil
}

func currencySymbolBuilder(params ...string) (interface{}, error) {
	return CurrencySymbol(), nil
}
