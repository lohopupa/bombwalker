package game

import (
	"fmt"
	"math/rand"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/window"
)

type Game struct {
	CurrentWindow   string
	Windows         map[string]*window.Window
	Score           uint
	GridSize        uint
	BombsTotal      uint
	BombsOpen       uint
	Map             []*Cell
	Alive           bool
	Win             bool
	PlayerX         uint
	PlayerY         uint
	PlayerDirection uint8 // 0, 1, 2, 3 top, right, down, left
}

type Cell struct {
	Open      bool
	Empty     bool
	Hover     bool
	Marked    bool
	BombsNear uint
}

func InitGameState(p platform.Platform) Game {
	g := Game{CurrentWindow: "hello"}
	g.Windows = make(map[string]*window.Window)
	g.Windows["hello"] = createHelloWindow(p, &g)
	g.Windows["main"] = createMainWindow(p, &g)
	g.Windows["help"] = createHelpWindow(p, &g)
	g.Windows["shit"] = createShitWindow(p, &g)
	g.BombsOpen = 0
	g.BombsTotal = 200
	g.GridSize = 40
	g.Score = g.GridSize * g.GridSize
	g.Alive = true
	g.generateMap()
	return g
}

func (g *Game) generateMap() {
	var cellCount int = int(g.GridSize * g.GridSize)
	g.Map = make([]*Cell, cellCount)
	for i := range g.Map {
		g.Map[i] = &Cell{
			Empty: true,
		}
	}
	var bombsSeted uint
	for bombsSeted = 0; bombsSeted < g.BombsTotal; {
		idx := rand.Intn(cellCount)
		if g.Map[idx].Empty {
			g.Map[idx].Empty = false
			bombsSeted += 1
			g.updateNeighbors(idx)
		}

	}
}
func (g *Game) updateNeighbors(idx int) {
	gs := int(g.GridSize)
	lx := -1
	rx := 1
	ly := -1
	ry := 1
	if idx%gs == 0 {
		lx = 0
	} else if idx%gs == gs-1 {
		rx = 0
	}
	if idx/gs == 0 {
		ly = 0
	} else if idx/gs == gs-1 {
		ry = 0
	}
	for x := lx; x <= rx; x += 1 {
		for y := ly; y <= ry; y += 1 {
			i := idx + x + y*gs
			if i >= 0 && i < gs*gs {
				g.Map[i].BombsNear += 1
			}
		}
	}
}

func (g *Game) Start() {
	w := g.Windows[g.CurrentWindow]
	w.Draw()
}

func (g *Game) ChangeWindow(windowName string) {
	//TODO: After window changes next event is loss
	if window := g.Windows[windowName]; window != nil {
		cw := g.Windows[g.CurrentWindow]
		cw.Stop()
		g.CurrentWindow = windowName
		window.Draw()
	} else {
		fmt.Printf("Could not find window %s\n", windowName)
	}
}

func (g *Game) Restart() {
	g.generateMap()
	g.Score = g.GridSize * g.GridSize
	g.Alive = true
	g.Win = false
	g.BombsOpen = 0
	g.PlayerX = 0
	g.PlayerY = 0

}

func (g *Game) Die() {
	g.Alive = false
	for _, c := range g.Map {
		c.Open = true
		if c.Marked && c.Empty {
			c.Marked = false
			g.Score -= 1
		}
	}
}

func (g *Game) CheckWin() bool {
	if !g.Alive {
		return false
	}
	oc := 0
	mc := 0
	for _, c := range g.Map {
		if c.Open && c.Empty {
			oc += 1
		}
		if c.Marked {
			if c.Empty {
				mc -= 1
			} else {
				mc += 1
			}
		}
	}
	if oc == int(g.GridSize*g.GridSize-g.BombsTotal) || mc == int(g.BombsTotal) {
		g.Win = true
		return true
	}
	return false
}

func (g *Game) OpenCell(idx int) {
	cell := g.Map[idx]
	if cell.Empty {
		if !cell.Open {
			g.OpenArea(idx, 20)
			cell.Open = true
			g.Score -= 1
			if cell.Marked {
				cell.Marked = false
				g.BombsOpen -= 1
			}
		}
	} else {
		g.Die()
	}
	g.CheckWin()
}

func (g *Game) MarkCell(idx int) {
	cell := g.Map[idx]
	if !cell.Open {
		cell.Marked = !cell.Marked
		if cell.Marked {
			g.BombsOpen += 1
		} else {
			g.BombsOpen -= 1

		}
	}
	g.CheckWin()
}

func (g *Game) OpenArea(idx, depth int) {
	cell := g.Map[idx]
	if !cell.Empty || cell.Open || cell.Marked || depth == 0 {
		return
	}
	if cell.BombsNear != 0 {
		cell.Open = true
		// g.Score -= 1
		return
	}
	cell.Open = true
	g.Score -= 1
	gs := int(g.GridSize)
	lx := -1
	rx := 1
	ly := -1
	ry := 1
	if idx%gs == 0 {
		lx = 0
	} else if idx%gs == gs-1 {
		rx = 0
	}
	if idx/gs == 0 {
		ly = 0
	} else if idx/gs == gs-1 {
		ry = 0
	}
	for x := lx; x <= rx; x += 1 {
		for y := ly; y <= ry; y += 1 {
			i := idx + x + y*gs
			if i >= 0 && i < gs*gs {
				// if cell.BombsNear != 0 {
					g.OpenArea(i, depth-1)
				// }
			}
		}
	}

}
