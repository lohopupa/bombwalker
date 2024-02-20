package color

import (
	"fmt"
	"math"
	"strconv"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func FromHexString(hex string) Color {
	if len(hex) != 9 || hex[0] != '#' {
		return Color{}
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return Color{}
	}

	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return Color{}
	}

	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return Color{}
	}

	a, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return Color{}
	}

	return Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
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

func (c Color) ToRGBAString() string {
	return fmt.Sprintf("rgb(%d %d %d)", c.R, c.G, c.B)
}

func (c Color) SetA(a uint8) Color {
	c.A = a
	return c
}

func (c Color) SetR(r uint8) Color {
	c.R = r
	return c
}
func (c Color) SetG(g uint8) Color {
	c.G = g
	return c
}
func (c Color) SetB(b uint8) Color {
	c.B = b
	return c
}