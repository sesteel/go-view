package layout

import (
	"view"
)

// Fill is a simple layout that expands a single
// child to the size allotted to the layout by 
// the target.
type Fill struct {
	target view.Composite
	child  view.Drawer
}

func NewFill(target view.Composite) *Fill {
	l := new(Fill)
	l.target = target
	return l
}

func (self *Fill) SetChild(d view.Drawer) {
	self.child = d
}

func (self *Fill) Child() view.Drawer {
	return self.child 
}

func (self *Fill) Draw(surface *view.Surface) {
	self.child.Draw(surface)
}
