package faker

import (
	"encoding/hex"
)

// Bool will build a random boolean value (true or false).
func Bool() bool {
	return IntInRange(0, 1) == 0
}

// PhoneNumber will build a random phone number string.
func PhoneNumber() string {
	formats := []string{"???-???-????", "(???) ???-????", "1-???-???-????", "???.???.????"}
	i := IntInRange(0, len(formats)-1)
	return Numerify(formats[i])
}

// UUID will build a random UUID string.
func UUID() string {
	version := byte(4)
	uuid := make([]byte, 16)
	_, err := random.Read(uuid)
	if err != nil {
		panic(err)
	}

	// Set version
	uuid[6] = (uuid[6] & 0x0f) | (version << 4)

	// Set variant
	uuid[8] = (uuid[8] & 0xbf) | 0x80

	buf := make([]byte, 36)
	var dash byte = '-'
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], uuid[10:])

	return string(buf)
}

// Builder functions

func boolBuilder(params ...string) (interface{}, error) {
	return Bool(), nil
}

func phoneNumberBuilder(params ...string) (interface{}, error) {
	return PhoneNumber(), nil
}

func uuidBuilder(params ...string) (interface{}, error) {
	return UUID(), nil
}
