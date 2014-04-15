package layout

import (
	"view"
	"view/event"
)

type absGroup struct {
	view view.View
	bounds view.Bounds
	mouseIn bool
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

func (self *Absolute) Add(d view.View, bounds view.Bounds) {
	self.children = append(self.children, &absGroup{d, bounds, false})
}

func (self *Absolute) Draw(s *view.Surface) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		ns := view.NewSurface(view.FORMAT_ARGB32, int(g.bounds.Width), int(g.bounds.Height))
		defer ns.Destroy()
		g.view.Draw(ns)
		s.SetSourceSurface(ns, g.bounds.X, g.bounds.Y)
		s.Paint()
		
	}
}

func (self *Absolute) Redraw() {
	self.target.Redraw()
}

func (self *Absolute) MouseEnter(ev event.Mouse) {}

func (self *Absolute) MouseExit(ev event.Mouse) {}

func (self *Absolute) MousePosition(ev event.Mouse) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		if g.bounds.Contains(float64(ev.X), float64(ev.Y)) {
			g.view.MousePosition(ev)
			if !g.mouseIn {
				g.mouseIn = true
				g.view.MouseEnter(ev)
			} 
		} else if g.mouseIn {
			g.mouseIn = false
			g.view.MouseExit(ev)
		}
	}
}

func (self *Absolute) MouseButtonPress(ev event.Mouse) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		if g.bounds.Contains(float64(ev.X), float64(ev.Y)) {
			g.view.MouseButtonPress(ev)
		}
	}
}

func (self *Absolute) MouseButtonRelease(ev event.Mouse) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		if g.bounds.Contains(float64(ev.X), float64(ev.Y)) {
			g.view.MouseButtonRelease(ev)
		}
	}
}

func (self *Absolute) MouseWheelUp(ev event.Mouse) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		if g.bounds.Contains(float64(ev.X), float64(ev.Y)) {
			g.view.MouseWheelUp(ev)
		}
	}
}

func (self *Absolute) MouseWheelDown(ev event.Mouse) {
	for i := 0; i < len(self.children); i++ {
		g := self.children[i]
		if g.bounds.Contains(float64(ev.X), float64(ev.Y)) {
			g.view.MouseWheelDown(ev)
		}
	}
}