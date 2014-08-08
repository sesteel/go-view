package text

import (
	"view"
	"view/tokenizer"
)

type line struct {
	parent styler
	tokens []*tokenizer.Token

	// nil when dirty
	surface *view.Surface
}

func (self *line) Draw(s *view.Surface) {
	if self.surface == nil {

	}
}
