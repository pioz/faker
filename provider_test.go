package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestRegisterProvider(t *testing.T) {
	err := factory.UnregisterProvider("foo", "string")
	assert.NotNil(t, err)
	assert.Equal(t, "Provider not registered", err.Error())

	err = factory.RegisterProvider("foo", "string", func(...string) (interface{}, error) {
		return "bar", nil
	})
	assert.Nil(t, err)

	err = factory.RegisterProvider("foo", "string", func(...string) (interface{}, error) {
		return "bar", nil
	})
	assert.NotNil(t, err)
	assert.Equal(t, "Provider already registered", err.Error())

	err = factory.UnregisterProvider("foo", "string")
	assert.Nil(t, err)
}
