package types

import (
	"math"
	"strconv"
	"fmt"
)

type Color struct {
	R, G, B, A uint8
}

func FromHexString(hex string) Color {
	if len(hex) != 7 && len(hex) != 9 || hex[0] != '#' {
		return Color{}
	}

	parseHex := func(s string) uint8 {
		val, _ := strconv.ParseUint(s, 16, 8)
		return uint8(val)
	}

	r := parseHex(hex[1:3])
	g := parseHex(hex[3:5])
	b := parseHex(hex[5:7])

	a := uint8(255)
	if len(hex) == 9 {
		a = parseHex(hex[7:9])
	}

	return Color{R: r, G: g, B: b, A: a}
}

func FromHSV(h uint, s, v float64) Color {
	if h < 0 || h >= 360 || s < 0 || s > 1 || v < 0 || v > 1 {
		return Color{}
	}
	C := v * s
	X := C * (1 - math.Abs(math.Mod(float64(h)/60, 2)-1))
	m := v - C
	var Rnot, Gnot, Bnot float64
	switch {
	case 0 <= h && h < 60:
		Rnot, Gnot, Bnot = C, X, 0
	case 60 <= h && h < 120:
		Rnot, Gnot, Bnot = X, C, 0
	case 120 <= h && h < 180:
		Rnot, Gnot, Bnot = 0, C, X
	case 180 <= h && h < 240:
		Rnot, Gnot, Bnot = 0, X, C
	case 240 <= h && h < 300:
		Rnot, Gnot, Bnot = X, 0, C
	case 300 <= h && h < 360:
		Rnot, Gnot, Bnot = C, 0, X
	}
	r := uint8(math.Round((Rnot + m) * 255))
	g := uint8(math.Round((Gnot + m) * 255))
	b := uint8(math.Round((Bnot + m) * 255))
	return Color{R: r, G: g, B: b, A: 255}
}

func (c Color) ToHexString() string {

	return fmt.Sprintf("#%02X%02X%02X%02X", c.R, c.G, c.B, c.A)
}