// +build linux,!goci
package view

import (
	"ui/event"
)

type Drawer interface {
	Draw(*Surface)//Bounds, ScrollOffset)
}

type View interface {
	Parent() View
	Surface() *Surface
	SetText(string) 
	Text() string
	Position() (float64, float64)
	SetSize(float64, float64) 
	Size() (float64, float64)
	SetStyle(Style) 
	Style() Style
	Drawer
	AddMouseEnterHandler(func(event.MouseEnter))
	AddMouseExitHandler(func(event.MouseExit))
}

type DefaultView struct {
	parent             View
	surface            *Surface
	text               string
	width              float64
	height             float64
	x, y               float64
	focus              bool
	style              Style
	event.EventDispatcher
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

func (self *DefaultView) SetText(text string) {
	self.text = text
}

func (self *DefaultView) Text() string {
	return self.text
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
