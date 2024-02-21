package elements

import (
	"minewalker/game/pkg/goui/platform"
)

type Label struct {
	ElementBase
}

func NewLabel(x, y float64, text string) *Label {
	return &Label{
		ElementBase: ElementBase{
			PosX:             x,
			PosY:             y,
			Width:            0,
			Height:           0,
			GetText:          func() string { return text },
			TextSize:         20,
			AlignX:           AlignXLeft,
			AlignY:           AlignYBottom,
			AdjustSizeToText: true,
			Padding:          5,
			ColorScheme:      nil,
		},
	}
}

func (b *Label) Draw(r platform.Platform) {
	x, y, _, _ := (*b).GetBoundary(r)
	r.Text(x, y, b.TextSize, b.FontFamily, b.GetText(), b.ColorScheme.AccentColor)
}

func (b *Label) GetBoundary(r platform.Platform) (x, y, w, h float64) {
	w = r.TextWidth(b.GetText(), b.FontFamily, b.TextSize)
	h = b.TextSize
	x = b.PosX
	y = b.PosY

	switch b.AlignX {
	case AlignXLeft:
		break
	case AlignXCenter:
		x -= w / 2
		break
	case AlignXRight:
		x -= w
		break
	}
	switch b.AlignY {
	case AlignYBottom:
		break
	case AlignYCenter:
		y -= h / 2
		break
	case AlignYTop:
		y -= h
	}
	return x, y, w, h
}

func (b *Label) HandleEvent(platform.Event, platform.Platform) {

}
