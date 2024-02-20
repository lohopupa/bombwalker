package gojs

import (
	"syscall/js"
)

// TODO: Adapt to js 
// type UIEvent struct {}

const (
	KeyDownEvent = "keydown"
	KeyUpEvent = "keyup"
	KeyPressEvent = "keypress"

	MouseDownEvent = "mousedown"
	MouseUpEvent = "mouseup"
	MouseClickEvent = "click"
	MouseDblClickEvent = "dblclick"
	MouseMove = "mousemove"
	// TODO: Add support for other events
	// MouseEnter = "mouseenter"
	// MouseLeave = "mouseleave"
	// MouseOut = "mouseout"
	// MouseOver = "mouseover"
)

type Event interface {
}
// TODO: Add support for gestures
type KeyboardEvent struct {
	EventType string
	AltKey bool
	CtrlKey bool
	ShiftKey bool
	MetaKey bool
	CharCode uint8
	Code string
	Key string
	KeyCode int
}

func KeyboardEventFromArgs(args []js.Value) KeyboardEvent {
	arg := args[0]
	return KeyboardEvent{
		EventType: arg.Get("type").String(),
		AltKey: arg.Get("altKey").Bool(),
		CtrlKey: arg.Get("ctrlKey").Bool(),
		ShiftKey: arg.Get("shiftKey").Bool(),
		MetaKey: arg.Get("metaKey").Bool(),
		CharCode: uint8(arg.Get("charCode").Int()),
		Code: arg.Get("code").String(),
		Key: arg.Get("key").String(),
		KeyCode: arg.Get("keyCode").Int(),
	}
}

type MouseEvent struct {
	EventType string
	AltKey bool
	CtrlKey bool
	ShiftKey bool
	MetaKey bool
	Button int
	ClientX int
	ClientY int
	OffsetX int
	OffsetY int
}


func MouseEventFromArgs(args []js.Value) MouseEvent {
	arg := args[0]
	return MouseEvent{
		EventType: arg.Get("type").String(),
		// AltKey: arg.Get("altkey").Bool(),
		// CtrlKey: arg.Get("ctrlKey").Bool(),
		// ShiftKey: arg.Get("shiftKey").Bool(),
		// MetaKey: arg.Get("metaKey").Bool(),
		Button: arg.Get("button").Int(),
		ClientX: arg.Get("clientX").Int(),
		ClientY: arg.Get("clientY").Int(),
		OffsetX: arg.Get("offsetX").Int(),
		OffsetY: arg.Get("offsetY").Int(),
	}
}