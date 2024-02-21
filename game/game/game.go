package game

import (
	"fmt"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/window"
)

type Game struct {
	CurrentWindow string
	Windows       map[string]*window.Window
	Score uint
	GridSize uint
	BombsTotal uint
	BombsOpen uint
}

func InitGameState(p platform.Platform) Game {
	g := Game{CurrentWindow: "hello"}
	g.Windows = make(map[string]*window.Window)
	g.Windows["hello"] = createHelloWindow(p, &g)
	g.Windows["main"] = createMainWindow(p, &g)
	g.Windows["help"] = createHelpWindow(p, &g)
	g.Windows["shit"] = createShitWindow(p, &g)
	g.BombsOpen = 5
	g.BombsTotal = 10
	g.GridSize = 10
	g.Score = 69
	return g
}

func (g Game) Start() {
	w := g.Windows[g.CurrentWindow]
	w.Draw()
}

func (g *Game) ChangeWindow(windowName string) {
	// TODO: After window changes next event is loss
	fmt.Println(windowName)
	if window := g.Windows[windowName]; window != nil {
		cw := g.Windows[g.CurrentWindow]
		cw.Stop()
		g.CurrentWindow = windowName
		window.Draw()
	} else {
		fmt.Printf("Could not find window %s\n", windowName)
	}
}
