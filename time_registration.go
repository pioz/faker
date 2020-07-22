package factory

func init() {
	RegisterProvider("DurationInRange", "time.Duration", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxDuration(params...)
		if err != nil {
			return nil, err
		}
		return DurationInRange(min, max), nil
	})

	RegisterProvider("Duration", "time.Duration", func(params ...string) (interface{}, error) {
		return Duration(), nil
	})

	RegisterProvider("", "time.Duration", func(params ...string) (interface{}, error) {
		return Duration(), nil
	})

	RegisterProvider("Time", "time.Time", func(params ...string) (interface{}, error) {
		return Time(), nil
	})

	RegisterProvider("", "time.Time", func(params ...string) (interface{}, error) {
		return Time(), nil
	})
}
