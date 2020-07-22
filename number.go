package factory

import (
	"math"
)

func IntInRange(min, max int) int {
	if min >= max {
		return min
	}
	return random.Intn(max-min+1) + min
}

func Int() int {
	return IntInRange(math.MinInt32, math.MaxInt32)
}

func Int64InRange(min, max int64) int64 {
	if min >= max {
		return min
	}

	return random.Int63n(max-min) + min
}

func Int64() int64 {
	return random.Int63n(math.MaxInt64) + math.MinInt64
}

func Int32InRange(min, max int32) int32 {
	return int32(Int64InRange(int64(min), int64(max)))
}

func Int32() int32 {
	return Int32InRange(math.MinInt32, math.MaxInt32)
}

func Int16InRange(min, max int16) int16 {
	return int16(Int64InRange(int64(min), int64(max)))
}

func Int16() int16 {
	return Int16InRange(math.MinInt16, math.MaxInt16)
}

func Int8InRange(min, max int8) int8 {
	return int8(Int64InRange(int64(min), int64(max)))
}

func Int8() int8 {
	return Int8InRange(math.MinInt8, math.MaxInt8)
}

func UintInRange(min, max uint) uint {
	if min >= max {
		return min
	}
	return uint(random.Intn(int(max)-int(min)+1) + int(min))
}

func Uint() uint {
	return uint(IntInRange(0, math.MaxUint32))
}

func Uint64InRange(min, max uint64) uint64 {
	if min >= max {
		return min
	}
	return uint64(random.Int63n(int64(max)-int64(min)) + int64(min))
}

func Uint64() uint64 {
	return Uint64InRange(0, math.MaxInt64) + Uint64InRange(0, math.MaxInt64)
}

func Uint32InRange(min, max uint32) uint32 {
	return uint32(Uint64InRange(uint64(min), uint64(max)))
}

func Uint32() uint32 {
	return Uint32InRange(0, math.MaxUint32)
}

func Uint16InRange(min, max uint16) uint16 {
	return uint16(Uint64InRange(uint64(min), uint64(max)))
}

func Uint16() uint16 {
	return Uint16InRange(0, math.MaxUint16)
}

func Uint8InRange(min, max uint8) uint8 {
	return uint8(Uint64InRange(uint64(min), uint64(max)))
}

func Uint8() uint8 {
	return Uint8InRange(0, math.MaxUint8)
}

func Float64InRange(min, max float64) float64 {
	if min >= max {
		return min
	}
	return random.Float64()*(max-min) + min
}

func Float64() float64 {
	return Float64InRange(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

func Float32InRange(min, max float32) float32 {
	if min >= max {
		return min
	}
	return random.Float32()*(max-min) + min
}

func Float32() float32 {
	return Float32InRange(math.SmallestNonzeroFloat32, math.MaxFloat32)
}
