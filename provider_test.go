package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestRegisterProvider(t *testing.T) {
	err := faker.UnregisterProvider("foo", "string")
	assert.NotNil(t, err)
	assert.Equal(t, "Provider not registered", err.Error())

	err = faker.RegisterProvider("foo", "string", func(...string) (interface{}, error) {
		return "bar", nil
	})
	assert.Nil(t, err)

	err = faker.RegisterProvider("foo", "string", func(...string) (interface{}, error) {
		return "bar", nil
	})
	assert.NotNil(t, err)
	assert.Equal(t, "Provider already registered", err.Error())

	err = faker.UnregisterProvider("foo", "string")
	assert.Nil(t, err)
}
