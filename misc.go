package factory

func Bool() bool {
	return IntInRange(0, 1) == 0
}
