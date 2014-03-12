// +build linux,!goci
package view

import (
	"ui/event"
)

type Drawer interface {
	Draw()
}

type View interface {
	Parent() View
	Surface() *Surface
	SetText(string) 
	Text() string
	Position() (uint, uint)
	SetSize(uint, uint) 
	Size() (uint, uint)
	SetStyle(*Style) 
	Style() *Style
	Drawer
	AddMouseEnterHandler(func(event.MouseEnter))
	AddMouseExitHandler(func(event.MouseExit))
}

type DefaultView struct {
	parent             View
	surface            *Surface
	text               string
	width              uint
	height             uint
	x, y               uint
	focus              bool
	style              *Style
	event.EventDispatcher
}

func (self *DefaultView) Parent() View {
	return self.parent
}

func (self *DefaultView) Position() (uint, uint) {
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

func (self *DefaultView) SetSize(width, height uint) {
	self.width = width
	self.height = height
}

func (self *DefaultView) Size() (uint, uint) {
	return self.width, self.height
}

func (self *DefaultView) SetStyle(style *Style) {
	self.style = style
}

func (self *DefaultView) Style() *Style {
	return self.style
}

func (self *DefaultView) Draw() {
	// default drawing does here
}
