package editor

import (
	"github.com/sesteel/go-view"
)

type extents struct {
	name    string
	surface *view.Surface
	mapping map[rune]*view.TextExtents
}

func (self *extents) Extents(r rune) *view.TextExtents {
	e := self.mapping[r]
	if e == nil {
		e = self.surface.TextExtents(string(r))
		self.mapping[r] = e
	}
	return e
}

var extentMaps map[string]*extents

func init() {
	extentMaps = make(map[string]*extents, 0)
}
