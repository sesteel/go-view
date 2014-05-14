// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"view/tokenizer"
	//	"regexp"
)

// Model represents a renderable set of text
// with the UI.  It provides a means to modify,
// search, and tokenize a set of text.
type Model interface {
	//	Find(*regexp.Regexp)
	//	FindAll(*regexp.Regexp)
	//	Insert(string, uint)
	//	Overwrite(string, uint)
	//	CursorPos() uint
	Lines() Lines
	Frame(uint, uint) Lines
}

type Lines [][]*tokenizer.Token

type defaultModel struct {
	name      string
	text      string
	lines     Lines
	cursorPos uint
}

func NewModel(name, text string) Model {
	d := new(defaultModel)
	d.cursorPos = 0
	d.name = name
	d.text = text
	d.lines = tokenizer.ToLines(tokenizer.Tokenize(text))
	return d
}

// Frame returns the lines from the model from the
// start index to the end index.
func (self *defaultModel) Frame(start, end uint) Lines {
	return self.lines[start:end]
}

// Lines returns the lines of the model.
func (self *defaultModel) Lines() Lines {
	return self.lines
}
