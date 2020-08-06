package faker

import (
	"fmt"
)

func Username() string {
	value, err := GetData("internet", "username")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

func Domain() string {
	domainSuffix, err := GetData("internet", "domain_suffix")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s.%s", Username(), domainSuffix)
}

func Email() string {
	return fmt.Sprintf("%s@%s", Username(), Domain())
}

func FreeEmail() string {
	domain, err := GetData("internet", "free_email")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s@%s", Username(), domain)
}

func SafeEmail() string {
	domain, err := GetData("internet", "safe_email")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s@%s", Username(), domain)
}

func Slug() string {
	return Parameterize(Sentence())
}

func Url() string {
	protocol, err := GetData("internet", "protocol")
	if err != nil {
		panic(err)
	}
	www := ""
	if Bool() {
		www = "www."
	}
	return fmt.Sprintf("%s://%s%s/%s", protocol, www, Domain(), Slug())
}

// Builder functions

func usernameBuilder(params ...string) (interface{}, error) {
	return Username(), nil
}

func domainBuilder(params ...string) (interface{}, error) {
	return Domain(), nil
}

func emailBuilder(params ...string) (interface{}, error) {
	return Email(), nil
}

func freeEmailBuilder(params ...string) (interface{}, error) {
	return FreeEmail(), nil
}

func safeEmailBuilder(params ...string) (interface{}, error) {
	return SafeEmail(), nil
}

func slugBuilder(params ...string) (interface{}, error) {
	return Slug(), nil
}

func urlBuilder(params ...string) (interface{}, error) {
	return Url(), nil
}
