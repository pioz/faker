package factory

func init() {
	RegisterProvider("Bool", "bool", func(params ...string) (interface{}, error) {
		return Bool(), nil
	})

	RegisterProvider("", "bool", func(params ...string) (interface{}, error) {
		return Bool(), nil
	})
}
