// +build linux,!goci
package view

import (
	"view/event"
	"fmt"
)

type Drawer interface {
	Draw(*Surface) //Bounds, ScrollOffset)
	Redraw()
}

type View interface {
	SetParent(View)
	Parent() View
	Surface() *Surface
	SetName(string)
	Name() string
	Position() (float64, float64)
	SetSize(float64, float64)
	Size() (float64, float64)
	SetStyle(Style)
	Style() Style
	Drawer
	event.FocusNotifier
	event.FocusHandler
	event.MouseNotifier
	event.MouseHandler
}

type DefaultView struct {
	parent  View
	surface *Surface
	name    string
	width   float64
	height  float64
	x, y    float64
	focus   bool
	style   Style
}

func (self *DefaultView) SetParent(parent View) {
	self.parent = parent
}

func (self *DefaultView) Parent() View {
	return self.parent
}

func (self *DefaultView) Position() (float64, float64) {
	return self.x, self.y
}

func (self *DefaultView) Surface() *Surface {
	return self.surface
}

func (self *DefaultView) SetName(name string) {
	self.name = name
}

func (self *DefaultView) Name() string {
	return self.name
}

func (self *DefaultView) SetSize(width, height float64) {
	self.width = width
	self.height = height
}

func (self *DefaultView) Size() (float64, float64) {
	return self.width, self.height
}

func (self *DefaultView) SetStyle(style Style) {
	self.style = style
}

func (self *DefaultView) Style() Style {
	return self.style
}

func (self *DefaultView) Draw(surface *Surface) {
	// default drawing does here
}

func (self *DefaultView) Redraw() {
	if DEBUG {
		fmt.Println("View.Redraw()")
	}
	if self.parent != nil {
		self.parent.Redraw()
	}
}