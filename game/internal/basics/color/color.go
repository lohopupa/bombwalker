package color

import (
	"fmt"
	"strconv"
	"math"
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

func FromHSV(h, s, v uint8) Color {
	
	hi := float64(h) / 60.0
	hiFloor := math.Floor(hi)
	f := hi - hiFloor
	sf := float64(s) / 255.0
	vf := float64(v) / 255.0

	p := uint8(vf * (1 - sf))
	q := uint8(vf * (1 - sf*f))
	t := uint8(vf * 1 - sf*(1.0-f))

	switch hiFloor {
	case 0:
		return Color{R: v, G: t, B: p, A: 255}
	case 1:
		return Color{R: q, G: v, B: p, A: 255}
	case 2:
		return Color{R: p, G: v, B: t, A: 255}
	case 3:
		return Color{R: p, G: q, B: v, A: 255}
	case 4:
		return Color{R: t, G: p, B: v, A: 255}
	default:
		return Color{R: v, G: p, B: q, A: 255}
	}
}

func (c Color) ToHexString() string {
	
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}
