package view

import ()



type Color struct {
	R, G, B, A float64
}

func clritof (c uint8) float64 {
	if c == 0 {
		return 0
	} else {
		return float64(c) / 255
	}
}

func UintColor(r, g, b, a uint8) Color {
	return Color{clritof(r), clritof(g), clritof(b), clritof(a)}
}

func HexColor(hex uint32) Color {
	return Color{clritof(uint8(hex >> 24 & 0xFF)), 
	             clritof(uint8(hex >> 16 & 0xFF)), 
	             clritof(uint8(hex >> 8 & 0xFF)), 
	             clritof(uint8(hex & 0xFF))}
}

type Bounds struct {
	X, Y, Width, Height uint
}