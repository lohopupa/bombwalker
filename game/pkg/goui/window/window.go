package window

import (
	"minewalker/game/pkg/goui/elements"
	"minewalker/game/pkg/goui/platform"
	"minewalker/game/pkg/goui/types"
)

type Window struct {
	Title       string
	SizeX       float64
	SizeY       float64
	Elements    []*elements.Element
	Platform    *platform.Platform
	ColorScheme types.ColorScheme
	FontFamily  string
	stopChan    *chan bool
}

func NewWindow(title string, r platform.Platform) *Window {
	x, y := r.GetSize()
	stopch := make(chan bool, 10)
	return &Window{
		Title:       title,
		SizeX:       x,
		SizeY:       y,
		ColorScheme: types.DefaultColorScheme(),
		Platform:    &r,
		stopChan:    &stopch,
	}
}

func (w *Window) AddElement(e elements.Element) {
	e.SetColors(w.ColorScheme)
	e.SetFontFamily(w.FontFamily)
	w.Elements = append(w.Elements, &e)
}

func (w *Window) AddElements(es... elements.Element) {
	for _, e := range es {
		w.AddElement(e)
	}
}

func (w *Window) Draw() {
	draw := func(int) {
		(*w.Platform).ClearRect(0, 0, w.SizeX, w.SizeY)
		(*w.Platform).FillRect(0, 0, w.SizeX, w.SizeY, w.ColorScheme.PrimaryColor)
		for _, e := range w.Elements {
			(*e).Draw(*w.Platform)
		}
	}
	go (*w.Platform).StartRendering(draw)
	go w.HandleEvents()

}

func (w *Window) HandleEvents() {
	for event := range (*w.Platform).GetEventsChan() {
		select {
		case <-*w.stopChan:
			return
		default:
			{
				for _, element := range w.Elements {
					(*element).HandleEvent(event, *w.Platform)
				}
			}
		}
	}
}

func (w *Window) Stop() {
	*w.stopChan <- true
	(*w.Platform).StopRendering()
}
