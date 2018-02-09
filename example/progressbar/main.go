package main

import (
	"time"
	"view"
	. "view/common"
	"view/layout"
	"view/widget/progressbar"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Progress Bar Example", 100, 100, 400, 50)
	abs := layout.NewAbsolute(win)
	a := progressbar.New(win, "OK", 500)
	abs.Add(a, Bounds{Point{10, 10}, Size{380, 25}})
	win.SetLayout(abs)
	win.Start()
	for i := 0.0; i <= 501.0; i++ {
		a.SetValue(i)
		a.Redraw()
		time.Sleep(10 * time.Millisecond)
		if i >= 500 {
			i = 0
			time.Sleep(1000 * time.Millisecond)
		}
	}
	<-waitOnExit
}
