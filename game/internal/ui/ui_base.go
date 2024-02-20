package ui

import (
	"minewalker/game/internal/basics/color"
	"minewalker/game/internal/basics/shapes"
	"minewalker/game/internal/renderer"

	la "minewalker/game/internal/basics/linear_algebra"
)

type vec2 = la.Vector2[float64]

type Page struct {
	Rect shapes.Rect

	Color color.Color

	Elements  []*UiElement
	PopupOpen bool
	Popup     *Page
}

func NewPage(rect shapes.Rect, e ...UiElement) *Page {
	p := &Page{
		Rect:      rect,
		PopupOpen: false,
	}
	for _, ee := range e {
		p.Elements = append(p.Elements, &ee)
	}
	return p
}

func (p Page) AddElement(e UiElement) {
	p.Elements = append(p.Elements, &e)
}

func (p Page) Draw(r renderer.Renderer) {
	r.FillRect(p.Rect, p.Color)
	// for _, e := range p.Elements {
	// 	go (*e).Draw(r)
	// }
	if p.PopupOpen {
		r.FillRect(p.Rect, color.FromHexString("#50505002"))
		p.Popup.Draw(r)
	}
}

func (p Page) GetElementUnderPoind(point vec2) *UiElement {
	if p.PopupOpen {
		return p.Popup.GetElementUnderPoind(point)
	}
	for _, e := range p.Elements {
		if (*e).GetRect().IsPointInside(point) {
			return e
		}
	}
	return nil
}

type UiElement interface {
	Draw(renderer.Renderer, bool)
	OnClick()
	GetRect() shapes.Rect
	SetHover(bool)
}

type TextAlign uint8

const (
	AlignLeft TextAlign = iota
	AlignCenter
	AlignRight
)

type UiConfig struct {
	Page *Page
	Rect shapes.Rect

	Align TextAlign

	Text      string
	TextColor color.Color
	TextSize  int
	TextFont  string

	BorderSize  int
	BorderColor color.Color
	// BorderRadius int

	BackgroundColor color.Color

	Hover bool

	ClickCallBack func()
}

func (b UiConfig) OnClick() {
	b.ClickCallBack()
}

func (b UiConfig) SetHover(h bool) {
	b.Hover = h
}

func (b UiConfig) GetRect() shapes.Rect {
	return b.Rect
}

type Button struct {
	UiConfig
}

func (b Button) Draw(r renderer.Renderer, Hover bool) {
	rect := b.Rect
	switch b.Align {
	case AlignCenter:{
		rect.Pos.X = rect.Pos.X - (rect.Size.X / 2)
		break
	}
	case AlignLeft:{
		break
	}
	case AlignRight:{
		rect.Pos.X = rect.Pos.X - rect.Size.X
		break
	}

	}
	c := b.BackgroundColor
	if Hover {
		c = color.Color{R: c.R+10, G: c.G+10, B: c.B+10, A: c.A}
	}
		r.FillRect(rect, c)
		r.StrokeRect(rect, b.BorderColor, b.BorderSize)
	textSize := vec2{
		X: r.MesureTextStyled(rect.Pos, b.TextColor, b.TextSize, b.TextFont, b.Text).Width,
		Y: float64(b.TextSize),
	}
	textPos := rect.Size.Sub(textSize).Scale(0.5).Add(rect.Pos)
	
	r.TextStyled(textPos, b.TextColor, b.TextSize, b.TextFont, b.Text)
}

type Label struct {
	UiConfig
}

func (b Label) Draw(r renderer.Renderer, Hover bool) {
	rect := b.Rect
	// if b.Hover {
	// 	rect.Pos = rect.Pos.Sub(la.Vector2New[float64](2, 2))
	// 	rect.Size = rect.Pos.Add(la.Vector2New[float64](4, 4))
	// }
	// r.FillRect(rect, b.BackgroundColor)
	// r.StrokeRect(rect, b.BorderColor, b.BorderSize)
	wX := r.MesureTextStyled(rect.Pos, b.TextColor, b.TextSize, b.TextFont, b.Text).Width
	switch b.Align {
	case AlignCenter:{
		rect.Pos.X = rect.Pos.X - (wX / 2)
		break
	}
	case AlignLeft:{
		break
	}
	case AlignRight:{
		rect.Pos.X = rect.Pos.X - wX
		break
	}

	}
	r.TextStyled(rect.Pos, b.TextColor, b.TextSize, b.TextFont, b.Text)
}
