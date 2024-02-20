package renderer

import (
	"fmt"
	"minewalker/game/internal/basics/color"
	la "minewalker/game/internal/basics/linear_algebra"
	"minewalker/game/internal/basics/shapes"
	"minewalker/gojs"
)

type Renderer struct {
	ctx gojs.Context2d
}

func Init(ctx gojs.Context2d) Renderer{
	return Renderer{ctx: ctx}
}

func (r Renderer) FillRect(rect shapes.Rect, c color.Color) {
	r.ctx.SetFillStyle(c.ToHexString())
	r.ctx.FillRect(rect.Pos.X, rect.Pos.Y, rect.Size.X, rect.Size.Y)
}

func (r Renderer) StrokeRect(rect shapes.Rect, c color.Color, width int) {
	r.ctx.SetStrokeStyle(c.ToHexString())
	r.ctx.SetLineWidth(width)
	r.ctx.StrokeRect(rect.Pos.X, rect.Pos.Y, rect.Size.X, rect.Size.Y)
}

func (r Renderer) FillCircle(circle shapes.Circle, c color.Color) {
	r.ctx.BeginPath();
	r.ctx.Arc(circle.Center.X, circle.Center.Y, circle.Radius, 0, 2 * 3.1415, false);
	r.ctx.SetFillStyle(c.ToHexString())
	r.ctx.Fill();
}

func (r Renderer) StrokeCircle(circle shapes.Circle, c color.Color) {
	r.ctx.BeginPath();
	r.ctx.Arc(circle.Center.X, circle.Center.Y, circle.Radius, 0, 2 * 3.1415, false);
	r.ctx.SetStrokeStyle(c.ToHexString())
	r.ctx.Stroke();
}

func (r Renderer) Clear(c color.Color) {
	r.ctx.SetFillStyle(c.ToHexString())
	r.ctx.FillRect(0, 0, r.ctx.Width, r.ctx.Width)
}

func (r Renderer) Text(pos la.Vector2[float64], text string) {
	r.ctx.SetFillStyle(color.Color{R: 0, G: 0, B: 0, A: 255}.ToHexString())
	r.ctx.FillText(text, pos.X, pos.Y)
}

func (r Renderer) TextStyled(pos la.Vector2[float64], color color.Color, fontSizePX int, fontFamily, text string) {
	r.ctx.SetFillStyle(color.ToHexString())
	r.ctx.SetFont(fmt.Sprintf("%dpx %s", fontSizePX, fontFamily))
	r.ctx.FillText(text, pos.X, pos.Y + float64(fontSizePX))
}

func (r Renderer) MesureTextStyled(pos la.Vector2[float64], color color.Color, fontSizePX int, fontFamily, text string) gojs.TextMetrics {
	r.ctx.SetFillStyle(color.ToHexString())
	r.ctx.SetFont(fmt.Sprintf("%dpx %s", fontSizePX, fontFamily))
	// r.ctx.FillText(text, pos.X, pos.Y + float64(fontSizePX))
	return r.ctx.MeasureText(text)
}