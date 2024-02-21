package elements

import (
	"minewalker/game/pkg/goui/platform"
	"strings"
)

type AlignText uint8

const (
	AlignTextLeft   AlignText = iota
	AlignTextCenter AlignText = iota
	AlignTextRight  AlignText = iota
)

type LabelML struct {
	ElementBase
	AlignText AlignText
	TextPad   float64
}

func NewLabelML(x, y, w float64, text string) *LabelML {
	return &LabelML{
		ElementBase: ElementBase{
			PosX:             x,
			PosY:             y,
			Width:            w,
			Height:           0,
			GetText:          func() string { return text },
			TextSize:         20,
			AlignX:           AlignXLeft,
			AlignY:           AlignYBottom,
			AdjustSizeToText: true,
			Padding:          5,
			ColorScheme:      nil,
		},
		AlignText: AlignTextCenter,
		TextPad:   10,
	}
}

func (b *LabelML) Draw(r platform.Platform) {
	x, y, w, _ := (*b).GetBoundary(r)
	lines := splitTextOnChunks(b.GetText(), b.FontFamily, b.TextSize, b.Width, r)
	for i, l := range lines {
		tw := r.TextWidth(l, b.FontFamily, b.TextSize)
		tx := x
		switch b.AlignText {
		case AlignTextCenter:
			tx = x + (w-tw)/2
		case AlignTextRight:
			tx = x + (w - tw)
		}
		r.Text(tx, y+(b.TextSize+b.TextPad)*float64(i), b.TextSize, b.FontFamily, l, b.ColorScheme.AccentColor)
	}
}

func (b *LabelML) GetBoundary(r platform.Platform) (x, y, w, h float64) {
	lines := splitTextOnChunks(b.GetText(), b.FontFamily, b.TextSize, b.Width, r)
	w = b.Width
	h = (b.TextSize + b.TextPad) * float64(len(lines))
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

func (b *LabelML) HandleEvent(platform.Event, platform.Platform) {

}

// func splitTextOnChunks(text, fontFamily string, fontSize, width float64, p platform.Platform) []string {
// 	words := strings.Split(text, " ")
// 	result := make([]string, 1)
// 	for _, w := range words {
// 		t := result[len(result)-1] + " " + w
// 		if p.TextWidth(t, fontFamily, fontSize) < width {
// 			result[len(result)-1] += " " + w
// 		} else {
// 			result = append(result, w)
// 		}
// 	}
// 	return result
// }

func splitTextOnChunks(text, fontFamily string, fontSize, width float64, p platform.Platform) []string {
	words := strings.Split(text, " ")
	result := make([]string, 0)
	for _, w := range words {
		if len(result) == 0 {
			result = append(result, w)
			continue
		}

		t := result[len(result)-1] + " " + w
		if p.TextWidth(t, fontFamily, fontSize) < width {
			result[len(result)-1] += " " + w
		} else {
			result = append(result, w)
		}
	}
	return result
}
