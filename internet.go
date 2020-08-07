package faker

import (
	"fmt"
)

// Username will build a random username string.
func Username() string {
	value, err := GetData("internet", "username")
	if err != nil {
		panic(err)
	}
	return value.(string)
}

// Domain will build a random domain string.
func Domain() string {
	domainSuffix, err := GetData("internet", "domain_suffix")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s.%s", Username(), domainSuffix)
}

// Email will build a random email address string.
func Email() string {
	return fmt.Sprintf("%s@%s", Username(), Domain())
}

// FreeEmail will build a random free email address string (gmail, hotmail, yahoo...).
func FreeEmail() string {
	domain, err := GetData("internet", "free_email")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s@%s", Username(), domain)
}

// SafeEmail will build a random email address string whose domain is always example.com.
func SafeEmail() string {
	domain, err := GetData("internet", "safe_email")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s@%s", Username(), domain)
}

// Slug will build a random slug string.
func Slug() string {
	return Parameterize(Sentence())
}

// URL will build a random URL string.
func URL() string {
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
	return URL(), nil
}
