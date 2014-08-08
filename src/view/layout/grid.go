package layout

import (
	// "log"
	"view"
	"view/common"
	"view/event"
)

type gridViewContainer struct {
	view           view.View
	row, col, w, h int
}

const (
	DefaultGridCellWidth  = 100
	DefaultGridCellHeight = 40
)

// Grid is a simple layout that organizes
// child views by way of a grid system.
type Grid struct {
	target        view.View
	cellSpacing   float64
	children      []*gridViewContainer
	cellHeights   []float64
	cellWidths    []float64
	allowOverflow bool
}

func NewGrid(target view.View) *Grid {
	g := &Grid{
		target,
		3,
		make([]*gridViewContainer, 0),
		make([]float64, 0),
		make([]float64, 0),
		false,
	}
	return g
}

func (self *Grid) Add(child view.View, row, col, w, h int) error {
	gvc := &gridViewContainer{child, row, col, w, h}

	self.children = append(self.children, gvc)

	for i := len(self.children); i < row; i++ {
		self.children = append(self.children, gvc)
	}

	for i := len(self.cellHeights); row+h-1 >= len(self.cellHeights); i++ {
		self.cellHeights = append(self.cellHeights, DefaultGridCellHeight)
	}

	for i := len(self.cellWidths); col+w-1 >= len(self.cellWidths); i++ {
		self.cellWidths = append(self.cellWidths, DefaultGridCellWidth)
	}

	return nil
}

func (self *Grid) Draw(surface *view.Surface) {
	for _, child := range self.children {
		b := self.bounds(child)
		s := view.NewSurface(view.FORMAT_ARGB32, int(b.Width), int(b.Height))
		child.view.Draw(s)
		surface.SetSourceSurface(s, b.X, b.Y)
		surface.Paint()
		s.Destroy()
	}
}

func (self *Grid) bounds(child *gridViewContainer) common.Bounds {
	x, y := 0.0, 0.0
	w, h := 0.0, 0.0

	for i := 0; i < child.col+child.w && i < len(self.cellWidths); i++ {
		if i < child.col {
			x += self.cellSpacing
			x += self.cellWidths[i]
		} else {
			w += self.cellSpacing
			w += self.cellWidths[i]
		}
	}

	for i := 0; i < child.row+child.h && i < len(self.cellHeights); i++ {
		if i < child.row {
			y += self.cellSpacing
			y += self.cellHeights[i]
		} else {
			h += self.cellSpacing
			h += self.cellHeights[i]
		}
	}
	return common.Bounds{common.Point{x, y}, common.Size{w, h}}
}

func (self *Grid) Animate(surface *view.Surface) {
	for _, child := range self.children {
		if anim, ok := child.view.(view.Animator); ok {
			b := self.bounds(child)
			s := view.NewSurface(view.FORMAT_ARGB32, int(b.Width), int(b.Height))
			anim.Animate(s)
			surface.SetSourceSurface(s, b.X, b.Y)
			surface.Paint()
			s.Destroy()
		}
	}
}

func (self *Grid) Redraw() {
	self.target.Redraw()
}

func (self *Grid) mapMouseEventToBounds(ev event.Mouse, f func(*gridViewContainer, event.Mouse)) {
	for _, child := range self.children {
		b := self.bounds(child)
		if b.Contains(ev.X, ev.Y) {
			e := ev.Normalize(b.Point)
			f(child, e)
		}
	}
}

func (self *Grid) MousePosition(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MousePosition(ev)
	})
}

func (self *Grid) MouseButtonPress(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MouseButtonPress(ev)
	})
}

func (self *Grid) MouseButtonRelease(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MouseButtonRelease(ev)
	})
}

func (self *Grid) MouseEnter(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MouseEnter(ev)
	})
}

func (self *Grid) MouseExit(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MouseExit(ev)
	})
}

func (self *Grid) MouseWheelUp(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MouseWheelUp(ev)
	})
}

func (self *Grid) MouseWheelDown(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *gridViewContainer, ev event.Mouse) {
		child.view.MouseWheelDown(ev)
	})
}
