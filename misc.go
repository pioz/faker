package factory

func Bool() bool {
	return IntInRange(0, 1) == 0
}

// Provider functions

func boolFn(params ...string) (interface{}, error) {
	return Bool(), nil
}
