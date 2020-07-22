package factory

func init() {
	RegisterProvider("IntInRange", "int", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return IntInRange(min, max), nil
	})

	RegisterProvider("Int", "int", func(params ...string) (interface{}, error) {
		return Int(), nil
	})

	RegisterProvider("", "int", func(params ...string) (interface{}, error) {
		return Int(), nil
	})

	RegisterProvider("Int64InRange", "int64", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Int64InRange(int64(min), int64(max)), nil
	})

	RegisterProvider("Int64", "int64", func(params ...string) (interface{}, error) {
		return Int64(), nil
	})

	RegisterProvider("", "int64", func(params ...string) (interface{}, error) {
		return Int64(), nil
	})

	RegisterProvider("Int32InRange", "int32", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Int32InRange(int32(min), int32(max)), nil
	})

	RegisterProvider("Int32", "int32", func(params ...string) (interface{}, error) {
		return Int32(), nil
	})

	RegisterProvider("", "int32", func(params ...string) (interface{}, error) {
		return Int32(), nil
	})

	RegisterProvider("Int16InRange", "int16", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Int16InRange(int16(min), int16(max)), nil
	})

	RegisterProvider("Int16", "int16", func(params ...string) (interface{}, error) {
		return Int16(), nil
	})

	RegisterProvider("", "int16", func(params ...string) (interface{}, error) {
		return Int16(), nil
	})

	RegisterProvider("Int8InRange", "int8", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Int8InRange(int8(min), int8(max)), nil
	})

	RegisterProvider("Int8", "int8", func(params ...string) (interface{}, error) {
		return Int8(), nil
	})

	RegisterProvider("", "int8", func(params ...string) (interface{}, error) {
		return Int8(), nil
	})

	RegisterProvider("UintInRange", "uint", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return UintInRange(uint(min), uint(max)), nil
	})

	RegisterProvider("Uint", "uint", func(params ...string) (interface{}, error) {
		return Uint(), nil
	})

	RegisterProvider("", "uint", func(params ...string) (interface{}, error) {
		return Uint(), nil
	})

	RegisterProvider("Uint64InRange", "uint64", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Uint64InRange(uint64(min), uint64(max)), nil
	})

	RegisterProvider("Uint64", "uint64", func(params ...string) (interface{}, error) {
		return Uint64(), nil
	})

	RegisterProvider("", "uint64", func(params ...string) (interface{}, error) {
		return Uint64(), nil
	})

	RegisterProvider("Uint32InRange", "uint32", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Uint32InRange(uint32(min), uint32(max)), nil
	})

	RegisterProvider("Uint32", "uint32", func(params ...string) (interface{}, error) {
		return Uint32(), nil
	})

	RegisterProvider("", "uint32", func(params ...string) (interface{}, error) {
		return Uint32(), nil
	})

	RegisterProvider("Uint16InRange", "uint16", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Uint16InRange(uint16(min), uint16(max)), nil
	})

	RegisterProvider("Uint16", "uint16", func(params ...string) (interface{}, error) {
		return Uint16(), nil
	})

	RegisterProvider("", "uint16", func(params ...string) (interface{}, error) {
		return Uint16(), nil
	})

	RegisterProvider("Uint8InRange", "uint8", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxInt(params...)
		if err != nil {
			return nil, err
		}
		return Uint8InRange(uint8(min), uint8(max)), nil
	})

	RegisterProvider("Uint8", "uint8", func(params ...string) (interface{}, error) {
		return Uint8(), nil
	})

	RegisterProvider("", "uint8", func(params ...string) (interface{}, error) {
		return Uint8(), nil
	})

	RegisterProvider("Float64InRange", "float64", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxFloat64(params...)
		if err != nil {
			return nil, err
		}
		return Float64InRange(min, max), nil
	})

	RegisterProvider("Float64", "float64", func(params ...string) (interface{}, error) {
		return Float64(), nil
	})

	RegisterProvider("", "float64", func(params ...string) (interface{}, error) {
		return Float64(), nil
	})

	RegisterProvider("Float32InRange", "float32", func(params ...string) (interface{}, error) {
		min, max, err := paramsToMinMaxFloat64(params...)
		if err != nil {
			return nil, err
		}
		return Float32InRange(float32(min), float32(max)), nil
	})

	RegisterProvider("Float32", "float32", func(params ...string) (interface{}, error) {
		return Float32(), nil
	})

	RegisterProvider("", "float32", func(params ...string) (interface{}, error) {
		return Float32(), nil
	})
}
