package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleLangName() {
	faker.SetSeed(1100)
	fmt.Println(faker.LangName())
	// Output: Kirghiz
}

func ExampleLangCode() {
	faker.SetSeed(1101)
	fmt.Println(faker.LangCode())
	// Output: sk
}

func TestLangBuild(t *testing.T) {
	faker.SetSeed(1110)
	s := &struct {
		Field1 string `faker:"LangName"`
		Field2 string `faker:"LangCode"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.Equal(t, "Samoan", s.Field1)
	assert.Equal(t, "fr", s.Field2)
}
