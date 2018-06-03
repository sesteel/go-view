package editor

import (
	"github.com/sesteel/go-view/color"
)

// TokenStyle provides the styling for various token types or subtypes.
type TokenStyle struct {
	Weight int
	Slant  int
	Color  color.RGBA
}
