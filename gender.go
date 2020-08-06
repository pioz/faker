package faker

func Gender() string {
	value, err := GetData("gender", "types")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func BinaryGender() string {
	value, err := GetData("gender", "binary_types")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func ShortBinaryGender() string {
	value, err := GetData("gender", "short_binary_types")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Builder functions

func genderBuilder(params ...string) (interface{}, error) {
	return Gender(), nil
}

func binaryGenderBuilder(params ...string) (interface{}, error) {
	return BinaryGender(), nil
}

func shortBinaryGenderBuilder(params ...string) (interface{}, error) {
	return ShortBinaryGender(), nil
}
