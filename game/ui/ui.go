package ui

import (
	"fmt"
	"minewalker/game/pkg/goui/elements"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/window"
)

func CreateMainPage(p platform.Platform) *window.Window {
	w := window.NewWindow("Hello", p)
	w.FontFamily = "merchant"
	psx, psy := p.GetSize()
	btn := elements.NewButton(psx/2, psy/3*2, "Start")
	btn.AlignY = elements.AlignYCenter
	btn.AlignX = elements.AlignXCenter
	btn.TextSize = 100
	btn.BorderThick = 10
	btn.OnClick = func() {
		fmt.Println("Click from wASSm")
	}
	lbl := elements.NewLabel(psx/2, psy/4, "Welcome to MineWalker game")
	lbl.AlignX = elements.AlignXCenter
	lbl.TextSize = 70
	w.AddElement(btn)
	w.AddElement(lbl)
	return w
}