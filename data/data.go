package data

type Pool []interface{}
type PoolGroup map[string]Pool
type PoolData map[string]PoolGroup

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
