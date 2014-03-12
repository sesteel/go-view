package layout

import (
	"ui/view"
)

type xypGroup struct {
	drawer view.Drawer
	x, y   float64
}

type XYPositioned struct {
	target   view.Composite
	children []*xypGroup
}

func NewXYPositioned(target view.Composite) *XYPositioned {
	l := new(XYPositioned)
	l.target = target
	return l
}

func (self *XYPositioned) Add(d view.Drawer, x, y float64) {
	self.children = append(self.children, &xypGroup{d, x, y})
}

func (self *XYPositioned) Draw() {
	s := self.target.Surface()
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		//s.PushGroup()
		s.Translate(g.x, g.y)
		g.drawer.Draw()
		s.Translate(-g.x, -g.y)
		//s.PopGroup()
	}
}
