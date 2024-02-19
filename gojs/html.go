package gojs

import (
	"syscall/js"
)

type HtmlElement struct {
	instance js.Value
	id string
}

type HtmlCanvasElement struct {
	HtmlElement
	width int
	height int
}

func (he HtmlElement) ToCanvas() HtmlCanvasElement {
	// TODO: Check if element is canvas
	return HtmlCanvasElement{
		HtmlElement: he,
		width: he.instance.Get("width").Int(),
		height: he.instance.Get("height").Int(),
	}
}

func (this HtmlCanvasElement) GetContext2d() *Context2d {
	ctx := Context2d{
		instance: this.instance.Call("getContext", "2d"),
	}
	return &ctx
} 



