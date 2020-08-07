package data

// Pool type is a slice that contains fake data.
type Pool []interface{}

// PoolGroup type is a map that groups Pool types.
type PoolGroup map[string]Pool

// PoolData type is a map that groups PoolGroup types.
type PoolData map[string]PoolGroup

// DB is a map (PoolData) that contains fake data grupped in two levels of
// categories.
var DB = PoolData{
	"address":  Address,
	"country":  Country,
	"currency": Currency,
	"gender":   Gender,
	"internet": Internet,
	"lang":     Lang,
	"name":     Name,
	"sentence": Sentence,
	"timezone": Timezone,
}
