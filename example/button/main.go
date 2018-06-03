package main

import (
	"fmt"

	"github.com/sesteel/go-view"
	"github.com/sesteel/go-view/event"
	"github.com/sesteel/go-view/geometry"
	"github.com/sesteel/go-view/layout"
	"github.com/sesteel/go-view/widget/button"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Button Example", 100, 100, 500, 500)
	abs := layout.NewAbsolute(win)

	a := button.New(win, "OK")
	a.AddMouseButtonPressHandler(func(m event.Mouse) {
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("OK Pressed")
		}
	})

	b := button.New(win, "Cancel")
	b.AddMouseButtonPressHandler(func(m event.Mouse) {
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("Cancel Pressed")
		}
	})

	abs.Add(a, geometry.Bounds{geometry.Point{10, 10}, geometry.Size{180, 60}})
	abs.Add(b, geometry.Bounds{geometry.Point{190, 10}, geometry.Size{180, 60}})
	win.SetLayout(abs)
	win.Start()
	<-waitOnExit
}
