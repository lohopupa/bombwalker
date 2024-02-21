package game

import (
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/types"
)
type GameUIElement struct {
	PosX, PosY, SizeX, SizeY float64 
}

func (e *GameUIElement) SetColors(cs types.ColorScheme) {
}

func (e *GameUIElement) SetFontFamily(ff string) {
}

func (e *GameUIElement) Draw(p platform.Platform) {
	p.FillRect(e.PosX, e.PosY, e.SizeX, e.SizeY, types.FromHexString("#181818"))
	p.FillCircle(e.PosX + 200, e.PosY + 425, 120, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 200, e.PosY + 275, 120, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 300, e.PosY + 350, 100, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 350, e.PosY + 350, 100, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 400, e.PosY + 350, 100, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 450, e.PosY + 350, 100, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 500, e.PosY + 350, 100, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 550, e.PosY + 350, 100, types.FromHexString("#DB9065"))
	p.FillCircle(e.PosX + 600, e.PosY + 350, 120, types.FromHexString("#FF3030"))
}
func (e *GameUIElement) GetBoundary(platform.Platform) (float64, float64, float64, float64) {
	return e.PosX, e.PosY, e.SizeX, e.SizeY
}

func (e *GameUIElement) HandleEvent(pe platform.Event, p platform.Platform) {
	
}