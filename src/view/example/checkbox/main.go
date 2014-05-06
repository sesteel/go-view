package main

import (
	"view"
	"view/layout"
	"view/widget/checkbox"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Check Box Example", 100, 100, 400, 70) 
	abs := layout.NewAbsolute(win)
	a := checkbox.New(win, "checkbox1", "I agree to the terms and conditions.")
	b := checkbox.New(win, "checkbox2", "Please automatically send me software updates.")
	abs.Add(a, view.Bounds{10, 10, view.Size{380, 25}})
	abs.Add(b, view.Bounds{10, 40, view.Size{380, 25}})
	win.SetLayout(abs)
	win.Start()
	<-waitOnExit
}
