package shapes

import (
	la "minewalker/game/internal/basics/linear_algebra"
)

type MU = float32 // Measure Unit
type vec2 = la.Vector2[MU]

type Rect struct {
	pos vec2
	size vec2
}

type Circle struct {
	center vec2
	radius MU
}

func (r Rect) IsPointInside(point vec2) bool {
	return point.X >= r.pos.X && point.X <= r.pos.X+r.size.X &&
		point.Y >= r.pos.Y && point.Y <= r.pos.Y+r.size.Y
}

func (c Circle) IsPointInside(point vec2) bool {
	return c.center.Distance(point) < c.radius
}