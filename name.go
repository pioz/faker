package faker

import (
	"strings"
)

// MaleFirstName will build a random male first name string.
func MaleFirstName() string {
	value, err := GetData("name", "male_first_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// FemaleFirstName will build a random female first name string.
func FemaleFirstName() string {
	value, err := GetData("name", "female_first_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// NeutralFirstName will build a random neutral first name string.
func NeutralFirstName() string {
	value, err := GetData("name", "neutral_first_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// FirstName will build a random first name string.
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

// LastName will build a random last name string.
func LastName() string {
	value, err := GetData("name", "last_name")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// NamePrefix will build a random name prefix string.
func NamePrefix() string {
	value, err := GetData("name", "prefix")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// NameSuffix will build a random name suffix string.
func NameSuffix() string {
	value, err := GetData("name", "suffix")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// FullName will build a random full name string.
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

// NameInitials will build a random name initials of 2 characters string.
func NameInitials() string {
	return strings.ToUpper(Lexify("??"))
}

// Builder functions

func maleFirstNameBuilder(params ...string) (interface{}, error) {
	return MaleFirstName(), nil
}

func femaleFirstName(params ...string) (interface{}, error) {
	return FemaleFirstName(), nil
}

func neutralFirstNameBuilder(params ...string) (interface{}, error) {
	return NeutralFirstName(), nil
}

func firstNameBuilder(params ...string) (interface{}, error) {
	return FirstName(), nil
}

func lastNameBuilder(params ...string) (interface{}, error) {
	return LastName(), nil
}

func namePrefixBuilder(params ...string) (interface{}, error) {
	return NamePrefix(), nil
}

func nameSuffixBuilder(params ...string) (interface{}, error) {
	return NameSuffix(), nil
}

func fullNameBuilder(params ...string) (interface{}, error) {
	return FullName(), nil
}

func nameInitialsBuilder(params ...string) (interface{}, error) {
	return NameInitials(), nil
}
