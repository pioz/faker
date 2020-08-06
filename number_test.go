package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func TestIntInRange(t *testing.T) {
	faker.SetSeed(3)

	value := faker.IntInRange(10, 20)
	t.Log(value)
	assert.Equal(t, 16, value)

	value = faker.IntInRange(-20, -10)
	t.Log(value)
	assert.Equal(t, -14, value)

	value = faker.IntInRange(-20, -30)
	assert.Equal(t, -20, value)
}

func TestInt(t *testing.T) {
	faker.SetSeed(4)
	value := faker.Int()
	t.Log(value)
	assert.Equal(t, 477987042, value)
}

func TestInt64InRange(t *testing.T) {
	faker.SetSeed(5)

	value := faker.Int64InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int64(10), value)

	value = faker.Int64InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int64(-19), value)

	value = faker.Int64InRange(-20, -30)
	assert.Equal(t, int64(-20), value)
}

func TestInt64(t *testing.T) {
	faker.SetSeed(6)
	value := faker.Int64()
	t.Log(value)
	assert.Equal(t, int64(-5917743806733054187), value)
}

func TestInt32InRange(t *testing.T) {
	faker.SetSeed(7)

	value := faker.Int32InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int32(15), value)

	value = faker.Int32InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int32(-16), value)

	value = faker.Int32InRange(-20, -30)
	assert.Equal(t, int32(-20), value)
}

func TestInt32(t *testing.T) {
	faker.SetSeed(8)
	value := faker.Int32()
	t.Log(value)
	assert.Equal(t, int32(-417194768), value)
}

func TestInt16InRange(t *testing.T) {
	faker.SetSeed(9)

	value := faker.Int16InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int16(14), value)

	value = faker.Int16InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int16(-16), value)

	value = faker.Int16InRange(-20, -30)
	assert.Equal(t, int16(-20), value)
}

func TestInt16(t *testing.T) {
	faker.SetSeed(10)
	value := faker.Int16()
	t.Log(value)
	assert.Equal(t, int16(-24938), value)
}

func TestInt8InRange(t *testing.T) {
	faker.SetSeed(11)

	value := faker.Int8InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int8(15), value)

	value = faker.Int8InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int8(-14), value)

	value = faker.Int8InRange(-20, -30)
	assert.Equal(t, int8(-20), value)
}

func TestInt8(t *testing.T) {
	faker.SetSeed(12)
	value := faker.Int8()
	t.Log(value)
	assert.Equal(t, int8(-101), value)
}

func TestUintInRange(t *testing.T) {
	faker.SetSeed(13)

	value := faker.UintInRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint(15), value)

	value = faker.UintInRange(20, 10)
	assert.Equal(t, uint(20), value)
}

func TestUint(t *testing.T) {
	faker.SetSeed(14)
	value := faker.Uint()
	t.Log(value)
	assert.Equal(t, uint(2486533097), value)
}

func TestUint64InRange(t *testing.T) {
	faker.SetSeed(15)

	value := faker.Uint64InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint64(13), value)

	value = faker.Uint64InRange(20, 10)
	assert.Equal(t, uint64(20), value)
}

func TestUint64(t *testing.T) {
	faker.SetSeed(16)
	value := faker.Uint64()
	t.Log(value)
	assert.Equal(t, uint64(16676020418646319060), value)
}

func TestUint32InRange(t *testing.T) {
	faker.SetSeed(17)

	value := faker.Uint32InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint32(16), value)

	value = faker.Uint32InRange(20, 10)
	assert.Equal(t, uint32(20), value)
}

func TestUint32(t *testing.T) {
	faker.SetSeed(18)
	value := faker.Uint32()
	t.Log(value)
	assert.Equal(t, uint32(1185777406), value)
}

func TestUint16InRange(t *testing.T) {
	faker.SetSeed(19)

	value := faker.Uint16InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint16(13), value)

	value = faker.Uint16InRange(20, 10)
	assert.Equal(t, uint16(20), value)
}

func TestUint16(t *testing.T) {
	faker.SetSeed(20)
	value := faker.Uint16()
	t.Log(value)
	assert.Equal(t, uint16(29810), value)
}

func TestUint8InRange(t *testing.T) {
	faker.SetSeed(21)

	value := faker.Uint8InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint8(18), value)

	value = faker.Uint8InRange(20, 10)
	assert.Equal(t, uint8(20), value)
}

func TestUint8(t *testing.T) {
	faker.SetSeed(22)

	value := faker.Uint8()
	t.Log(value)
	assert.Equal(t, uint8(231), value)
}

func TestFloat64InRange(t *testing.T) {
	faker.SetSeed(23)

	value := faker.Float64InRange(10, 20)
	t.Log(value)
	assert.Equal(t, float64(18.120965248489753), value)

	value = faker.Float64InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, float64(-14.513652135502497), value)

	value = faker.Float64InRange(-20, -30)
	assert.Equal(t, float64(-20), value)
}

func TestFloat64(t *testing.T) {
	faker.SetSeed(24)

	value := faker.Float64()
	t.Log(value)
	assert.Equal(t, float64(6.696671980874496e+307), value)
}

func TestFloat32InRange(t *testing.T) {
	faker.SetSeed(25)

	value := faker.Float32InRange(10, 20)
	t.Log(value)
	assert.Equal(t, float32(18.961363), value)

	value = faker.Float32InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, float32(-15.832848), value)

	value = faker.Float32InRange(-20, -30)
	assert.Equal(t, float32(-20), value)
}

func TestFloat32(t *testing.T) {
	faker.SetSeed(26)
	value := faker.Float32()
	t.Log(value)
	assert.Equal(t, float32(1.5426473e+38), value)
}

func TestNumberBuild(t *testing.T) {
	faker.SetSeed(30)
	s := &struct {
		IntInRangeField     int `faker:"intinrange(1,10)"`
		IntField            int `faker:"int"`
		DefaultIntField     int
		Int64InRangeField   int64 `faker:"int64inrange(1,10)"`
		Int64Field          int64 `faker:"int64"`
		DefaultInt64Field   int64
		Int32InRangeField   int32 `faker:"int32inrange(1,10)"`
		Int32Field          int32 `faker:"int32"`
		DefaultInt32Field   int32
		Int16InRangeField   int16 `faker:"int16inrange(1,10)"`
		Int16Field          int16 `faker:"int16"`
		DefaultInt16Field   int16
		Int8InRangeField    int8 `faker:"int8inrange(1,10)"`
		Int8Field           int8 `faker:"int8"`
		DefaultInt8Field    int8
		UintInRangeField    uint `faker:"uintinrange(1,10)"`
		UintField           uint `faker:"uint"`
		DefaultUintField    uint
		Uint64InRangeField  uint64 `faker:"uint64inrange(1,10)"`
		Uint64Field         uint64 `faker:"uint64"`
		DefaultUint64Field  uint64
		Uint32InRangeField  uint32 `faker:"uint32inrange(1,10)"`
		Uint32Field         uint32 `faker:"uint32"`
		DefaultUint32Field  uint32
		Uint16InRangeField  uint16 `faker:"uint16inrange(1,10)"`
		Uint16Field         uint16 `faker:"uint16"`
		DefaultUint16Field  uint16
		Uint8InRangeField   uint8 `faker:"uint8inrange(1,10)"`
		Uint8Field          uint8 `faker:"uint8"`
		DefaultUint8Field   uint8
		Float64InRangeField float64 `faker:"Float64InRange(1,10.6)"`
		Float64Field        float64 `faker:"float64"`
		DefaultFloat64Field float64
		Float32InRangeField float32 `faker:"Float32InRange(1,10.6)"`
		Float32Field        float32 `faker:"float32"`
		DefaultFloat32Field float32
	}{}
	err := faker.Build(&s)
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
