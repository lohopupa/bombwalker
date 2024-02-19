package linearalgebra

import (
	"math"
)

type Numbers interface {
	// TODO: Add all numeric types
	uint | int | float32
}

type Vector2[T Numbers] struct {
	X T
	Y T
}

func Vector2New[T Numbers](x, y T) Vector2[T] {
	return Vector2[T]{X: x, Y: y}
}

func Vector2Zero[T Numbers]() Vector2[T] {
	return Vector2[T]{X: 0, Y: 0}
}

func (v Vector2[T]) Add(v2 Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector2[T]) Sub(v2 Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vector2[T]) Mul(v2 Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
	}
}

func (v Vector2[T]) Scale(scalar T) Vector2[T] {
	return Vector2[T]{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2[T]) Div(scalar T) Vector2[T] {
	return Vector2[T]{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2[T]) Dot(v2 Vector2[T]) T {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector2[T]) Magnitude() float32 {
	return v.Distance(v)
}

func (v Vector2[T]) Distance(v2 Vector2[T]) float32 {
	return float32(math.Sqrt(float64(v.X*v2.X + v.Y*v2.Y)))
}

func (v Vector2[T]) Normalize() Vector2[float32] {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return Vector2[float32]{X: 0, Y: 0}
	}

	return Vector2[float32]{
		X: float32(v.X),
		Y: float32(v.Y),
	}.Div(magnitude)
}
