package main

import (
	"ui/view"
	"ui/view/layout"
	"ui/view/widget"
	"fmt"
)

func main() {
	var waitOnExit chan bool
	fmt.Println(view.HexColor(0xff77ff00))
	win := view.NewWindow("Test Application", 0, 0, 800, 600)
	win.SetText("Test App")
	win.SetSize(1000, 300)
	l := layout.NewXYPositioned(win)
	l.Add(widget.NewButton(win, "ABCdefg"), 10, 10)
	win.SetLayout(l)
	<-waitOnExit
}
