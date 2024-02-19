package shapes

import (
	la "minewalker/game/internal/basics/linear_algebra"
)

type MU = float64 // Measure Unit
type vec2 = la.Vector2[MU]

type Rect struct {
	Pos  vec2
	Size vec2
}

type Circle struct {
	Center vec2
	Radius MU
}

func (r Rect) IsPointInside(point vec2) bool {
	return point.X >= r.Pos.X && point.X <= r.Pos.X+r.Size.X &&
		point.Y >= r.Pos.Y && point.Y <= r.Pos.Y+r.Size.Y
}

func (c Circle) IsPointInside(point vec2) bool {
	return c.Center.Distance(point) < c.Radius
}
