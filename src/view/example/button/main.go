package main 

import (
	. "view/common"
	"fmt"
	"view/widget/button"
	"view"
	"view/layout"
	"view/event"
)
  
func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Button Example", 100, 100, 500, 500)
	abs := layout.NewAbsolute(win)
	
	a := button.New(win, "OK")
	a.AddMouseButtonPressHandler(func(m event.Mouse){
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("OK Pressed") 
		}
	})
	
	b := button.New(win, "Cancel")
	b.AddMouseButtonPressHandler(func(m event.Mouse){
		if m.Button == event.MOUSE_BUTTON_LEFT {
			fmt.Println("Cancel Pressed")
		}
	})
	
	abs.Add(a, Bounds{10, 10, Size{80, 30}})
	abs.Add(b, Bounds{100, 10, Size{80, 30}})
	win.SetLayout(abs)
	win.Start()
	<- waitOnExit  
}

