package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestMiscBuild(t *testing.T) {
	factory.SetSeed(500)
	s := &struct {
		BoolField        bool `factory:"bool"`
		DefaultBoolField bool
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)
	assert.True(t, s.BoolField)
	assert.True(t, s.DefaultBoolField)
}
