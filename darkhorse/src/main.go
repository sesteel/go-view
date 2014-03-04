package main 

import (
	"view/core"    
)

func main() {
	var waitOnExit chan bool
	win := core.NewWindow("Test Application", 0, 0, 800, 600)
	win.SetName("XXXXXXX")
	win.SetSize(1000,300)
	<-waitOnExit 
}

