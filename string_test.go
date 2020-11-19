package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleStringWithSize() {
	faker.SetSeed(900)
	fmt.Println(faker.StringWithSize(5))
	// Output: zVCVi
}

func ExampleString() {
	faker.SetSeed(901)
	fmt.Println(faker.String())
	// Output: PMIWvi6i
}

func ExampleDigitsWithSize() {
	faker.SetSeed(902)
	fmt.Println(faker.DigitsWithSize(6))
	// Output: 083232
}

func ExampleDigits() {
	faker.SetSeed(903)
	fmt.Println(faker.Digits())
	// Output: 615512
}

func ExampleLettersWithSize() {
	faker.SetSeed(904)
	fmt.Println(faker.LettersWithSize(7))
	// Output: ZHCbqwV
}

func ExampleLetters() {
	faker.SetSeed(905)
	fmt.Println(faker.Letters())
	// Output: HZ
}

func ExampleLexify() {
	faker.SetSeed(906)
	fmt.Println(faker.Lexify("ab?c??d?"))
	// Output: abxcZhdZ
}

func ExampleNumerify() {
	faker.SetSeed(907)
	fmt.Println(faker.Numerify("ab?c??d???"))
	// Output: ab5c30d754
}

func ExampleParameterize() {
	faker.SetSeed(908)
	fmt.Println(faker.Parameterize("The Amazing Zanzò 153 "))
	// Output: the-amazing-zanz-153
}

func ExamplePick() {
	faker.SetSeed(909)
	fmt.Println(faker.Pick("cat", "dog", "mouse", "lion", "bear"))
	// Output: dog
}

func TestPickWithNoArgs(t *testing.T) {
	faker.SetSeed(910)
	assert.Equal(t, "", faker.Pick())
}

func TestStringBuild(t *testing.T) {
	faker.SetSeed(920)
	s := &struct {
		Field1  string `faker:"StringWithSize(4)"`
		Field2  string `faker:"String"`
		Field3  string `faker:"DigitsWithSize(5)"`
		Field4  string `faker:"Digits(4)"`
		Field5  string `faker:"LettersWithSize(6)"`
		Field6  string `faker:"Letters"`
		Field7  string `faker:"Lexify(a?b?c)"`
		Field8  string `faker:"Numerify(???x)"`
		Field9  string `faker:"Parameterize(Go is the best programming language!)"`
		Field10 string `faker:"Pick(cat,dog,mouse,lion,bear)"`
	}{}
	err := faker.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, "1nkm", s.Field1)
	assert.Equal(t, "TeF", s.Field2)
	assert.Equal(t, "61017", s.Field3)
	assert.Equal(t, "43939689", s.Field4)
	assert.Equal(t, "odEMfI", s.Field5)
	assert.Equal(t, "VBilXp", s.Field6)
	assert.Equal(t, "abbqc", s.Field7)
	assert.Equal(t, "111x", s.Field8)
	assert.Equal(t, "go-is-the-best-programming-language", s.Field9)
	assert.Equal(t, "cat", s.Field10)
}

func TestStringInvalidParams(t *testing.T) {
	faker.SetSeed(921)
	s1 := &struct {
		Field string `faker:"DigitsWithSize(a)"`
	}{}
	err := faker.Build(&s1)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())

	s2 := &struct {
		Field string `faker:"LettersWithSize(a)"`
	}{}
	err = faker.Build(&s2)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())

	s3 := &struct {
		Field string `faker:"Lexify"`
	}{}
	err = faker.Build(&s3)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s4 := &struct {
		Field string `faker:"Numerify"`
	}{}
	err = faker.Build(&s4)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s5 := &struct {
		Field string `faker:"Parameterize"`
	}{}
	err = faker.Build(&s5)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())
}
