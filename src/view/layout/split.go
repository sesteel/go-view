package layout

import (
	"view"
)

// VSplit is a simple layout that expands a single
// child to the size allotted to the layout by the
// target.
type VSplit struct {
	target      view.Composite
	left, right view.Drawer
	division    float64
}

func NewVSplit(target view.Composite) *VSplit {
	l := new(VSplit)
	l.target = target
//	division
	return l
}

func (self *VSplit) SetLeft(d view.Drawer) {
	self.left = d
}

func (self *VSplit) SetRight(d view.Drawer) {
	self.right = d
}

func (self *VSplit) Left() view.Drawer {
	return self.left
}

func (self *VSplit) Right() view.Drawer {
	return self.right
}

func (self *VSplit) Draw(bounds view.Bounds, offset view.ScrollOffset) {
//	s := self.target.Surface()
//	s.Rectangle(float64(0), float64(0), )
//	s.Clip()
//	self.left.Draw()
//	self.right.Draw()
}
