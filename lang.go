package faker

func LangName() string {
	value, err := GetData("lang", "name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func LangCode() string {
	value, err := GetData("lang", "code")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Provider functions

func langNameProvider(params ...string) (interface{}, error) {
	return LangName(), nil
}

func langCodeProvider(params ...string) (interface{}, error) {
	return LangCode(), nil
}
