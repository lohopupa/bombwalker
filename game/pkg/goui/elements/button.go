package elements

import (	
	"minewalker/game/pkg/goui/platform"
)

type Button struct {
	ElementBase
	BorderThick float64
	OnClick func()
}

func NewButton(x, y float64, text string) *Button {
	return &Button{
		BorderThick: 5,
		ElementBase: ElementBase{
			PosX:             x,
			PosY:             y,
			Width:            0,
			Height:           0,
			Text:             text,
			TextSize:         30,
			AlignX:           AlignXLeft,
			AlignY:           AlignYBottom,
			AdjustSizeToText: true,
			Padding:          15,
			ColorScheme:      nil,
		},
	}
}

func (b *Button) Draw(r platform.Platform) {
	x, y, w, h := (*b).GetBoundary(r)
	pc, ac := b.ColorScheme.PrimaryColor, b.ColorScheme.AccentColor
	if b.Hover {
		pc, ac = b.ColorScheme.PrimaryColorHighlight, b.ColorScheme.AccentColorHighlight
	}
	r.FillRect(x, y, w, h, pc)
	r.StrokeRect(x, y, w, h, b.BorderThick, ac)
	tw := r.TextWidth(b.Text, b.FontFamily, b.TextSize)
	tx := x + (w-tw)/2
	ty := y + (h-b.TextSize)/2
	r.Text(tx, ty, b.TextSize, b.FontFamily, b.Text, ac)
}

func (b *Button) GetBoundary(r platform.Platform) (x, y, w, h float64) {
	w = b.Width
	h = b.Height
	x = b.PosX
	y = b.PosY
	if b.AdjustSizeToText {
		w = r.TextWidth(b.Text, b.FontFamily, b.TextSize) + b.Padding*2
		h = b.TextSize + b.Padding*2
	}

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

func (b *Button) HandleEvent(e platform.Event, p platform.Platform) {
	x, y, w, h := b.GetBoundary(p)
	if e.MousePosX > x && e.MousePosX < x+w && e.MousePosY > y && e.MousePosY < y+h {
		if e.EventType == platform.EventTypeMouseClick {
			b.OnClick()
		}
		b.Hover = true
	} else {
		b.Hover = false
	}
}
