package faker

import (
	"fmt"
	"strconv"
	"time"
)

func paramsToMinMaxInt(params ...string) (int, int, error) {
	if len(params) != 2 {
		return 0, 0, parametersError(nil)
	}
	var min, max int
	var err error
	min, err = strconv.Atoi(params[0])
	if err != nil {
		return 0, 0, parametersError(err)
	}
	max, err = strconv.Atoi(params[1])
	if err != nil {
		return 0, 0, parametersError(err)
	}
	return min, max, nil
}

func paramsToMinMaxFloat64(params ...string) (float64, float64, error) {
	if len(params) != 2 {
		return 0.0, 0.0, parametersError(nil)
	}
	var min, max float64
	var err error
	min, err = strconv.ParseFloat(params[0], 64)
	if err != nil {
		return 0.0, 0.0, parametersError(err)
	}
	max, err = strconv.ParseFloat(params[1], 64)
	if err != nil {
		return 0.0, 0.0, parametersError(err)
	}
	return min, max, nil
}

func paramsToMinMaxDuration(params ...string) (time.Duration, time.Duration, error) {
	if len(params) != 2 {
		return 0, 0, parametersError(nil)
	}
	var min, max time.Duration
	var err error
	min, err = time.ParseDuration(params[0])
	if err != nil {
		return 0, 0, parametersError(err)
	}
	max, err = time.ParseDuration(params[1])
	if err != nil {
		return 0, 0, parametersError(err)
	}
	return min, max, nil
}

func paramsToInt(params ...string) (int, error) {
	if len(params) != 1 {
		return 0, parametersError(nil)
	}
	number, err := strconv.Atoi(params[0])
	if err != nil {
		return 0, parametersError(err)
	}
	return number, nil
}

func parametersError(err error) error {
	if err == nil {
		return fmt.Errorf("invalid parameters")
	}
	return fmt.Errorf("invalid parameters: %s", err.Error())
}
