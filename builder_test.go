package faker_test

import (
	"fmt"
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

func ExampleRegisterBuilder() {
	faker.SetSeed(1802)

	// Define a new builder
	builder := func(params ...string) (interface{}, error) {
		if len(params) > 0 && params[0] == "melee" {
			return faker.Pick("Barbarian", "Bard", "Fighter", "Monk", "Paladin", "Rogue"), nil
		}
		return faker.Pick("Cleric", "Druid", "Ranger", "Sorcerer", "Warlock", "Wizard"), nil
	}

	// Register a new builder named "dndClass" for string types
	err := faker.RegisterBuilder("dndClass", "string", builder)
	if err != nil {
		panic(err)
	}

	player := &struct {
		Class string `faker:"dndClass(melee)"`
		// other fields ...
	}{}

	// Build a struct with fake data
	faker.Build(&player)

	fmt.Println(player.Class)
	// Output: Paladin
}
