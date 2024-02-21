package gojs

import (
	"errors"
	// "fmt"
	"syscall/js"
)

func GetElementById(id string) (*HtmlElement, error) {
	jsElement := js.Global().Get("document").Call("getElementById", id)
	if jsElement.IsNull() {
		return nil, errors.New("Could not get element " + id)
	}
	return &HtmlElement{
		instance: jsElement,
		id:       id,
	}, nil
}

type AnimationCallback = func(timestamp int)

func RequestAnimationFrame(callback AnimationCallback) uint64 {
	js.Global().Call("requestAnimationFrame", js.FuncOf(func(this js.Value, args []js.Value) any {
		timeStamp := args[0].Int()
		go callback(timeStamp)
		return nil
	}))
	return 0
}

func AddEventListener(eventType string, callback func(Event)) {
	js.Global().Get("document").Call("addEventListener", eventType, js.FuncOf(func(this js.Value, args []js.Value) any {
		switch eventType {
		case KeyDownEvent:
			fallthrough
		case KeyPressEvent:
			fallthrough
		case KeyUpEvent:
			go callback(KeyboardEventFromArgs(args))
			break
		case MouseClickEvent:
			fallthrough
		case MouseDblClickEvent:
			fallthrough
		case MouseDownEvent:
			fallthrough
		case MouseUpEvent:
			fallthrough
		case MouseMoveEvent:
			go callback(MouseEventFromArgs(args))
			break
		}
		return nil
	}))
}

func (e HtmlElement) AddEventListener(eventType string, callback func(Event)) {
	e.instance.Call("addEventListener", eventType, js.FuncOf(func(this js.Value, args []js.Value) any {
		switch eventType {
		case KeyDownEvent:
			fallthrough
		case KeyPressEvent:
			fallthrough
		case KeyUpEvent:
			go callback(KeyboardEventFromArgs(args))
			break
		case MouseClickEvent:
			fallthrough
		case MouseDblClickEvent:
			fallthrough
		case MouseDownEvent:
			fallthrough
		case MouseUpEvent:
			fallthrough
		case MouseMoveEvent:
			go callback(MouseEventFromArgs(args))
			break
		}
		return nil
	}))
}
