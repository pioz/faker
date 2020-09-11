# Faker

![Go](https://github.com/pioz/faker/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/pioz/faker)](https://goreportcard.com/report/github.com/pioz/faker)
[![codecov](https://codecov.io/gh/pioz/faker/branch/master/graph/badge.svg)](https://codecov.io/gh/pioz/faker)
[![awesome-go](https://camo.githubusercontent.com/13c4e50d88df7178ae1882a203ed57b641674f94/68747470733a2f2f63646e2e7261776769742e636f6d2f73696e647265736f726875732f617765736f6d652f643733303566333864323966656437386661383536353265336136336531353464643865383832392f6d656469612f62616467652e737667)](https://github.com/avelino/awesome-go)
[![GoReference](https://pkg.go.dev/badge/mod/github.com/pioz/faker)](https://pkg.go.dev/github.com/pioz/faker)
<!-- [![GoDoc](https://godoc.org/github.com/pioz/faker?status.svg)](https://godoc.org/github.com/pioz/faker) -->

Random fake data and struct generator for Go.

* More than 100 generator functions
* Struct generator
* Unique data generator
* Builtin types support
* Easily customizable
* Zero dependencies
* Recursive infinite loop detector
* Benchmarks (coming soon)

## Installation

    go get github.com/pioz/faker

## Example

```go
faker.SetSeed(623)

fmt.Println(faker.Username())
fmt.Println(faker.String())
fmt.Println(faker.IntInRange(1, 10))
fmt.Println(faker.CurrencyName())
// Output: spicule
// gzaazJyRGt3jDnVh6ik7R9FO0AU1HcOzdOXbmNVBRQ5pq8n9tHf9B21PIFozLEzsoY4wILvZjTxSLQmD3UOAamDgVR411T3YHleDTgLuz90XSO3NFZm1AnaJiJamVRcNGD2zmi4qWkcjKF3E4JKgn1DiCeC3eSb5WELsw8XqRzlvJqG
// 10
// Myanmar Kyat
```

Please refer to the [godoc](https://godoc.org/github.com/pioz/faker) for all available functions and more.

## Struct Builder Example

```go
faker.SetSeed(622)

// Define a new builder
colorBuilder := func(params ...string) (interface{}, error) {
  return faker.Pick("Red", "Yellow", "Blue", "Black", "White"), nil
}

// Register a new builder named "color" for string type
err := faker.RegisterBuilder("color", "string", colorBuilder)
if err != nil {
  panic(err)
}

type Animal struct {
  Name  string `faker:"username"`
  Color string `faker:"color"` // Use custom color builder
}

type Person struct {
  FirstName string            `faker:"firstName"`         // Any available function case insensitive
  LastName  *string           `faker:"lastName"`          // Pointer are also supported
  Age       int               `faker:"intinrange(0,120)"` // Can call with parameters
  UUID      string            `faker:"uuid;unique"`       // Guarantees a unique value
  Number    int               `faker:"-"`                 // Skip this field
  Code      string            // No tag to use default builder for this field type
  Pet       Animal            // Recursively fill this struct
  Nicknames []string          `faker:"username;len=3"`          // Build an array of size 3 using faker.Username function
  Extra     map[string]string `faker:"stringWithSize(3);len=2"` // map are supported
}

p := Person{}
err = faker.Build(&p)
if err != nil {
  panic(err)
}
fmt.Println(p.FirstName)
fmt.Println(*p.LastName)
fmt.Println(p.Age)
fmt.Println(p.UUID)
fmt.Println(p.Number)
fmt.Println(p.Code)
fmt.Println(p.Pet.Name)
fmt.Println(p.Pet.Color)
fmt.Println(len(p.Nicknames))
fmt.Println(p.Nicknames[0])
fmt.Println(p.Nicknames[1])
fmt.Println(p.Nicknames[2])
fmt.Println(p.Extra)
// Output: Wilber
// Gutkowski
// 25
// ff8d6917-b920-46e6-b1be-dc2d48becfcb
// 0
// z
// honegger
// Red
// 3
// teagan
// polypeptide
// chinfest
// map[70w:3F6 gQS:isq]
```

## Factory

One of the nice things about Faker is that it can also be used as a factory
library. In fact when we call the `faker.Build` function if a value is not
zero then it is not modified, leaving the original value. This allows you to
create factory functions very easily:

```go
faker.SetSeed(623)

type User struct {
    Username string `faker:"username"`
    Email    string `faker:"email"`
    Country  string `faker:"CountryAlpha2"`
}

italianUserFactory := func() *User {
    u := &User{Country: "IT"}
    faker.Build(u)
    return u
}

italianUser := italianUserFactory()
fmt.Println(italianUser)
// Output: &{spicule hoag@ornamented.biz IT}
```

## Customization

A builder is a variadic function that will take an arbitrary number of strings
as arguments and return an interface or an error. This function (builder) can
be used to generate fake data and customize the behaviour of Faker. Here an example:

```go
faker.SetSeed(1802)

// Define a new builder
builder := func(params ...string) (interface{}, error) {
    if len(params) > 0 && params[0] == "melee" {
        return faker.Pick("Barbarian", "Bard", "Fighter", "Monk", "Paladin", "Rogue"), nil
    }
    return faker.Pick("Cleric", "Druid", "Ranger", "Sorcerer", "Warlock", "Wizard"), nil
}

// Register a new builder named "dndClass" for string type
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
```

## Contributing with new generator functions

If you want to contribute to faker with new generator functions please read this [wiki](https://github.com/pioz/faker/wiki/Contributing-with-new-generator-functions)!

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/pioz/faker/issues.

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).






