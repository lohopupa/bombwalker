package main

import (
	"minewalker/game/pkg/platform"
	"minewalker/game/ui"
	// "time"
)

func main() {
	platform := platform.InitWasm("canvas")
	w := ui.CreateMainPage(platform)
	w.Draw()
	// w.Stop()
	select {}
}
