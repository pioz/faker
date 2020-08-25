package faker

import (
	"reflect"
)

// Slice will build a random slice of interface{} of length n. The elements of
// the slice are genereted by the function fn.
func Slice(size int, fn func() interface{}) []interface{} {
	slice := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, fn())
	}
	return slice
}

// Sample return a random element of slice. Sample panics if slice is not a
// slice, is nil or is empty.
func Sample(slice interface{}) interface{} {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic("faker.Sample param must be a slice")
	}
	sliceReflectValue := reflect.ValueOf(slice)
	if sliceReflectValue.IsNil() {
		panic("faker.Sample param is nil")
	}
	if sliceReflectValue.Len() == 0 {
		panic("faker.Sample param is empty")
	}
	i := IntInRange(0, sliceReflectValue.Len()-1)
	return sliceReflectValue.Index(i).Interface()
}
