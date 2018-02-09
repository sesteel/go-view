package main

import (
	// "fmt"
	"view"
	"view/color"
	// "view/event"
	"view/layout"
	"view/widget/editor"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Grid Layout Example", 100, 100, 600, 300)
	grid := layout.NewGrid(win)

	// a := label.New(win, "label1", "This is a label.")
	// a.Style().SetBackground(color.Red2)
	// a.AddMouseButtonPressHandler(func(m event.Mouse) {
	// 	if m.Button == event.MOUSE_BUTTON_LEFT {
	// 		fmt.Println("OK Pressed")
	// 	}
	// })

	// b := label.New(win, "label2", "This too is a another label.")
	// b.Style().SetBackground(color.Green2)
	// b.AddMouseButtonPressHandler(func(m event.Mouse) {
	// 	if m.Button == event.MOUSE_BUTTON_LEFT {
	// 		fmt.Println("Cancel Pressed")
	// 	}
	// })

	// c := label.New(win, "label3", "This label is centered. \n asdasd")
	// c.Style().SetTextAlignment(view.STYLE_TEXT_CENTERED)
	// c.Style().SetBackground(color.Blue2)
	// c.AddMouseButtonPressHandler(func(m event.Mouse) {
	// 	if m.Button == event.MOUSE_BUTTON_LEFT {
	// 		fmt.Println("Cancel Pressed")
	// 	}
	// })

	d := editor.New(win, "editor", "a, \nb, \nc")
	d.Style.Background = color.Blue2
	// grid.Add(a, 0, 0, 4, 1)
	// grid.Add(b, 1, 0, 1, 1)
	// grid.Add(c, 2, 0, 1, 1)
	grid.Add(d, 0, 0, 2, 2)
	grid.Add(d, 3, 0, 2, 2)
	win.SetLayout(grid)
	win.Start()
	<-waitOnExit
}
