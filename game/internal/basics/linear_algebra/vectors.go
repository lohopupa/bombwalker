package linearalgebra

import (
	"math"	
)

type Numbers interface {
	// TODO: Add all numeric types
	uint | int | float32 | float64
}

type Vector2[T Numbers] struct {
	x T
	y T
}

func Vector2New[T Numbers](x, y T) Vector2[T] {
	return Vector2[T]{x: x, y: y}
}

func Vector2Zero[T Numbers]() Vector2[T] {
	return Vector2[T]{x: 0, y: 0}
}

func (v Vector2[T]) Add(v2 Vector2[T]) Vector2[T] {
	return Vector2[T]{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}

func (v Vector2[T]) Sub(v2 Vector2[T]) Vector2[T] {
	return Vector2[T]{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}
}

func (v Vector2[T]) Mul(v2 Vector2[T]) Vector2[T] {
	return Vector2[T]{
		x: v.x * v2.x,
		y: v.y * v2.y,
	}
}

func (v Vector2[T]) Scale(scalar T) Vector2[T] {
	return Vector2[T]{
		x: v.x * scalar,
		y: v.y * scalar,
	}
}

func (v Vector2[T]) Div(scalar T) Vector2[T] {
	return Vector2[T]{
		x: v.x / scalar,
		y: v.y / scalar,
	}
}


func (v Vector2[T]) Dot(v2 Vector2[T]) T {
	return v.x*v2.x + v.y*v2.y
}

func (v Vector2[T]) Magnitude() T {
	return T(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}

func (v Vector2[T]) Normalize() Vector2[T] {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector2[T]{x: 0, y: 0}
	}
	return v.Div(magnitude)
}