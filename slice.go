package faker

func Slice(size int, fn func() interface{}) []interface{} {
	slice := make([]interface{}, 0, size)
	for i := 0; i < size; i++ {
		slice = append(slice, fn())
	}
	return slice
}
