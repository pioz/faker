package factory

func init() {
	RegisterProvider("StringWithSize", "string", func(params ...string) (interface{}, error) {
		size, err := paramsToInt(params...)
		if err != nil {
			return nil, err
		}
		return StringWithSize(size), nil
	})

	RegisterProvider("String", "string", func(params ...string) (interface{}, error) {
		return String(), nil
	})

	RegisterProvider("DigitsWithSize", "string", func(params ...string) (interface{}, error) {
		size, err := paramsToInt(params...)
		if err != nil {
			return nil, err
		}
		return DigitsWithSize(size), nil
	})

	RegisterProvider("Digits", "string", func(params ...string) (interface{}, error) {
		return Digits(), nil
	})

	RegisterProvider("LettersWithSize", "string", func(params ...string) (interface{}, error) {
		size, err := paramsToInt(params...)
		if err != nil {
			return nil, err
		}
		return LettersWithSize(size), nil
	})

	RegisterProvider("Letters", "string", func(params ...string) (interface{}, error) {
		return Letters(), nil
	})

	RegisterProvider("Lexify", "string", func(params ...string) (interface{}, error) {
		if len(params) != 1 {
			return nil, parametersError(nil)
		}
		return Lexify(params[0]), nil
	})

	RegisterProvider("Numerify", "string", func(params ...string) (interface{}, error) {
		if len(params) != 1 {
			return nil, parametersError(nil)
		}
		return Numerify(params[0]), nil
	})
}
