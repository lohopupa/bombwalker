package platform

import (
	"fmt"
	"minewalker/game/pkg/gojs"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/types"
)

type WASM struct {
	ctx    gojs.Context2d
	events platform.EventsChan

	draw   bool
}

func (w *WASM) GetEventsChan() platform.EventsChan {
	return w.events
}

func (w *WASM) StartRendering(draw func(int)) {
	w.draw = true
	var innerFunc func(int)
	innerFunc = func(timestamp int) {
		draw(timestamp)
		if w.draw {
			gojs.RequestAnimationFrame(innerFunc)
		}
	}
	go gojs.RequestAnimationFrame(innerFunc)
}

func (w *WASM) StopRendering() {
	w.draw = false
}

func handleEvent(e gojs.Event, ch chan platform.Event) {
	switch event := e.(type) {
	case gojs.KeyboardEvent:
		{
			ch <- platform.Event{
				EventType: platform.EventTypeKeyPress,
				KeyCode:   event.Key,
			}
		}
	case gojs.MouseEvent:
		{
			var et platform.EventType
			switch event.EventType {
			case gojs.MouseClickEvent:
				{
					et = platform.EventTypeMouseClick
				}
			case gojs.MouseMoveEvent:
				{
					et = platform.EventTypeMouseMove
				}
			}
			ch <- platform.Event{
				EventType:   et,
				MouseButton: platform.MouseButton(event.Button),
				MousePosX:   float64(event.OffsetX),
				MousePosY:   float64(event.OffsetY),
			}
		}
	}
}

func InitWasm(canvasId string) *WASM {
	canvas, _ := gojs.GetElementById(canvasId)
	events := make(chan platform.Event)
	eventListener := func(e gojs.Event) {
		handleEvent(e, events)
	}
	gojs.AddEventListener(gojs.MouseClickEvent, eventListener)
	gojs.AddEventListener(gojs.MouseMoveEvent, eventListener)
	gojs.AddEventListener(gojs.KeyPressEvent, eventListener)
	return &WASM{ctx: *canvas.ToCanvas().GetContext2d(), events: events}
}
func (r WASM) ClearRect(x, y, width, height float64) {
	r.ctx.ClearRect(x, y, width, height)
}
func (r WASM) FillRect(x, y, width, height float64, color types.Color) {
	r.ctx.SetFillStyle(color.ToHexString())
	r.ctx.FillRect(x, y, width, height)

}
func (r WASM) StrokeRect(x, y, width, height, thick float64, color types.Color) {
	r.ctx.SetStrokeStyle(color.ToHexString())
	r.ctx.SetLineWidth(int(thick))
	r.ctx.StrokeRect(x, y, width, height)
}
func (r WASM) FillCircle(x, y, radius float64, color types.Color) {
	r.ctx.BeginPath()
	r.ctx.Arc(x, y, radius, 0, 2*3.1415, false)
	r.ctx.SetFillStyle(color.ToHexString())
	r.ctx.Fill()
}
func (r WASM) StrokeCircle(x, y, radius, thick float64, color types.Color) {
	r.ctx.BeginPath()
	r.ctx.Arc(x, y, radius, 0, 2*3.1415, false)
	r.ctx.SetLineWidth(int(thick))
	r.ctx.SetStrokeStyle(color.ToHexString())
	r.ctx.Stroke()
}
func (r WASM) Text(x, y, textSize float64, fontFamily, text string, color types.Color) {
	r.ctx.SetFillStyle(color.ToHexString())
	r.ctx.SetFont(fmt.Sprintf("%dpx %s", int(textSize), fontFamily))
	r.ctx.FillText(text, x, y+textSize)
}
func (r WASM) TextWidth(text, fontFamily string, textSize float64) float64 {
	r.ctx.SetFont(fmt.Sprintf("%dpx %s", int(textSize), fontFamily))
	return r.ctx.MeasureText(text).Width
}
func (r WASM) Line(sX, sY, eX, eY, thickness float32, color types.Color) {

}
func (r WASM) GetSize() (float64, float64) {
	return r.ctx.Width, r.ctx.Height
}
