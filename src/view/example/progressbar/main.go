package main 

import (
	"view/widget/progressbar"
	"view"
	"view/layout"
	"time"
)

func main() {
	var waitOnExit chan bool  
	win := view.NewWindow("Progress Bar Example", 100, 100, 400, 50)
	abs := layout.NewAbsolute(win)
	a := progressbar.New(win, "OK", 500)
	abs.Add(a, view.Bounds{10, 10, view.Size{380, 25}})
	win.SetLayout(abs)
	win.Start()
	for i:=0.0; i <= 501.0; i++ {
		a.SetValue(i)
		a.Redraw()
		time.Sleep(10 * time.Millisecond)
		if (i >= 500) {
			i = 0
			time.Sleep(1000 * time.Millisecond)
		}
	}
	<- waitOnExit  
}
