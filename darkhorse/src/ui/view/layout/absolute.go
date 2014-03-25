package layout

import (
	"ui/view"
)

type absGroup struct {
	drawer view.Drawer
	x, y   float64
}

type Absolute struct {
	target   view.Composite
	children []*absGroup
}

func NewAbsolute(target view.Composite) *Absolute {
	l := new(Absolute)
	l.target = target
	return l
}

func (self *Absolute) Add(d view.Drawer, x, y float64) {
	self.children = append(self.children, &absGroup{d, x, y})
}

func (self *Absolute) Draw(s *view.Surface) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		g.drawer.Draw(s)
	}
}
