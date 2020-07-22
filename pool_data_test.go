package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestTimezonePool(t *testing.T) {
	factory.SetSeed(902)
	value, err := factory.GetData("timezone", "offset")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "7", value)
}
