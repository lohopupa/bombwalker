package game

import (
	"minewalker/game/pkg/goui/elements"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/window"
)

func createHelloWindow(p platform.Platform, g *Game) *window.Window {
	w := window.NewWindow("Hello", p)
	w.FontFamily = "merchantD"
	psx, psy := p.GetSize()
	btn := elements.NewButton(psx/2, psy/3*2, "Start")
	btn.AlignY = elements.AlignYCenter
	btn.AlignX = elements.AlignXCenter
	btn.TextSize = 100
	btn.BorderThick = 10
	btn.OnClick = func() {
		g.ChangeWindow("main")
	}
	lbl := elements.NewLabel(psx/2, psy/4, "Welcome to MineWalker game")
	lbl.AlignX = elements.AlignXCenter
	lbl.TextSize = 70
	w.AddElement(btn)
	w.AddElement(lbl)
	return w
}

func createMainWindow(p platform.Platform, g *Game) *window.Window {
	w := window.NewWindow("Main", p)
	w.FontFamily = "merchantD"
	psx, psy := p.GetSize()
	btnWidth := 200.
	btnHeight := 75.
	btnOffset := 30.

	exitBtn := mainWindowButton("Exit", btnWidth, btnHeight)
	exitBtn.PosX = btnOffset
	exitBtn.PosY = btnOffset

	restartBtn := mainWindowButton("Restart", btnWidth, btnHeight)
	restartBtn.PosX = psx - btnOffset
	restartBtn.PosY = btnOffset
	restartBtn.AlignX = elements.AlignXRight

	settingsBtn := mainWindowButton("Settings", btnWidth, btnHeight)
	settingsBtn.PosX = btnOffset
	settingsBtn.PosY = psy - btnOffset
	settingsBtn.AlignY = elements.AlignYTop

	helpBtn := mainWindowButton("Help", btnWidth, btnHeight)
	helpBtn.PosX = psx - btnOffset
	helpBtn.PosY = psy - btnOffset
	helpBtn.AlignX = elements.AlignXRight
	helpBtn.AlignY = elements.AlignYTop

	gameE := GameUIElement{300, 30, psx - 600, psy - 60}
	
	w.AddElements(exitBtn, restartBtn, settingsBtn, helpBtn, &gameE)
	return w
}

func mainWindowButton(text string, w, h float64) *elements.Button {
	return &elements.Button{
		ElementBase: elements.ElementBase{
			Text: text,
			Width: w,
			Height: h,
			AdjustSizeToText: false,
			TextSize: 32,
		},
		BorderThick: 5,
	}
}
