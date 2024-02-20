package ui

import (
	"fmt"
	"minewalker/game/internal/basics/color"
	"minewalker/game/internal/basics/shapes"
)

func CreateMainPage(rect shapes.Rect) *Page {

	labelHello := Label{
		UiConfig: UiConfig{
			// Page: p,
			Text:            "Wellcome to MineWalker game",
			Rect:            shapes.Rect{Pos: vec2{X: rect.Size.X / 2, Y: rect.Pos.Y + 150}, Size: vec2{X: 0, Y: 0}},
			TextColor:       color.FromHexString("#AAAAAAFF"),
			TextSize:        60,
			TextFont:        "merchant",
			BackgroundColor: color.Color{R: 0, G: 0, B: 0, A: 0},
			Hover:           false,
			Align: AlignCenter,
		},
	}
	buttonStart := Button{
		UiConfig: UiConfig{
			// Page: p,
			Text:            "Start",			
			Rect:            shapes.Rect{Pos: vec2{X: rect.Size.X / 2, Y: rect.Pos.Y + 400}, Size: vec2{X: 400, Y: 150}},
			TextColor:       color.FromHexString("#AAAAAAFF"),
			TextSize:        80,
			TextFont:        "merchant",
			BorderSize:      10,
			BorderColor:     color.FromHexString("#AAAAAAFF"),
			BackgroundColor: color.FromHexString("#000000FF"),
			ClickCallBack:   func() { fmt.Println("Hello button") },
			Hover:           false,
			Align: AlignCenter,
		},
	}
	p := NewPage(rect, labelHello, buttonStart)
	p.Color = color.FromHexString("#181818FF")
	return p

}
