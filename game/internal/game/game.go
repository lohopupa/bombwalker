package game

import (
	la "minewalker/game/internal/basics/linear_algebra"
	"minewalker/game/internal/renderer"
	"minewalker/game/internal/ui"
	"sync"
)

type GameContext struct {
	PrevTime int
	Renderer renderer.Renderer
	Pages    map[string]ui.Page
	MousePos la.Vector2[float64]
}

var instance *GameContext
var once sync.Once

func GetGameContext() *GameContext {
	once.Do(func() {
		instance = &GameContext{}
	})
	return instance
}
