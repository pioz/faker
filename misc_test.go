package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	factory.SetSeed(102)
	var value bool
	value = factory.Bool()
	t.Log(value)
	assert.True(t, value)
	value = factory.Bool()
	t.Log(value)
	assert.False(t, value)
}
