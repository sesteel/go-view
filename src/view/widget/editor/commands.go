// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"view/tokenizer"
)

func (self *Editor) MoveCursorsLeft() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Column > 0 {
			c.Column--
		} else {
			if c.Line == 0 {
				return
			} else {
				c.Line--
				c.Column = len(self.Lines[c.Line]) - 1
			}
		}
		self.Cursors[i] = c
	}
}

func (self *Editor) MoveCursorsRight() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Column < len(self.Lines[c.Line])-1 {
			c.Column++
		} else {
			if c.Line == len(self.Lines)-1 {
				return
			} else {
				c.Line++
				c.Column = 0
			}
		}
		self.Cursors[i] = c
	}
}

func (self *Editor) MoveCursorsUp() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Line > 0 {
			c.Line--
			l := len(self.Lines[c.Line])
			if c.Column > l-1 {
				c.Column = l - 1
			}
		}
		self.Cursors[i] = c
	}
}

func (self *Editor) MoveCursorsDown() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Line < len(self.Lines)-1 {
			c.Line++
			l := len(self.Lines[c.Line])
			if c.Column > l-1 {
				c.Column = l - 1
			}
		}
		self.Cursors[i] = c
	}
}

func (self *Editor) DeleteCharBeforeCursors() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Column > 0 {
			pos := self.Lines[c.Line][c.Column].Position
			self.Text = self.Text[:pos-1] + self.Text[pos:]
			self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
			c.Column--

		} else {
			if c.Line == 0 {
				return
			} else {
				col := len(self.Lines[c.Line-1]) - 1
				pos := self.Lines[c.Line-1][col].Position
				self.Text = self.Text[:pos] + self.Text[pos+1:]
				self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
				c.Line--
				c.Column = col
			}
		}
		self.Cursors[i] = c
	}
}

func (self *Editor) DeleteCharAfterCursors() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Line == len(self.Lines)-1 && c.Column == len(self.Lines[c.Line]) - 1 {
			return
		}
		pos := self.Lines[c.Line][c.Column].Position
		self.Text = self.Text[:pos] + self.Text[pos+1:]
		self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
		self.Cursors[i] = c
	}
}

func (self *Editor) InsertCharAtCursors(r rune) {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		pos := self.Lines[c.Line][c.Column].Position
		self.Text = self.Text[:pos] + string(r) + self.Text[pos:]
		self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
		if r == '\n'{
			c.Column = 0
			c.Line++
		} else {
			c.Column++
		}
		self.Cursors[i] = c
	} 
}

func (self *Editor) j() {}
