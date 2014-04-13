package doc

import (
	"view/tokenizer"
//	"regexp"
)

// Document represents a renderable set of text
// with the UI.  It provides a means to modify,
// search, and tokenize a set of text.
type Document interface {
//	Find(*regexp.Regexp)
//	FindAll(*regexp.Regexp)
//	Insert(string, uint)
//	Overwrite(string, uint)
//	CursorPos() uint
	Lines() Lines
	Frame(uint, uint) Lines
}

type Lines [][]*tokenizer.Token

type defaultDocument struct {
	name   string
	text   string
	lines  Lines
	cursorPos uint
}

func NewDocument(name, text string) Document {
	d := new(defaultDocument)
	d.cursorPos = 0
	d.name = name
	d.text = text
	d.lines = tokenizer.ToLines(tokenizer.Tokenize(text))
	return d
}

// Frame returns the lines from the document from the 
// start index to the end index.
func (self *defaultDocument) Frame(start, end uint) Lines {
	return self.lines[start:end]
} 

// Lines returns the lines of the document.
func (self *defaultDocument) Lines() Lines {
	return self.lines
}