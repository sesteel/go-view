package editor

import (
	"view/color"
)

// TokenStyle provides the styling for various token types or subtypes.
type TokenStyle struct {
	Weight int
	Slant  int
	Color  color.RGBA
}
