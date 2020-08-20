package faker

import (
	"errors"
	"fmt"
	"strings"
)

type builderFunc func(...string) (interface{}, error)

func builderKey(builderName, builderType string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", builderName, builderType))
}

// RegisterBuilder register a new builder. A builder is a variadic function
// that will take an arbitrary number of strings as arguments and return an
// interface or an error. This function (builder) can be used to generate fake
// data. builderName is the name of the builder, builderType is the type of
// the interface returned by the builder.
func RegisterBuilder(builderName, builderType string, fn builderFunc) error {
	key := builderKey(builderName, builderType)
	if _, ok := builders[key]; ok {
		return errors.New("builder already registered")
	}
	builders[key] = fn
	return nil
}

// UnregisterBuilder unregister/remove a builder.
func UnregisterBuilder(builderName, builderType string) error {
	key := builderKey(builderName, builderType)
	if _, ok := builders[key]; !ok {
		return errors.New("builder not registered")
	}
	delete(builders, key)
	return nil
}
