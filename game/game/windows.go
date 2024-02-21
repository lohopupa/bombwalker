package game

import (
	"fmt"
	"minewalker/game/pkg/goui/elements"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/types"
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
	btn.ClickOnKey = "Enter"
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
	exitBtn.ClickOnKey = "Escape"
	exitBtn.OnClick = func() {
		g.ChangeWindow("hello")
	}

	restartBtn := mainWindowButton("Restart", btnWidth, btnHeight)
	restartBtn.PosX = psx - btnOffset
	restartBtn.PosY = btnOffset
	restartBtn.AlignX = elements.AlignXRight
	restartBtn.ClickOnKey = "KeyR"
	restartBtn.OnClick = func() {
		g.ChangeWindow("shit")
	}

	settingsBtn := mainWindowButton("Settings", btnWidth, btnHeight)
	settingsBtn.PosX = btnOffset
	settingsBtn.PosY = psy - btnOffset
	settingsBtn.AlignY = elements.AlignYTop
	settingsBtn.ClickOnKey = "KeyO"
	settingsBtn.OnClick = func() {
		g.ChangeWindow("shit")
	}

	helpBtn := mainWindowButton("Help", btnWidth, btnHeight)
	helpBtn.PosX = psx - btnOffset
	helpBtn.PosY = psy - btnOffset
	helpBtn.AlignX = elements.AlignXRight
	helpBtn.AlignY = elements.AlignYTop
	helpBtn.ClickOnKey = "KeyH"
	helpBtn.OnClick = func() {
		g.ChangeWindow("help")
	}

	infoLblsCount := 3
	infoTextSize := 30.
	infoGap := 20.
	infoPos := psy/2 - float64(infoLblsCount)*(infoTextSize+infoGap)/2
	scoreLbl := elements.NewLabel(btnOffset, infoPos, "")
	scoreLbl.GetText = func() string { return fmt.Sprintf("Score: %d", g.Score) }
	scoreLbl.TextSize = infoTextSize
	infoPos += infoTextSize + infoGap

	gridSizeLbl := elements.NewLabel(btnOffset, infoPos, "")
	gridSizeLbl.GetText = func() string { return fmt.Sprintf("Grid: %dx%d", g.GridSize, g.GridSize) }
	gridSizeLbl.TextSize = infoTextSize
	infoPos += infoTextSize + infoGap

	bombsLbl := elements.NewLabel(btnOffset, infoPos, "")
	bombsLbl.GetText = func() string { return fmt.Sprintf("Bombs: %d/%d", g.BombsOpen, g.BombsTotal) }
	bombsLbl.TextSize = infoTextSize

	gameE := GameUIElement{300, 30, psx - 600, psy - 60}

	w.AddElements(exitBtn, restartBtn, settingsBtn, helpBtn, scoreLbl, gridSizeLbl, bombsLbl, &gameE)
	return w
}

func createHelpWindow(p platform.Platform, g *Game) *window.Window {
	w := window.NewWindow("Help", p)
	w.FontFamily = "merchantD"
	psx, psy := p.GetSize()

	textSize := 30.

	rulesText := "The rules of this game are the same as those of the classic minesweeper, but instead of simply clicking on mines, you need to walk on them without getting blown up!"
	rulesPad := 50.
	rulesWidth := psx - rulesPad*2
	rulesLbl := elements.NewLabelML(rulesPad, rulesPad, rulesWidth, rulesText)
	rulesLbl.TextSize = textSize

	h1 := elements.NewLabelML(rulesPad, 300, rulesWidth, "- Use the arrow keys or WASD to navigate")
	h1.TextSize = textSize
	// h1.AlignText = elements.AlignTextLeft

	h2 := elements.NewLabelML(rulesPad, 360, rulesWidth, "- Click on the cell you want to visit")
	h2.TextSize = textSize
	// h2.AlignText = elements.AlignTextLeft

	h3 := elements.NewLabelML(rulesPad, 420, rulesWidth, "- Press the space bar or right mouse button to mark a cell")
	h3.TextSize = textSize
	// h3.AlignText = elements.AlignTextLeft

	okBtn := elements.NewButton(psx/2, psy/6*5, "Ok!")
	okBtn.AlignX = elements.AlignXCenter
	okBtn.AlignY = elements.AlignYCenter
	okBtn.AdjustSizeToText = false
	okBtn.Width = 200
	okBtn.Height = 70
	okBtn.TextSize = textSize
	okBtn.ClickOnKey = "Enter"
	okBtn.OnClick = func() {
		g.ChangeWindow("main")
	}

	w.AddElements(rulesLbl, h1, h2, h3, okBtn)
	return w
}

func mainWindowButton(text string, w, h float64) *elements.Button {
	cs := types.DefaultColorScheme()
	// cs.AccentColor = types.FromHexString("#ADD8E6")
	// cs.AccentColorHighlight = types.FromHexString("#FFFFFF")
	return &elements.Button{
		ElementBase: elements.ElementBase{
			GetText:          func() string { return text },
			Width:            w,
			Height:           h,
			AdjustSizeToText: false,
			TextSize:         32,
			ColorScheme:      &cs,
		},
		BorderThick: 5,
	}
}

func createShitWindow(p platform.Platform, g *Game) *window.Window {
	w := window.NewWindow("Shit", p)
	w.FontFamily = "merchantD"
	psx, psy := p.GetSize()
	btn := elements.NewButton(psx/2, psy/3*2, "Back")
	btn.AlignY = elements.AlignYCenter
	btn.AlignX = elements.AlignXCenter
	btn.TextSize = 100
	btn.BorderThick = 10
	btn.ClickOnKey = "Enter"
	btn.OnClick = func() {
		g.ChangeWindow("main")
	}
	lbl := elements.NewLabel(psx/2, psy/4, "Oh shit, there is nothing here!")
	lbl.AlignX = elements.AlignXCenter
	lbl.TextSize = 65
	w.AddElement(btn)
	w.AddElement(lbl)
	return w
}
