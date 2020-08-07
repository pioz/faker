package faker

// LangName will build a random language name string.
func LangName() string {
	value, err := GetData("lang", "name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// LangCode will build a random language code string.
func LangCode() string {
	value, err := GetData("lang", "code")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Builder functions

func langNameBuilder(params ...string) (interface{}, error) {
	return LangName(), nil
}

func langCodeBuilder(params ...string) (interface{}, error) {
	return LangCode(), nil
}
