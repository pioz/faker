package data

type Pool []interface{}
type PoolGroup map[string]Pool
type PoolData map[string]PoolGroup

var DB = PoolData{
	"address":  Address,
	"country":  Country,
	"currency": Currency,
	"lang":     Lang,
	"timezone": Timezone,
}
