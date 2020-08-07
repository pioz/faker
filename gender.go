package faker

// Gender will build a random gender string.
func Gender() string {
	value, err := GetData("gender", "types")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// BinaryGender will build a random binary gender string (Male or Female).
func BinaryGender() string {
	value, err := GetData("gender", "binary_types")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// ShortBinaryGender will build a random short binary gender string (m or f).
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
