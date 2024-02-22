package game

import (
	"fmt"
	"math"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/types"
)

type GameUIElement struct {
	PosX, PosY, SizeX, SizeY float64
	GameState                *Game
}

func (e *GameUIElement) SetColors(cs types.ColorScheme) {
}

func (e *GameUIElement) SetFontFamily(ff string) {
}

func (e *GameUIElement) Draw(p platform.Platform) {
	// x, y, w, h := e.GetBoundary(p)
	// p.FillRect(x, y, w, h, types.FromHexString("#181818"))
	e.drawGrid(p)
	// e.drawPlayer(p)
	if !e.GameState.Alive {
		e.drawDiedScreen(p)
	}
	if e.GameState.Win {
		e.drawWinScreen(p)
	}

}
func (e *GameUIElement) GetBoundary(platform.Platform) (float64, float64, float64, float64) {
	x, y, w, h := e.PosX, e.PosY, e.SizeX, e.SizeY
	sz := min(h, w)
	px := x - (sz-w)/2
	py := y - (sz-h)/2
	return px, py, sz, sz
}

func (g *GameUIElement) HandleEvent(e platform.Event, p platform.Platform) {
	if !g.GameState.Alive || g.GameState.Win {
		return
	}
	x, y, w, h := g.GetBoundary(p)
	mouseOver := e.MousePosX > x && e.MousePosX < x+w && e.MousePosY > y && e.MousePosY < y+h
	switch e.EventType {
	case platform.EventTypeMouseClick:
		{
			if mouseOver {
				cellIdx := g.coordsToIndex(e.MousePosX, e.MousePosY, p)
				switch e.MouseButton {
				case 0:
						g.GameState.OpenCell(cellIdx)
					case 2:
						g.GameState.MarkCell(cellIdx)
					
				}
			}
		}
	case platform.EventTypeMouseMove:
		{
			for idx, c := range g.GameState.Map {
				c.Hover = idx == g.coordsToIndex(e.MousePosX, e.MousePosY, p)
			}
		}
	case platform.EventTypeKeyPress:
		{
			switch e.KeyCode {
			case "ArrowUp":
				fallthrough
			case "KeyW":
				{
					if g.GameState.PlayerY > 0 {
						g.GameState.PlayerY -= 1
					}
				}
			case "ArrowDown":
				fallthrough
			case "KeyS":
				{
					if g.GameState.PlayerY < g.GameState.GridSize - 1 {
						g.GameState.PlayerY += 1
					}
				}
			case "ArrowRight":
				fallthrough
			case "KeyD":
				{
					if g.GameState.PlayerX < g.GameState.GridSize - 1 {
						g.GameState.PlayerX += 1
					}
				}
			case "ArrowLeft":
				fallthrough
			case "KeyA":
				{
					if g.GameState.PlayerX > 0 {
						g.GameState.PlayerX -= 1
					}
				}
			}
			g.GameState.OpenCell(int(g.GameState.PlayerX * g.GameState.GridSize + g.GameState.PlayerY))

		}
	}
}

// func (e *GameUIElement)drawPlayer(p platform.Platform) {

// }

func (e *GameUIElement) drawGrid(p platform.Platform) {
	x, y, w, _ := e.GetBoundary(p)
	gs := int(e.GameState.GridSize)
	cellSize := w / float64(gs)
	for px := 0; px < gs; px += 1 {
		for py := 0; py < gs; py += 1 {
			c := e.GameState.Map[px*gs+py]
			cText, cColor := getCellTextAndColor(*c)
			var stroke float64
			if c.Hover && e.GameState.Alive && !e.GameState.Win {
				stroke = 4
			} else {
				stroke = 2
			}
			cx := x + float64(px)*cellSize
			cy := y + float64(py)*cellSize
			p.FillRect(cx, cy, cellSize, cellSize, cColor)
			p.StrokeRect(cx, cy, cellSize, cellSize, stroke, types.FromHexString("#505050"))
			textSize := cellSize * 0.8
			font := "merchantD"
			tw := p.TextWidth(cText, font, textSize)
			tx := cx + (cellSize-tw)/2
			ty := cy + (cellSize-textSize)/2
			p.Text(tx, ty, textSize, font, cText, types.FromHexString("#303030"))
		}
	}
	ps := cellSize * 0.75
	px := x + float64(e.GameState.PlayerX)*cellSize + (cellSize-ps)/2
	py := y + float64(e.GameState.PlayerY)*cellSize + (cellSize-ps)/2
	p.FillRect(px, py, ps, ps, types.FromHexString("#BBBBBBBB"))
}

func (e *GameUIElement) coordsToIndex(mouseX, mouseY float64, p platform.Platform) int {
	x, y, w, h := e.GetBoundary(p)
	offsetX := mouseX - x
	offsetY := mouseY - y
	if offsetX < 0 || offsetX > w || offsetY < 0 || offsetY > h {
		return -1
	}
	gs := int(e.GameState.GridSize)
	cellSize := w / float64(gs)
	cx := int(math.Floor(offsetX / cellSize))
	cy := int(math.Floor(offsetY / cellSize))
	idx := cx*gs + cy
	return idx
}

func getCellTextAndColor(c Cell) (string, types.Color) {
	var text, color string

	if c.Open {
		if c.Empty {
			color = "#BBAACC"
			if c.BombsNear == 0 {
				text = " "
			} else {
				text = fmt.Sprint(c.BombsNear)
			}
		} else {
			text = "X"
			color = "#CCAAAA"
		}
	} else {
		text = " "
		color = "#AAAACC"
	}
	if c.Marked {
		color = "#AACCAA"
		text = "*"
	}

	return text, types.FromHexString(color)
}

func (e *GameUIElement) drawDiedScreen(p platform.Platform) {
	x, y, w, h := e.GetBoundary(p)
	p.FillRect(x, y, w, h, types.FromHexString("#202020AA"))
	tw := p.TextWidth("WASTED", "merchantD", 100)
	tx := x + (w-tw)/2
	ty := y + (h-100)/2
	p.Text(tx, ty, 100, "merchantD", "WASTED", types.FromHexString("#FF2020"))
}

func (e *GameUIElement) drawWinScreen(p platform.Platform) {
	x, y, w, h := e.GetBoundary(p)
	p.FillRect(x, y, w, h, types.FromHexString("#202020AA"))
	tw := p.TextWidth("Congrats", "merchantD", 100)
	tx := x + (w-tw)/2
	ty := y + (h-100)/2
	p.Text(tx, ty, 100, "merchantD", "Congrats", types.FromHexString("#20FF20"))
}
