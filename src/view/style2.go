package view

import (
	"view/color"
)

type BorderColors struct {
	Top, Bottom, Left, Right color.RGBA
}

type BorderWidths struct {
	Top, Bottom, Left, Right float64
}

type CornerRadiuses struct {
	TopLeft, TopRight, BottomLeft, BottomRight float64
}

type Paddings struct {
	Top, Bottom, Left, Right float64
}

type OverflowX int

const (
	OVERFLOW_X_NONE OverflowX = iota
	OVERFLOW_X_SCROLL
	OVERFLOW_X_WRAP
)

type OverflowY int

const (
	OVERFLOW_Y_NONE OverflowY = iota
	OVERFLOW_Y_SCROLL
)

// type TabWidth int
type TextAlignment int

const (
	STYLE_TEXT_LEFT TextAlignment = iota
	STYLE_TEXT_CENTERED
	STYLE_TEXT_RIGHT
	STYLE_TEXT_JUSTIFIED
)

type Font struct {
	Name   string
	Weight int
	Slant  int
	Size   float64
}

// NewFont returns the a Font which contains the
// default font style information.  You may modify
// the returned font to a style suiting your needs.
func NewFont() *Font {
	return &Font{
		"Sans",
		FONT_WEIGHT_NORMAL,
		FONT_SLANT_NORMAL,
		14,
	}
}

func (self *Font) Configure(s *Surface) {
	s.SelectFontFace(self.Name, self.Slant, self.Weight)
	s.SetFontSize(self.Size)
}
