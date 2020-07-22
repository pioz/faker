package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestNumberBuild(t *testing.T) {
	factory.SetSeed(30)
	s := &struct {
		IntInRangeField     int `factory:"intinrange(1,10)"`
		IntField            int `factory:"int"`
		DefaultIntField     int
		Int64InRangeField   int64 `factory:"int64inrange(1,10)"`
		Int64Field          int64 `factory:"int64"`
		DefaultInt64Field   int64
		Int32InRangeField   int32 `factory:"int32inrange(1,10)"`
		Int32Field          int32 `factory:"int32"`
		DefaultInt32Field   int32
		Int16InRangeField   int16 `factory:"int16inrange(1,10)"`
		Int16Field          int16 `factory:"int16"`
		DefaultInt16Field   int16
		Int8InRangeField    int8 `factory:"int8inrange(1,10)"`
		Int8Field           int8 `factory:"int8"`
		DefaultInt8Field    int8
		UintInRangeField    uint `factory:"uintinrange(1,10)"`
		UintField           uint `factory:"uint"`
		DefaultUintField    uint
		Uint64InRangeField  uint64 `factory:"uint64inrange(1,10)"`
		Uint64Field         uint64 `factory:"uint64"`
		DefaultUint64Field  uint64
		Uint32InRangeField  uint32 `factory:"uint32inrange(1,10)"`
		Uint32Field         uint32 `factory:"uint32"`
		DefaultUint32Field  uint32
		Uint16InRangeField  uint16 `factory:"uint16inrange(1,10)"`
		Uint16Field         uint16 `factory:"uint16"`
		DefaultUint16Field  uint16
		Uint8InRangeField   uint8 `factory:"uint8inrange(1,10)"`
		Uint8Field          uint8 `factory:"uint8"`
		DefaultUint8Field   uint8
		Float64InRangeField float64 `factory:"Float64InRange(1,10.6)"`
		Float64Field        float64 `factory:"float64"`
		DefaultFloat64Field float64
		Float32InRangeField float32 `factory:"Float32InRange(1,10.6)"`
		Float32Field        float32 `factory:"float32"`
		DefaultFloat32Field float32
	}{}
	err := factory.Build(&s)
	assert.Nil(t, err)
	t.Log(s)

	assert.Equal(t, 9, s.IntInRangeField)
	assert.Equal(t, -1588912014, s.IntField)
	assert.Equal(t, -1491768569, s.DefaultIntField)

	assert.Equal(t, int64(5), s.Int64InRangeField)
	assert.Equal(t, int64(-5089836721814089834), s.Int64Field)
	assert.Equal(t, int64(-3837974256891962760), s.DefaultInt64Field)

	assert.Equal(t, int32(1), s.Int32InRangeField)
	assert.Equal(t, int32(-1416661755), s.Int32Field)
	assert.Equal(t, int32(-1243937080), s.DefaultInt32Field)

	assert.Equal(t, int16(5), s.Int16InRangeField)
	assert.Equal(t, int16(18620), s.Int16Field)
	assert.Equal(t, int16(-3431), s.DefaultInt16Field)

	assert.Equal(t, int8(1), s.Int8InRangeField)
	assert.Equal(t, int8(84), s.Int8Field)
	assert.Equal(t, int8(-17), s.DefaultInt8Field)

	assert.Equal(t, uint(8), s.UintInRangeField)
	assert.Equal(t, uint(1869826397), s.UintField)
	assert.Equal(t, uint(498610181), s.DefaultUintField)

	assert.Equal(t, uint64(3), s.Uint64InRangeField)
	assert.Equal(t, uint64(8118072143219141702), s.Uint64Field)
	assert.Equal(t, uint64(10753327993520087958), s.DefaultUint64Field)

	assert.Equal(t, uint32(1), s.Uint32InRangeField)
	assert.Equal(t, uint32(2486160442), s.Uint32Field)
	assert.Equal(t, uint32(2189451667), s.DefaultUint32Field)

	assert.Equal(t, uint16(6), s.Uint16InRangeField)
	assert.Equal(t, uint16(38236), s.Uint16Field)
	assert.Equal(t, uint16(55644), s.DefaultUint16Field)

	assert.Equal(t, uint8(1), s.Uint8InRangeField)
	assert.Equal(t, uint8(98), s.Uint8Field)
	assert.Equal(t, uint8(68), s.DefaultUint8Field)

	assert.Equal(t, float64(9.535379483774388), s.Float64InRangeField)
	assert.Equal(t, float64(1.6249093496264468e+308), s.Float64Field)
	assert.Equal(t, float64(2.2516638145714746e+307), s.DefaultFloat64Field)

	assert.Equal(t, float32(7.8864875), s.Float32InRangeField)
	assert.Equal(t, float32(3.0873776e+38), s.Float32Field)
	assert.Equal(t, float32(3.3773583e+38), s.DefaultFloat32Field)
}
