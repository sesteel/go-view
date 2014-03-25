package view

import (
	"ui/event"
//	"ui/view/color"
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

func (self *CompositeView) SetText(text string) {
	self.text = text
}

func (self *CompositeView) Text() string {
	return self.text
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
	
	if self.layout != nil {
		self.layout.Draw(surface)
	}
}
