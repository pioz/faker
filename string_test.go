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
