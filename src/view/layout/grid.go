package layout

import (
	"log"
	"view"
	"view/event"
	// "view/common"
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

	for i := len(self.cellHeights); row+w-1 >= len(self.cellHeights); i++ {
		self.cellHeights = append(self.cellHeights, DefaultGridCellHeight)
	}

	for i := len(self.cellWidths); col+h-1 >= len(self.cellWidths); i++ {
		self.cellWidths = append(self.cellWidths, DefaultGridCellWidth)
	}

	return nil
}

func (self *Grid) Draw(surface *view.Surface) {
	for j, child := range self.children {

		x, y := 0.0, 0.0
		w, h := 0.0, 0.0

		for i := 0; i < child.col+child.h && i < len(self.cellWidths); i++ {
			if i < child.col {
				x += self.cellSpacing
				x += self.cellWidths[i]
			} else {
				w += self.cellSpacing
				w += self.cellWidths[i]
			}
		}

		for i := 0; i < child.row+child.w && i < len(self.cellHeights); i++ {
			if i < child.row {
				y += self.cellSpacing
				y += self.cellHeights[i]
			} else {
				h += self.cellSpacing
				h += self.cellHeights[i]
			}
		}

		s := view.NewSurface(view.FORMAT_ARGB32, int(w), int(h))
		log.Println("::", j, w, h, x, y)
		child.view.Draw(s)
		surface.SetSourceSurface(s, x, y)
		surface.Paint()
		s.Destroy()
	}
}

func (self *Grid) Animate(surface *view.Surface) {
	for _, child := range self.children {
		if anim, ok := child.view.(view.Animator); ok {
			anim.Animate(surface)
		}
	}
}

func (self *Grid) Redraw() {
	self.target.Redraw()
}

func (self *Grid) MousePosition(ev event.Mouse) {
	// self.child.MousePosition(ev)
}

func (self *Grid) MouseButtonPress(ev event.Mouse) {
	// self.child.MouseButtonPress(ev)
}

func (self *Grid) MouseButtonRelease(ev event.Mouse) {
	// self.child.MouseButtonRelease(ev)
}

func (self *Grid) MouseEnter(ev event.Mouse) {
	// self.child.MouseEnter(ev)
}

func (self *Grid) MouseExit(ev event.Mouse) {
	// self.child.MouseExit(ev)
}

func (self *Grid) MouseWheelUp(ev event.Mouse) {
	// self.child.MouseWheelUp(ev)
}

func (self *Grid) MouseWheelDown(ev event.Mouse) {
	// self.child.MouseWheelDown(ev)
}
