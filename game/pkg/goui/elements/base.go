package elements

import (
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/types"
)

type Element interface {
	Draw(platform.Platform)
	GetBoundary(platform.Platform) (float64, float64, float64, float64)
	SetColors(types.ColorScheme)
	SetFontFamily(string)
	HandleEvent(platform.Event, platform.Platform)
}
type AlignX uint8
type AlignY uint8

const (
	AlignXLeft AlignX = iota
	AlignXRight
	AlignXCenter
	AlignYTop AlignY = iota
	AlignYBottom
	AlignYCenter
)

type ElementBase struct {
	PosX             float64
	PosY             float64
	Width            float64
	Height           float64
	GetText          func() string
	TextSize         float64
	FontFamily       string
	AlignX           AlignX
	AlignY           AlignY
	AdjustSizeToText bool
	Padding          float64
	ColorScheme      *types.ColorScheme
	Hover            bool
}

func (e *ElementBase) SetColors(cs types.ColorScheme) {
	if e.ColorScheme == nil {
		e.ColorScheme = &cs
	}
}

func (e *ElementBase) SetFontFamily(ff string) {
	e.FontFamily = ff
}
