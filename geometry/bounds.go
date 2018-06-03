package geometry

type Bounds struct {
	Point
	Size
}

func NewBounds(x, y, w, h float64) Bounds {
	return Bounds{Point{x, y}, Size{w, h}}
}

func (b Bounds) Contains(x, y float64) bool {
	return x >= b.X && x <= (b.X+b.Width) && y >= b.Y && y <= (b.Y+b.Height)
}
