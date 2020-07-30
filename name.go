package factory

import (
	"strings"
)

func MaleFirstName() string {
	value, err := GetData("name", "male_first_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func FemaleFirstName() string {
	value, err := GetData("name", "female_first_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func NeutralFirstName() string {
	value, err := GetData("name", "neutral_first_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func FirstName() string {
	kind := IntInRange(0, 2)
	switch kind {
	case 0:
		return MaleFirstName()
	case 1:
		return FemaleFirstName()
	default:
		return NeutralFirstName()
	}
}

func LastName() string {
	value, err := GetData("name", "last_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func NamePrefix() string {
	value, err := GetData("name", "prefix")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func NameSuffix() string {
	value, err := GetData("name", "suffix")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func FullName() string {
	templates := []string{"{first_name} {last_name}", "{prefix} {first_name} {last_name}", "{first_name} {suffix} {last_name}", "{prefix} {first_name} {suffix} {last_name}"}
	i := IntInRange(0, len(templates)*2)
	if i > len(templates)-1 {
		i = 0
	}
	name := templates[i]
	name = strings.Replace(name, "{first_name}", FirstName(), -1)
	name = strings.Replace(name, "{last_name}", LastName(), -1)
	name = strings.Replace(name, "{prefix}", NamePrefix(), -1)
	name = strings.Replace(name, "{suffix}", NameSuffix(), -1)
	return name
}

func NameInitials() string {
	return strings.ToUpper(Lexify("??"))
}

// Provider functions

func maleFirstNameProvider(params ...string) (interface{}, error) {
	return MaleFirstName(), nil
}

func femaleFirstName(params ...string) (interface{}, error) {
	return FemaleFirstName(), nil
}

func neutralFirstNameProvider(params ...string) (interface{}, error) {
	return NeutralFirstName(), nil
}

func firstNameProvider(params ...string) (interface{}, error) {
	return FirstName(), nil
}

func lastNameProvider(params ...string) (interface{}, error) {
	return LastName(), nil
}

func namePrefixProvider(params ...string) (interface{}, error) {
	return NamePrefix(), nil
}

func nameSuffixProvider(params ...string) (interface{}, error) {
	return NameSuffix(), nil
}

func fullNameProvider(params ...string) (interface{}, error) {
	return FullName(), nil
}

func nameInitialsProvider(params ...string) (interface{}, error) {
	return NameInitials(), nil
}
