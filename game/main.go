package main

import (
	"minewalker/game/game"
	"minewalker/game/pkg/platform"
	// "time"
)

func main() {
	platform := platform.InitWasm("canvas")
	game := game.InitGameState(platform)
	game.Start()
	// w.Draw()
	// w.Stop()
	select {}
}
