package main

import (
	"fmt"
	"minewalker/game/internal/basics/color"
	la "minewalker/game/internal/basics/linear_algebra"
	"minewalker/game/internal/basics/shapes"
	"minewalker/game/internal/renderer"
	"minewalker/gojs"
)

type Game struct {
	R        renderer.Renderer
	Val      int
	PrevTime int
	Rect     shapes.Rect
	Vel      la.Vector2[float64]
	Size     la.Vector2[float64]
}

var g Game

func main() {
	fmt.Println("Hello from WASM")

	htmlCanvas, err := gojs.GetElementById("canvas")
	if err != nil {
		fmt.Println(err)
	}
	canvas := htmlCanvas.ToCanvas()
	ctx := canvas.GetContext2d()
	g.R = renderer.Init(*ctx)
	g.Val = 0
	g.PrevTime = 0
	g.Rect = shapes.Rect{
		Pos:  la.Vector2New[float64](1, 1),
		Size: la.Vector2New[float64](100, 100),
	}
	g.Size = la.Vector2New(ctx.Width, ctx.Height)

	gojs.AddEventListener(gojs.KeyDownEvent, KeyboardEventHandler)
	gojs.RequestAnimationFrame(NextFrame)
	select {}
}

func NextFrame(timestamp int) {
	dt := timestamp - g.PrevTime
	g.PrevTime = timestamp
	g.Val = int((float64(g.Val) + float64(dt)*0.25)) % 360
	g.R.Clear(color.FromHSV(uint(g.Val), 0.8, 0.9))
	g.Rect.Pos = g.Rect.Pos.Add(g.Vel.Scale(float64(dt) * 0.05))
	// g.Vel = g.Vel.Scale(0.99)
	if g.Rect.Pos.X+g.Rect.Size.X >= g.Size.X {
		g.Rect.Pos.X = g.Size.X - g.Rect.Size.X
		// g.Vel.X *= -0.9
	}
	if g.Rect.Pos.X <= 0 {
		g.Rect.Pos.X = 0
		// g.Vel.X *= -0.9
	}
	if g.Rect.Pos.Y + g.Rect.Size.Y >= g.Size.Y {
		g.Rect.Pos.Y = g.Size.Y - g.Rect.Size.Y
		// g.Vel.Y *= -0.9
	}
	if g.Rect.Pos.Y <= 0 {
		g.Rect.Pos.Y = 0
		// g.Vel.Y *= -0.9
	}
	// g.Vel.Y += 0.05
	// g.Rect.Pos.Y *= 1.01
	g.R.FillRect(g.Rect, color.Color{})
	gojs.RequestAnimationFrame(NextFrame)
}

func KeyboardEventHandler(e gojs.Event) {
	switch ke := e.(type) {
	case gojs.KeyboardEvent:
		{
			switch ke.Code {
			case "KeyW":
				if ke.ShiftKey {
					g.Rect.Size.Y -= 10
				} else {
					g.Rect.Pos.Y -= 10
				}
				break
			case "KeyA":
				if ke.ShiftKey {
					g.Rect.Size.X -= 10
				} else {
					g.Rect.Pos.X -= 10
				}
				break
			case "KeyS":
				if ke.ShiftKey {
					g.Rect.Size.Y += 10
				} else {
					g.Rect.Pos.Y += 10
				}
				break
			case "KeyD":
				if ke.ShiftKey {
					g.Rect.Size.X += 10
				} else {
					g.Rect.Pos.X += 10
				}
				break
			}
		}
	default:
		return
	}
}
