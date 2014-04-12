package view

import (
	"ui/view/event"
)

type Composite interface {
	View
	SetLayout(Layout)
	GetLayout() Layout
}

type CompositeView struct {
	DefaultView
	layout  Layout
	event.MouseEventDispatcher
}

func (self *CompositeView) Parent() View {
	return self.parent
}

func (self *CompositeView) Position() (float64, float64) {
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

func (self *CompositeView) SetSize(width, height float64) {
	self.width = width
	self.height = height
}

func (self *CompositeView) Size() (float64, float64) {
	return self.width, self.height
}

func (self *CompositeView) Draw(surface *Surface) {
	// 1. save state
	// 2. create clip
	// 3. translate
	// 4. draw
	// 5. apply
	// 6. translate back
	// 7. pop
//	x, y := self.Position()
//	w, h := self.Size()
//	p := NewLinearPattern(x, y, x, h)
//	p.AddColorStop(0, color.Gray3)
//	p.AddColorStop(1, color.Gray3)
//	surface.Rectangle(x, y, w, h)
//	surface.SetSource(p)
//	surface.Fill()
//	p.Destroy()
	
	self.layout.Draw(surface)
	
}

func (self *CompositeView) MouseEnter(me event.Mouse) {
	self.MouseEventDispatcher.MouseEnter(me)
	self.layout.MouseEnter(me)
}

func (self *CompositeView) MouseExit(me event.Mouse) {
	self.MouseEventDispatcher.MouseExit(me)
	self.layout.MouseExit(me)
}

func (self *CompositeView) MousePosition(me event.Mouse) {
	self.MouseEventDispatcher.MousePosition(me)
	self.layout.MousePosition(me)
}

func (self *CompositeView) MouseWheelUp(me event.Mouse) {
	self.MouseEventDispatcher.MouseWheelUp(me)
	self.layout.MouseWheelUp(me)
}

func (self *CompositeView) MouseWheelDown(me event.Mouse) {
	self.MouseEventDispatcher.MouseWheelDown(me)
	self.layout.MouseWheelDown(me)
}

func (self *CompositeView) MouseButtonPress(me event.Mouse) {
	self.MouseEventDispatcher.MouseButtonPress(me)
	self.layout.MouseButtonPress(me)
}

func (self *CompositeView) MouseButtonRelease(me event.Mouse) {
	self.MouseEventDispatcher.MouseButtonRelease(me)
	self.layout.MouseButtonRelease(me)
}
