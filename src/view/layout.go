package view

import (
	"view/event"
)

// Layouts are special types that provide strategies
// for dividing the view space among components.  
// They are charged with several important responsibilities:
//
// - They mask or otherwise protect the rest of the 
//   target view surface from unintended drawing.
//
// - They define the boundries from which a component
//   or other composite view can be rendered
//
// - They distribute events to their child components
// 
type Layout interface {
	Drawer
	event.MouseHandler
}