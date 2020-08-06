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

// Provider functions

func genderProvider(params ...string) (interface{}, error) {
	return Gender(), nil
}

func binaryGenderProvider(params ...string) (interface{}, error) {
	return BinaryGender(), nil
}

func shortBinaryGenderProvider(params ...string) (interface{}, error) {
	return ShortBinaryGender(), nil
}
