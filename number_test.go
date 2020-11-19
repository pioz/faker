package faker_test

import (
	"fmt"
	"testing"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/assert"
)

func ExampleIntInRange() {
	faker.SetSeed(3)
	fmt.Println(faker.IntInRange(10, 20))
	fmt.Println(faker.IntInRange(-20, -10))
	fmt.Println(faker.IntInRange(-20, -30))
	// Output: 16
	// -14
	// -20
}

func ExampleInt() {
	faker.SetSeed(4)
	fmt.Println(faker.Int())
	// Output: 477987042
}

func ExampleInt64InRange() {
	faker.SetSeed(5)
	fmt.Println(faker.Int64InRange(10, 20))
	fmt.Println(faker.Int64InRange(-20, -10))
	fmt.Println(faker.Int64InRange(-20, -30))
	// Output: 10
	// -19
	// -20
}

func ExampleInt64() {
	faker.SetSeed(6)
	fmt.Println(faker.Int64())
	// Output: -5917743806733054187
}

func ExampleInt32InRange() {
	faker.SetSeed(7)
	fmt.Println(faker.Int32InRange(10, 20))
	fmt.Println(faker.Int32InRange(-20, -10))
	fmt.Println(faker.Int32InRange(-20, -30))
	// Output: 15
	// -16
	// -20
}

func ExampleInt32() {
	faker.SetSeed(8)
	fmt.Println(faker.Int32())
	// Output: -417194768
}

func ExampleInt16InRange() {
	faker.SetSeed(9)
	fmt.Println(faker.Int16InRange(10, 20))
	fmt.Println(faker.Int16InRange(-20, -10))
	fmt.Println(faker.Int16InRange(-20, -30))
	// Output: 14
	// -16
	// -20
}

func ExampleInt16() {
	faker.SetSeed(10)
	fmt.Println(faker.Int16())
	// Output: -24938
}

func ExampleInt8InRange() {
	faker.SetSeed(11)
	fmt.Println(faker.Int8InRange(10, 20))
	fmt.Println(faker.Int8InRange(-20, -10))
	fmt.Println(faker.Int8InRange(-20, -30))
	// Output: 15
	// -14
	// -20
}

func ExampleInt8() {
	faker.SetSeed(12)
	fmt.Println(faker.Int8())
	// Output: -101
}

func ExampleUintInRange() {
	faker.SetSeed(13)
	fmt.Println(faker.UintInRange(10, 20))
	fmt.Println(faker.UintInRange(20, 10))
	// Output: 15
	// 20
}

func ExampleUint() {
	faker.SetSeed(14)
	fmt.Println(faker.Uint())
	// Output: 2486533097
}

func ExampleUint64InRange() {
	faker.SetSeed(15)
	fmt.Println(faker.Uint64InRange(10, 20))
	fmt.Println(faker.Uint64InRange(20, 10))
	// Output: 13
	// 20
}

func ExampleUint64() {
	faker.SetSeed(16)
	fmt.Println(faker.Uint64())
	// Output: 16676020418646319060
}

func ExampleUint32InRange() {
	faker.SetSeed(17)
	fmt.Println(faker.Uint32InRange(10, 20))
	fmt.Println(faker.Uint32InRange(20, 10))
	// Output: 16
	// 20
}

func ExampleUint32() {
	faker.SetSeed(18)
	fmt.Println(faker.Uint32())
	// Output: 1185777406
}

func ExampleUint16InRange() {
	faker.SetSeed(19)
	fmt.Println(faker.Uint16InRange(10, 20))
	fmt.Println(faker.Uint16InRange(20, 10))
	// Output: 13
	// 20
}

func ExampleUint16() {
	faker.SetSeed(20)
	fmt.Println(faker.Uint16())
	// Output: 29810
}

func ExampleUint8InRange() {
	faker.SetSeed(21)
	fmt.Println(faker.Uint8InRange(10, 20))
	fmt.Println(faker.Uint8InRange(20, 10))
	// Output: 18
	// 20
}

func ExampleUint8() {
	faker.SetSeed(22)
	fmt.Println(faker.Uint8())
	// Output: 231
}

func ExampleFloat64InRange() {
	faker.SetSeed(23)
	fmt.Println(faker.Float64InRange(10, 20))
	fmt.Println(faker.Float64InRange(-20, -10))
	fmt.Println(faker.Float64InRange(-20, -30))
	// Output: 18.120965248489753
	// -14.513652135502497
	// -20
}

func ExampleFloat64() {
	faker.SetSeed(24)
	fmt.Println(faker.Float64())
	// Output: 6.696671980874496e+307
}

func ExampleFloat32InRange() {
	faker.SetSeed(25)
	fmt.Println(faker.Float32InRange(10, 20))
	fmt.Println(faker.Float32InRange(-20, -10))
	fmt.Println(faker.Float32InRange(-20, -30))
	// Output: 18.961363
	// -15.832848
	// -20
}

func ExampleFloat32() {
	faker.SetSeed(26)
	fmt.Println(faker.Float32())
	// Output: 1.5426473e+38
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

func TestNumberInvalidParams(t *testing.T) {
	faker.SetSeed(31)
	s1 := &struct {
		Field int64 `faker:"Int64InRange(a)"`
	}{}
	err := faker.Build(&s1)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s2 := &struct {
		Field int32 `faker:"Int32InRange(1,a)"`
	}{}
	err = faker.Build(&s2)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters: strconv.Atoi: parsing \"a\": invalid syntax", err.Error())

	s3 := &struct {
		Field int16 `faker:"Int16InRange(a)"`
	}{}
	err = faker.Build(&s3)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s4 := &struct {
		Field int8 `faker:"Int8InRange(a)"`
	}{}
	err = faker.Build(&s4)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s5 := &struct {
		Field uint64 `faker:"Uint64InRange(a)"`
	}{}
	err = faker.Build(&s5)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s6 := &struct {
		Field uint32 `faker:"Uint32InRange(a)"`
	}{}
	err = faker.Build(&s6)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s7 := &struct {
		Field uint16 `faker:"Uint16InRange(a)"`
	}{}
	err = faker.Build(&s7)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s8 := &struct {
		Field uint8 `faker:"Uint8InRange(a)"`
	}{}
	err = faker.Build(&s8)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s9 := &struct {
		Field float64 `faker:"float64InRange(a)"`
	}{}
	err = faker.Build(&s9)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())

	s10 := &struct {
		Field float32 `faker:"float32InRange(a)"`
	}{}
	err = faker.Build(&s10)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid parameters", err.Error())
}
