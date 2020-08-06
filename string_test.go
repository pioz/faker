package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestStringWithSize(t *testing.T) {
	faker.SetSeed(900)
	assert.Equal(t, "zVCVi", faker.StringWithSize(5))
}

func TestString(t *testing.T) {
	faker.SetSeed(901)
	assert.Equal(t, "PMIWvi6i", faker.String())
}

func TestDigitsWithSize(t *testing.T) {
	faker.SetSeed(902)
	assert.Equal(t, "083232", faker.DigitsWithSize(6))
}

func TestDigits(t *testing.T) {
	faker.SetSeed(903)
	assert.Equal(t, "615512400386075514115741266153386748264703852380812377508844034317824285597708278706114561792707006830894923584342024491056728416855686581847469966327751223766784279332851122716701382579799798697760692758485844", faker.Digits())
}

func TestLettersWithSize(t *testing.T) {
	faker.SetSeed(904)
	assert.Equal(t, "ZHCbqwV", faker.LettersWithSize(7))
}

func TestLetters(t *testing.T) {
	faker.SetSeed(905)
	assert.Equal(t, "HZEoOvFsmElJQnsbdXbkRhVXJupACokXppdhpWO", faker.Letters())
}

func TestLexify(t *testing.T) {
	faker.SetSeed(906)
	assert.Equal(t, "abxcZhdZ", faker.Lexify("ab?c??d?"))
}

func TestNumerify(t *testing.T) {
	faker.SetSeed(907)
	assert.Equal(t, "ab5c30d754", faker.Numerify("ab?c??d???"))
}

func TestParameterize(t *testing.T) {
	faker.SetSeed(908)
	assert.Equal(t, "the-amazing-zanz-153", faker.Parameterize("The Amazing Zanzò 153 "))
}

func TestPick(t *testing.T) {
	faker.SetSeed(909)
	assert.Equal(t, "dog", faker.Pick("cat", "dog", "mouse", "lion", "bear"))
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
	assert.Equal(t, "TeFS5Odn1cPhpV2u9WBocBKPvRUzZx1mLtxH9W1d", s.Field2)
	assert.Equal(t, "21470", s.Field3)
	assert.Equal(t, "91616148171577033030", s.Field4)
	assert.Equal(t, "hRMljr", s.Field5)
	assert.Equal(t, "krLqaIpuNJIyqpHnxwwlsVlyYmlIQJKpsekLGbykxAdtZOuUCoJwkQAMKjeEDYGYRgzhGDXhZkYUCasPAALIcPDsbkhMVsxJIYPytgJRFBCMWlEtukCKyKCYokIhiPYYofRkbpCPxtrAorUYCqRGejEFAGTikhUzvAgPuDHNonPnCbUxRJmxrqZoOjNEIQWuMhmgCljwIYQRymoOGiaANnzxPhQNIESQRxWsVtmLxcqIjCTTBbF", s.Field6)
	assert.Equal(t, "aXbGc", s.Field7)
	assert.Equal(t, "729x", s.Field8)
	assert.Equal(t, "go-is-the-best-programming-language", s.Field9)
	assert.Equal(t, "lion", s.Field10)
}
