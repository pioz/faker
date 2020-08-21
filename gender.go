package faker

var genderData = PoolGroup{
	"types":              {"Male", "Female", "Agender", "Androgyne", "Androgynous", "Bigender", "Cis", "Cisgender", "Cis Female", "Cis Male", "Cis Man", "Cis Woman", "Cisgender Female", "Cisgender Male", "Cisgender Man", "Cisgender Woman", "Female to Male", "FTM", "Gender Fluid", "Gender Nonconforming", "Gender Questioning", "Gender Variant", "Genderqueer", "Intersex", "Male to Female", "MTF", "Neither", "Neutrois", "Non-binary", "Other", "Pangender", "Trans", "Trans Female", "Trans Male", "Trans Man", "Trans Person", "Trans Woman", "Transfeminine", "Transgender", "Transgender Female", "Transgender Male", "Transgender Man", "Transgender Person", "Transgender Woman", "Transmasculine", "Transsexual", "Transsexual Female", "Transsexual Male", "Transsexual Man", "Transsexual Person", "Transsexual Woman", "Two-Spiri"},
	"binary_types":       {"Male", "Female"},
	"short_binary_types": {"f", "m"},
}

// Gender will build a random gender string.
func Gender() string {
	value, _ := GetData("gender", "types")
	return value.(string)
}

// BinaryGender will build a random binary gender string (Male or Female).
func BinaryGender() string {
	value, _ := GetData("gender", "binary_types")
	return value.(string)
}

// ShortBinaryGender will build a random short binary gender string (m or f).
func ShortBinaryGender() string {
	value, _ := GetData("gender", "short_binary_types")
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
