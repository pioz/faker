package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestRegisterBuilder(t *testing.T) {
	err := faker.UnregisterBuilder("foo", "string")
	assert.NotNil(t, err)
	assert.Equal(t, "Builder not registered", err.Error())

	err = faker.RegisterBuilder("foo", "string", func(...string) (interface{}, error) {
		return "bar", nil
	})
	assert.Nil(t, err)

	err = faker.RegisterBuilder("foo", "string", func(...string) (interface{}, error) {
		return "bar", nil
	})
	assert.NotNil(t, err)
	assert.Equal(t, "Builder already registered", err.Error())

	err = faker.UnregisterBuilder("foo", "string")
	assert.Nil(t, err)
}
