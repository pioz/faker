package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestStringBuild(t *testing.T) {
	factory.SetSeed(920)
	s := &struct {
		Field1 string `factory:"StringWithSize(4)"`
		Field2 string `factory:"String"`
		Field3 string `factory:"DigitsWithSize(5)"`
		Field4 string `factory:"Digits(4)"`
		Field5 string `factory:"LettersWithSize(6)"`
		Field6 string `factory:"Letters"`
		Field7 string `factory:"Lexify(a?b?c)"`
		Field8 string `factory:"Numerify(???x)"`
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
}
