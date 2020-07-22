package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/stretchr/testify/assert"
)

func TestIntInRange(t *testing.T) {
	factory.SetSeed(3)

	value := factory.IntInRange(10, 20)
	t.Log(value)
	assert.Equal(t, 16, value)

	value = factory.IntInRange(-20, -10)
	t.Log(value)
	assert.Equal(t, -14, value)

	value = factory.IntInRange(-20, -30)
	assert.Equal(t, -20, value)
}

func TestInt(t *testing.T) {
	factory.SetSeed(4)
	value := factory.Int()
	t.Log(value)
	assert.Equal(t, 477987042, value)
}

func TestInt64InRange(t *testing.T) {
	factory.SetSeed(5)

	value := factory.Int64InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int64(10), value)

	value = factory.Int64InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int64(-19), value)

	value = factory.Int64InRange(-20, -30)
	assert.Equal(t, int64(-20), value)
}

func TestInt64(t *testing.T) {
	factory.SetSeed(6)
	value := factory.Int64()
	t.Log(value)
	assert.Equal(t, int64(-5917743806733054187), value)
}

func TestInt32InRange(t *testing.T) {
	factory.SetSeed(7)

	value := factory.Int32InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int32(15), value)

	value = factory.Int32InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int32(-16), value)

	value = factory.Int32InRange(-20, -30)
	assert.Equal(t, int32(-20), value)
}

func TestInt32(t *testing.T) {
	factory.SetSeed(8)
	value := factory.Int32()
	t.Log(value)
	assert.Equal(t, int32(-417194768), value)
}

func TestInt16InRange(t *testing.T) {
	factory.SetSeed(9)

	value := factory.Int16InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int16(14), value)

	value = factory.Int16InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int16(-16), value)

	value = factory.Int16InRange(-20, -30)
	assert.Equal(t, int16(-20), value)
}

func TestInt16(t *testing.T) {
	factory.SetSeed(10)
	value := factory.Int16()
	t.Log(value)
	assert.Equal(t, int16(-24938), value)
}

func TestInt8InRange(t *testing.T) {
	factory.SetSeed(11)

	value := factory.Int8InRange(10, 20)
	t.Log(value)
	assert.Equal(t, int8(15), value)

	value = factory.Int8InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, int8(-14), value)

	value = factory.Int8InRange(-20, -30)
	assert.Equal(t, int8(-20), value)
}

func TestInt8(t *testing.T) {
	factory.SetSeed(12)
	value := factory.Int8()
	t.Log(value)
	assert.Equal(t, int8(-101), value)
}

func TestUintInRange(t *testing.T) {
	factory.SetSeed(13)

	value := factory.UintInRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint(15), value)

	value = factory.UintInRange(20, 10)
	assert.Equal(t, uint(20), value)
}

func TestUint(t *testing.T) {
	factory.SetSeed(14)
	value := factory.Uint()
	t.Log(value)
	assert.Equal(t, uint(2486533097), value)
}

func TestUint64InRange(t *testing.T) {
	factory.SetSeed(15)

	value := factory.Uint64InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint64(13), value)

	value = factory.Uint64InRange(20, 10)
	assert.Equal(t, uint64(20), value)
}

func TestUint64(t *testing.T) {
	factory.SetSeed(16)
	value := factory.Uint64()
	t.Log(value)
	assert.Equal(t, uint64(16676020418646319060), value)
}

func TestUint32InRange(t *testing.T) {
	factory.SetSeed(17)

	value := factory.Uint32InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint32(16), value)

	value = factory.Uint32InRange(20, 10)
	assert.Equal(t, uint32(20), value)
}

func TestUint32(t *testing.T) {
	factory.SetSeed(18)
	value := factory.Uint32()
	t.Log(value)
	assert.Equal(t, uint32(1185777406), value)
}

func TestUint16InRange(t *testing.T) {
	factory.SetSeed(19)

	value := factory.Uint16InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint16(13), value)

	value = factory.Uint16InRange(20, 10)
	assert.Equal(t, uint16(20), value)
}

func TestUint16(t *testing.T) {
	factory.SetSeed(20)
	value := factory.Uint16()
	t.Log(value)
	assert.Equal(t, uint16(29810), value)
}

func TestUint8InRange(t *testing.T) {
	factory.SetSeed(21)

	value := factory.Uint8InRange(10, 20)
	t.Log(value)
	assert.Equal(t, uint8(18), value)

	value = factory.Uint8InRange(20, 10)
	assert.Equal(t, uint8(20), value)
}

func TestUint8(t *testing.T) {
	factory.SetSeed(22)

	value := factory.Uint8()
	t.Log(value)
	assert.Equal(t, uint8(231), value)
}

func TestFloat64InRange(t *testing.T) {
	factory.SetSeed(23)

	value := factory.Float64InRange(10, 20)
	t.Log(value)
	assert.Equal(t, float64(18.120965248489753), value)

	value = factory.Float64InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, float64(-14.513652135502497), value)

	value = factory.Float64InRange(-20, -30)
	assert.Equal(t, float64(-20), value)
}

func TestFloat64(t *testing.T) {
	factory.SetSeed(24)

	value := factory.Float64()
	t.Log(value)
	assert.Equal(t, float64(6.696671980874496e+307), value)
}

func TestFloat32InRange(t *testing.T) {
	factory.SetSeed(25)

	value := factory.Float32InRange(10, 20)
	t.Log(value)
	assert.Equal(t, float32(18.961363), value)

	value = factory.Float32InRange(-20, -10)
	t.Log(value)
	assert.Equal(t, float32(-15.832848), value)

	value = factory.Float32InRange(-20, -30)
	assert.Equal(t, float32(-20), value)
}

func TestFloat32(t *testing.T) {
	factory.SetSeed(26)
	value := factory.Float32()
	t.Log(value)
	assert.Equal(t, float32(1.5426473e+38), value)
}
