package faker

import (
	"math"
)

// IntInRange will build a random int between min and max included.
func IntInRange(min, max int) int {
	if min >= max {
		return min
	}
	return random.Intn(max-min+1) + min
}

// Int will build a random int.
func Int() int {
	return IntInRange(math.MinInt32, math.MaxInt32)
}

// Int64InRange will build a random int64 between min and max included.
func Int64InRange(min, max int64) int64 {
	if min >= max {
		return min
	}

	return random.Int63n(max-min) + min
}

// Int64 will build a random int64.
func Int64() int64 {
	return random.Int63n(math.MaxInt64) + math.MinInt64
}

// Int32InRange will build a random int32 between min and max included.
func Int32InRange(min, max int32) int32 {
	return int32(Int64InRange(int64(min), int64(max)))
}

// Int32 will build a random int32.
func Int32() int32 {
	return Int32InRange(math.MinInt32, math.MaxInt32)
}

// Int16InRange will build a random int16 between min and max included.
func Int16InRange(min, max int16) int16 {
	return int16(Int64InRange(int64(min), int64(max)))
}

// Int16 will build a random int16.
func Int16() int16 {
	return Int16InRange(math.MinInt16, math.MaxInt16)
}

// Int8InRange will build a random int8 between min and max included.
func Int8InRange(min, max int8) int8 {
	return int8(Int64InRange(int64(min), int64(max)))
}

// Int8 will build a random int8.
func Int8() int8 {
	return Int8InRange(math.MinInt8, math.MaxInt8)
}

// UintInRange will build a random uint between min and max included.
func UintInRange(min, max uint) uint {
	if min >= max {
		return min
	}
	return uint(random.Intn(int(max)-int(min)+1) + int(min))
}

// Uint will build a random uint.
func Uint() uint {
	return uint(IntInRange(0, math.MaxUint32))
}

// Uint64InRange will build a random uint64 between min and max included.
func Uint64InRange(min, max uint64) uint64 {
	if min >= max {
		return min
	}
	return uint64(random.Int63n(int64(max)-int64(min)) + int64(min))
}

// Uint64 will build a random uint64.
func Uint64() uint64 {
	return Uint64InRange(0, math.MaxInt64) + Uint64InRange(0, math.MaxInt64)
}

// Uint32InRange will build a random uint32 between min and max included.
func Uint32InRange(min, max uint32) uint32 {
	return uint32(Uint64InRange(uint64(min), uint64(max)))
}

// Uint32 will build a random uint32.
func Uint32() uint32 {
	return Uint32InRange(0, math.MaxUint32)
}

// Uint16InRange will build a random uint16 between min and max included.
func Uint16InRange(min, max uint16) uint16 {
	return uint16(Uint64InRange(uint64(min), uint64(max)))
}

// Uint16 will build a random uint16.
func Uint16() uint16 {
	return Uint16InRange(0, math.MaxUint16)
}

// Uint8InRange will build a random uint8 between min and max included.
func Uint8InRange(min, max uint8) uint8 {
	return uint8(Uint64InRange(uint64(min), uint64(max)))
}

// Uint8 will build a random uint8.
func Uint8() uint8 {
	return Uint8InRange(0, math.MaxUint8)
}

// Float64InRange will build a random float64 between min and max included.
func Float64InRange(min, max float64) float64 {
	if min >= max {
		return min
	}
	return random.Float64()*(max-min) + min
}

// Float64 will build a random float64.
func Float64() float64 {
	return Float64InRange(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Float32InRange will build a random float32 between min and max included.
func Float32InRange(min, max float32) float32 {
	if min >= max {
		return min
	}
	return random.Float32()*(max-min) + min
}

// Float32 will build a random float32.
func Float32() float32 {
	return Float32InRange(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Builder functions

func intInRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return IntInRange(min, max), nil
}

func intBuilder(params ...string) (interface{}, error) {
	return Int(), nil
}

func int64InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int64InRange(int64(min), int64(max)), nil
}

func int64Builder(params ...string) (interface{}, error) {
	return Int64(), nil
}

func int32InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int32InRange(int32(min), int32(max)), nil
}

func int32Builder(params ...string) (interface{}, error) {
	return Int32(), nil
}

func int16InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int16InRange(int16(min), int16(max)), nil
}

func int16Builder(params ...string) (interface{}, error) {
	return Int16(), nil
}

func int8InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Int8InRange(int8(min), int8(max)), nil
}

func int8Builder(params ...string) (interface{}, error) {
	return Int8(), nil
}

func uintInRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return UintInRange(uint(min), uint(max)), nil
}

func uintBuilder(params ...string) (interface{}, error) {
	return Uint(), nil
}

func uint64InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint64InRange(uint64(min), uint64(max)), nil
}

func uint64Builder(params ...string) (interface{}, error) {
	return Uint64(), nil
}

func uint32InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint32InRange(uint32(min), uint32(max)), nil
}

func uint32Builder(params ...string) (interface{}, error) {
	return Uint32(), nil
}

func uint16InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint16InRange(uint16(min), uint16(max)), nil
}

func uint16Builder(params ...string) (interface{}, error) {
	return Uint16(), nil
}

func uint8InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxInt(params...)
	if err != nil {
		return nil, err
	}
	return Uint8InRange(uint8(min), uint8(max)), nil
}

func uint8Builder(params ...string) (interface{}, error) {
	return Uint8(), nil
}

func float64InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxFloat64(params...)
	if err != nil {
		return nil, err
	}
	return Float64InRange(min, max), nil
}

func float64Builder(params ...string) (interface{}, error) {
	return Float64(), nil
}

func float32InRangeBuilder(params ...string) (interface{}, error) {
	min, max, err := paramsToMinMaxFloat64(params...)
	if err != nil {
		return nil, err
	}
	return Float32InRange(float32(min), float32(max)), nil
}

func float32Builder(params ...string) (interface{}, error) {
	return Float32(), nil
}
