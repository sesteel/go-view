package view

import (
	"fmt"
	"ui/event"
)

type Composite interface {
	View
	SetLayout(Layout)
	GetLayout() Layout
}

type CompositeView struct {
	DefaultView
	layout  Layout
	event.EventDispatcher
}

func (self *CompositeView) Parent() View {
	return self.parent
}

func (self *CompositeView) Position() (uint, uint) {
	return self.x, self.y
}

func (self *CompositeView) Surface() *Surface {
	return self.surface
}

func (self *CompositeView) SetLayout(layout Layout) {
	self.layout = layout
}

func (self *CompositeView) GetLayout() Layout {
	return self.layout
}

func (self *CompositeView) SetText(text string) {
	self.text = text
}

func (self *CompositeView) Text() string {
	return self.text
}

func (self *CompositeView) SetSize(width, height uint) {
	self.width = width
	self.height = height
}

func (self *CompositeView) Size() (uint, uint) {
	return self.width, self.height
}

func (self *CompositeView) Draw() {
	// 1. save state
	// 2. create mask
	// 3. translate
	// 4. draw
	// 5. apply
	// 6. translate back
	// 7. pop
	if self.layout != nil {
	fmt.Println("DRAW")
		self.layout.Draw()
	}
}
