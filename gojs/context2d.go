package gojs

import (
	"syscall/js"
)

type Context2d struct {
	instance js.Value
	Height   float64
	Width    float64
}

type TextMetrics struct {
	Width                  float64
	ActualBoundingBoxLeft  float64
	ActualBoundingBoxRight float64
	FontBoundingBoxAscent  float64
	FontBoundingBoxDescent float64
	EmHeightAscent         float64
	EmHeightDescent        float64
	HangingBaseline        float64
	AlphabeticBaseline     float64
	IdeographicBaseline    float64
}

func (this Context2d) ClearRect(x, y, width, height float64) {
	this.instance.Call("clearRect", x, y, width, height)
}

func (c *Context2d) FillRect(x, y, width, height float64) {
	c.instance.Call("fillRect", x, y, width, height)
}

func (c *Context2d) StrokeRect(x, y, width, height float64) {
	c.instance.Call("strokeRect", x, y, width, height)
}

func (c *Context2d) FillText(text string, x, y float64) {
	c.instance.Call("fillText", text, x, y)
}

func (c *Context2d) StrokeText(text string, x, y float64) {
	c.instance.Call("strokeText", text, x, y)
}

func (c *Context2d) MeasureText(text string) TextMetrics {
	textMetrics := c.instance.Call("measureText", text)

	return TextMetrics{
		Width:                  textMetrics.Get("width").Float(),
		ActualBoundingBoxLeft:  textMetrics.Get("actualBoundingBoxLeft").Float(),
		ActualBoundingBoxRight: textMetrics.Get("actualBoundingBoxRight").Float(),
		FontBoundingBoxAscent:  textMetrics.Get("fontBoundingBoxAscent").Float(),
		FontBoundingBoxDescent: textMetrics.Get("fontBoundingBoxDescent").Float(),
		EmHeightAscent:         textMetrics.Get("emHeightAscent").Float(),
		EmHeightDescent:        textMetrics.Get("emHeightDescent").Float(),
		HangingBaseline:        textMetrics.Get("hangingBaseline").Float(),
		AlphabeticBaseline:     textMetrics.Get("alphabeticBaseline").Float(),
		IdeographicBaseline:    textMetrics.Get("ideographicBaseline").Float(),
	}
}

func (c *Context2d) MoveTo(x, y float64) {
	c.instance.Call("moveTo", x, y)
}

func (c *Context2d) LineTo(x, y float64) {
	c.instance.Call("lineTo", x, y)
}

func (c *Context2d) Arc(x, y, radius, startAngle, endAngle float64, anticlockwise bool) {
	c.instance.Call("arc", x, y, radius, startAngle, endAngle, anticlockwise)
}

func (c *Context2d) BeginPath() {
	c.instance.Call("beginPath")
}

func (c *Context2d) ClosePath() {
	c.instance.Call("closePath")
}

func (c *Context2d) Fill() {
	c.instance.Call("fill")
}

func (c *Context2d) Stroke() {
	c.instance.Call("stroke")
}

func (c *Context2d) SetFillStyle(style string) {
	c.instance.Set("fillStyle", style)
}

func (c *Context2d) SetStrokeStyle(style string) {
	c.instance.Set("strokeStyle", style)
}

func (c *Context2d) SetFont(font string) {
	c.instance.Set("font", font)
}
