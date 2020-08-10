package faker

// Slice will build a random slice of interface{} of length n. The elements of
// the slice are genereted by the function fn.
func Slice(size int, fn func() interface{}) []interface{} {
	slice := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, fn())
	}
	return slice
}
