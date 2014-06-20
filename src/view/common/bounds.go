package common

import ()

type Bounds struct {
	Point
	Size
}

func (b Bounds) Contains(x, y float64) bool {
	return x >= b.X && x <= (b.X+b.Width) && y >= b.Y && y <= (b.Y+b.Height)
}
