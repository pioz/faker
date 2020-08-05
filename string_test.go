package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestStringWithSize(t *testing.T) {
	factory.SetSeed(900)
	assert.Equal(t, "zVCVi", factory.StringWithSize(5))
}

func TestString(t *testing.T) {
	factory.SetSeed(901)
	assert.Equal(t, "PMIWvi6i", factory.String())
}

func TestDigitsWithSize(t *testing.T) {
	factory.SetSeed(902)
	assert.Equal(t, "083232", factory.DigitsWithSize(6))
}

func TestDigits(t *testing.T) {
	factory.SetSeed(903)
	assert.Equal(t, "615512400386075514115741266153386748264703852380812377508844034317824285597708278706114561792707006830894923584342024491056728416855686581847469966327751223766784279332851122716701382579799798697760692758485844", factory.Digits())
}

func TestLettersWithSize(t *testing.T) {
	factory.SetSeed(904)
	assert.Equal(t, "ZHCbqwV", factory.LettersWithSize(7))
}

func TestLetters(t *testing.T) {
	factory.SetSeed(905)
	assert.Equal(t, "HZEoOvFsmElJQnsbdXbkRhVXJupACokXppdhpWO", factory.Letters())
}

func TestLexify(t *testing.T) {
	factory.SetSeed(906)
	assert.Equal(t, "abxcZhdZ", factory.Lexify("ab?c??d?"))
}

func TestNumerify(t *testing.T) {
	factory.SetSeed(907)
	assert.Equal(t, "ab5c30d754", factory.Numerify("ab?c??d???"))
}

func TestParameterize(t *testing.T) {
	factory.SetSeed(908)
	assert.Equal(t, "the-amazing-zanz-153", factory.Parameterize("The Amazing Zanzò 153 "))
}

func TestPick(t *testing.T) {
	factory.SetSeed(909)
	assert.Equal(t, "dog", factory.Pick("cat", "dog", "mouse", "lion", "bear"))
}

func TestStringBuild(t *testing.T) {
	factory.SetSeed(920)
	s := &struct {
		Field1  string `factory:"StringWithSize(4)"`
		Field2  string `factory:"String"`
		Field3  string `factory:"DigitsWithSize(5)"`
		Field4  string `factory:"Digits(4)"`
		Field5  string `factory:"LettersWithSize(6)"`
		Field6  string `factory:"Letters"`
		Field7  string `factory:"Lexify(a?b?c)"`
		Field8  string `factory:"Numerify(???x)"`
		Field9  string `factory:"Parameterize(Go is the best programming language!)"`
		Field10 string `factory:"Pick(cat,dog,mouse,lion,bear)"`
	}{}
	err := factory.Build(&s)
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
