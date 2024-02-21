package platform

import (
	"minewalker/game/pkg/goui/types"
)

type Platform interface {
	ClearRect(x, y, width, height float64)
	FillRect(x, y, width, height float64, color types.Color)
	StrokeRect(x, y, width, height, thick float64, color types.Color)
	FillCircle(x, y, radius float64, color types.Color)
	StrokeCircle(x, y, radius, thick float64, color types.Color)
	Text(x, y, textSize float64, fontFamily, text string, color types.Color)
	TextWidth(text, fontFamily string, fontSize float64) float64
	Line(sX, sY, eX, eY, thickness float32, color types.Color)
	GetSize() (float64, float64)
	GetEventsChan() EventsChan
	StartRendering(func(int))
	StopRendering()
}

type EventType uint8
type MouseButton uint8
type EventsChan chan Event

const (
	EventTypeMouseClick EventType = iota
	EventTypeMouseMove
	EventTypeKeyPress

	MouseLeft EventType = iota
	MouseWheel
	MouseRight
)

type Event struct {
	EventType   EventType
	MouseButton MouseButton
	MousePosX   float64
	MousePosY   float64

	KeyCode string
}
