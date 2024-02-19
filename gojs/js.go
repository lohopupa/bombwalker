package gojs

import (
	"errors"
	"syscall/js"
)

type JS struct {
	document js.Value
}

func Init() (JS, error) {
	js := JS{
		document: js.Global().Get("document"),
	}
	return js, nil
}

func (this JS) GetElementById(id string) (*HtmlElement, error) {
	jsElement := this.document.Call("getElementById", id)
	if jsElement.IsNull() {
		return nil, errors.New("Could not get element " + id)
	}
	return &HtmlElement{
		instance: jsElement,
		id:       id,
	}, nil
}
