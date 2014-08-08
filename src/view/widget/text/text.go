package text

import (
	"view"
	"view/color"
	"view/tokenizer"
	"view/tokenizer/plaintext"
)

type styler interface {
	style(*tokenizer.Token)
}

type MultiLineLatin struct {
	view.DefaultComponent
	Tokenizer tokenizer.Tokenizer
	lines     []*line
	text      string // TODO Consider how to optimize this away as it gets created repeatedly
}

func NewMultiLineLatin(parent view.View, name string, text string) *MultiLineLatin {
	tknzr := plaintext.New()
	e := &MultiLineLatin{
		*view.NewComponent(parent, name),
		tknzr,
		nil,
		text,
	}
	return e
}

func (self *MultiLineLatin) SetText(text string) {
	lines := tokenizer.ToLinesOfTokens(self.Tokenizer.Tokenize(text))
	self.lines = make([]*line, len(lines))
	for i, l := range lines {
		self.lines[i] = &line{
			self,
			l,
			nil,
		}
	}
	self.text = text
}

func (self *MultiLineLatin) style(t *tokenizer.Token) {

}

func (self *MultiLineLatin) Text() string {
	return self.text
}

func (self *MultiLineLatin) Draw(s *view.Surface) {
	for _, l := range self.lines {
		l.Draw(s)
	}

}
