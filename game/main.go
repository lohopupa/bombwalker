package main

import (
	"log"
	"minewalker/gojs"
)

func main(){
	x, err := gojs.Init()
	if err != nil {
		log.Fatal(err)
	}
	htmlCanvas, err := x.GetElementById("canvas")
	if err != nil {
		log.Fatal(err)
	}
	canvas := htmlCanvas.ToCanvas()
	canvas.HtmlElement = canvas.HtmlElement

}