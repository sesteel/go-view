package main

import (
	"github.com/sesteel/go-view"
	. "github.com/sesteel/go-view/geometry"
	"github.com/sesteel/go-view/layout"
	"github.com/sesteel/go-view/widget/checkbox"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Check Box Example", 100, 100, 400, 70)
	abs := layout.NewAbsolute(win)
	a := checkbox.New(win, "checkbox1", "I agree to the terms and conditions.")
	b := checkbox.New(win, "checkbox2", "Please automatically send me software updates.")
	abs.Add(a, Bounds{Point{10, 10}, Size{380, 25}})
	abs.Add(b, Bounds{Point{10, 40}, Size{380, 25}})
	win.SetLayout(abs)
	win.Start()
	<-waitOnExit
}
