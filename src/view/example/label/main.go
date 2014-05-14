package main

import (
	"fmt"
	"view"
	"view/color"
	"view/event"
	"view/layout"
	"view/widget/label"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Label Example", 100, 100, 200, 200)
	abs := layout.NewAbsolute(win)

	a := label.New(win, "label1", "This is a label.")
	a.AddMouseButtonPressHandler(func(m event.Mouse) {
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("OK Pressed")
		}
	})

	b := label.New(win, "label2", "This too is a another label.")
	b.Style().SetBorderColor(color.Gray11)
	b.AddMouseButtonPressHandler(func(m event.Mouse) {
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("Cancel Pressed")
		}
	})

	c := label.New(win, "label3", "This label is centered. \n asdasd")
	c.Style().SetTextAlignment(view.STYLE_TEXT_CENTERED)
	c.Style().SetBorderColor(color.Blue1)
	c.AddMouseButtonPressHandler(func(m event.Mouse) {
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("Cancel Pressed")
		}
	})

	abs.Add(a, view.Bounds{10, 10, view.Size{180, 30}})
	abs.Add(b, view.Bounds{10, 40, view.Size{180, 30}})
	abs.Add(c, view.Bounds{10, 70, view.Size{180, 30}})
	//	abs.Add(d, view.Bounds{10, 70, view.Size{180, 30}})
	//	abs.Add(e, view.Bounds{10, 90, view.Size{180, 30}})
	win.SetLayout(abs)
	win.Start()
	<-waitOnExit
}
