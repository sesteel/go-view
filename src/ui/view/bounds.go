package view

import ()

type ScrollOffset float64 

type Size struct {
	Width, Height float64
}

type Bounds struct {
	X, Y float64
	Size
}

func (b Bounds) Contains(x, y float64) bool {
	return x >= b.X && 
	       x <= (b.X + b.Width) && 
	       y >= b.Y && 
	       y <= (b.Y + b.Height)
}

