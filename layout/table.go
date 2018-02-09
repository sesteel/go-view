package layout

import (
	// "log"
	"view"
	"view/common"
	"view/event"
)

type cell struct {
	view           view.View
	row, col, w, h int
}

const (
	cellWidth  = 100
	cellHeight = 20
)

// Table is a simple layout that organizes child views by way of a
// classic cell system.  The Table differs grom the Grid in several
// ways; most notably Table overflows the view space while a Grid
// does not.
type Table struct {
	target      view.View
	cellSpacing float64
	children    []*cell
	cellHeights []float64
	cellWidths  []float64
}

func NewTable(target view.View) *Table {
	g := &Table{
		target,
		3,
		make([]*cell, 0),
		make([]float64, 0),
		make([]float64, 0),
	}
	return g
}

func (self *Table) SetRowHeight(row int, h float64) {
	for row >= len(self.cellHeights) {
		self.cellHeights = append(self.cellHeights, cellHeight)
	}
	self.cellHeights[row] = h
}

func (self *Table) SetColWidth(col int, w float64) {
	for col >= len(self.cellWidths) {
		self.cellWidths = append(self.cellWidths, cellWidth)
	}
	self.cellWidths[col] = w
}

func (self *Table) AddCellView(child view.View, row, col int) error {
	return self.AddMultiCellView(child, row, col, 1, 1)
}

func (self *Table) AddMultiCellView(child view.View, row, col, w, h int) error {
	gvc := &cell{child, row, col, w, h}

	self.children = append(self.children, gvc)

	for i := len(self.children); i < row; i++ {
		self.children = append(self.children, gvc)
	}

	for row+h-1 >= len(self.cellHeights) {
		self.cellHeights = append(self.cellHeights, cellHeight)
	}

	for col+w-1 >= len(self.cellWidths) {
		self.cellWidths = append(self.cellWidths, cellWidth)
	}

	return nil
}

func (self *Table) Draw(surface *view.Surface) {
	for _, child := range self.children {
		b := self.bounds(child)
		s := view.NewSurface(view.FORMAT_ARGB32, int(b.Width), int(b.Height))
		child.view.Draw(s)
		surface.SetSourceSurface(s, b.X, b.Y)
		surface.Paint()
		s.Destroy()
	}
}

func (self *Table) bounds(child *cell) common.Bounds {
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

func (self *Table) Animate(surface *view.Surface) {
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

func (self *Table) Redraw() {
	self.target.Redraw()
}

func (self *Table) mapMouseEventToBounds(ev event.Mouse, f func(*cell, event.Mouse)) {
	for _, child := range self.children {
		b := self.bounds(child)
		if b.Contains(ev.X, ev.Y) {
			e := ev.Normalize(b.Point)
			f(child, e)
		}
	}
}

func (self *Table) MousePosition(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MousePosition(ev)
	})
}

func (self *Table) MouseButtonPress(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MouseButtonPress(ev)
	})
}

func (self *Table) MouseButtonRelease(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MouseButtonRelease(ev)
	})
}

func (self *Table) MouseEnter(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MouseEnter(ev)
	})
}

func (self *Table) MouseExit(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MouseExit(ev)
	})
}

func (self *Table) MouseWheelUp(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MouseWheelUp(ev)
	})
}

func (self *Table) MouseWheelDown(ev event.Mouse) {
	self.mapMouseEventToBounds(ev, func(child *cell, ev event.Mouse) {
		child.view.MouseWheelDown(ev)
	})
}
