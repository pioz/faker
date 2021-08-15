package faker

import (
	"fmt"
)

var colorData = PoolGroup{
	"name": {"amaranth", "amber", "amethyst", "apricot", "aquamarine", "azure", "beige", "black", "blue", "blush", "bronze", "brown", "burgundy", "byzantium", "carmine", "cerise", "cerulean", "champagne", "chocolate", "coffee", "copper", "coral", "crimson", "cyan", "emerald", "erin", "fuchsia", "gold", "gray", "green", "grey", "harlequin", "indigo", "ivory", "jade", "lavender", "lemon", "lilac", "lime", "magenta", "maroon", "mauve", "ochre", "olive", "orange", "orchid", "peach", "pear", "periwinkle", "pink", "plum", "puce", "purple", "raspberry", "red", "rose", "ruby", "salmon", "sangria", "sapphire", "scarlet", "silver", "sky", "tan", "taupe", "teal", "turquoise", "ultramarine", "violet", "viridian", "white", "yellow"},
}

// ColorName will build a random color name string.
func ColorName() string {
	value, _ := GetData("color", "name")
	return value.(string)
}

// ColorHex will build a random hex color string. First element is the Red value
// from 0 to 255; second element is the Green value from 0 to 255; third element
// is the Blue value from 0 to 255.
func ColorHex() string {
	c := ColorRGB()
	hex := "#" + getHex(c[0]) + getHex(c[1]) + getHex(c[2])
	return hex
}

// ColorRGB will build a random RGB color.
func ColorRGB() [3]int {
	var c [3]int
	c[0] = IntInRange(0, 255)
	c[1] = IntInRange(0, 255)
	c[2] = IntInRange(0, 255)
	return c
}

// ColorHSL will build a random HSL color. First element is Hue, a degree on the
// color wheel from 0 to 360. 0 is red, 120 is green, 240 is blue. Second
// element is Saturation, a percentage value; 0 means a shade of gray and 100 is
// the full color. Third element is Lightness, also a percentage; 0 is black,
// 100 is white.
func ColorHSL() [3]int {
	var c [3]int
	c[0] = IntInRange(0, 360)
	c[1] = IntInRange(0, 100)
	c[2] = IntInRange(0, 100)
	return c
}

// Builder functions

func colorNameBuilder(params ...string) (interface{}, error) {
	return ColorName(), nil
}

func colorHexBuilder(params ...string) (interface{}, error) {
	return ColorHex(), nil
}

func colorRGBBuilder(params ...string) (interface{}, error) {
	return ColorRGB(), nil
}

func colorHSLBuilder(params ...string) (interface{}, error) {
	return ColorHSL(), nil
}

// Private functions

func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}
