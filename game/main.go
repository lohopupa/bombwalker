package main

import (
	"fmt"
	la "minewalker/game/internal/basics/linear_algebra"
	"minewalker/game/internal/basics/shapes"
	"minewalker/game/internal/game"
	"minewalker/game/internal/renderer"
	"minewalker/game/internal/ui"
	"minewalker/gojs"
)

var p *ui.Page

func main() {
	game_ctx := game.GetGameContext()
	htmlCanvas, err := gojs.GetElementById("canvas")
	if err != nil {
		fmt.Println(err)
	}
	canvas := htmlCanvas.ToCanvas()
	ctx := canvas.GetContext2d()
	game_ctx.Renderer = renderer.Init(*ctx)

	canvas.AddEventListener(gojs.MouseMove, MouseMoveHandler)

	p = ui.CreateMainPage(
		shapes.Rect{
			Pos: la.Vector2Zero[float64](), 
			Size: la.Vector2[float64]{
				X: ctx.Width, 
				Y: ctx.Height,
			},
		},
	)
	gojs.RequestAnimationFrame(NextFrame)
	select {}
}

func NextFrame(timestamp int) {
	gctx := game.GetGameContext()
	// dt := timestamp - game.GetGameContext().PrevTime
	for _, e := range p.Elements {
		hover:=(*e).GetRect().IsPointInside(gctx.MousePos)
		(*e).Draw(gctx.Renderer, hover)
	}
	p.Draw(gctx.Renderer)

	gojs.RequestAnimationFrame(NextFrame)
}

func MouseMoveHandler(e gojs.Event) {
	gctx := game.GetGameContext()
	switch ev := e.(type) {
	case gojs.MouseEvent:
		{
			gctx.MousePos = la.Vector2[float64]{X: float64(ev.OffsetX), Y: float64(ev.OffsetY)}
		}
	}
}
