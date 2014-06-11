// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	// "fmt"
	"view/tokenizer"
)

func (self *Editor) MoveCursor(x, y float64) {
	findChar := func(x, y float64) bool {
		// position the first cursor
		cur := &self.Cursors[0]
		for l := 0; l < len(self.Lines); l++ {
			// get the last char for sampling
			linelen := len(self.Lines[l])
			last := self.Lines[l][linelen-1]

			if y >= last.Bounds.Y && y <= last.Bounds.Y+last.Bounds.Height {
				if x >= last.Bounds.X+last.Bounds.Width {
					cur.Column = linelen - 1
					cur.Line = l
					return true
				}

				for c := 0; c < linelen; c++ {
					char := self.Lines[l][c]
					if char.Bounds.Contains(x, y) {
						cur.Column = c
						cur.Line = l
						return true
					}
				}
			}
		}
		return false
	}

	// LIMITATION
	// Clicking in between lines of text will only seek
	// 1000px upwards or downwards for a character to
	// place the cursor at.
	for i := 0; !findChar(x, y) && i < 200; i++ {
		if y <= self.Style().PaddingTop() {
			y += 5
		} else {
			y -= 5
		}
	}

	return
}

func (self *Editor) MoveCursorsToPreviousToken() {
	for i := 0; i < len(self.Cursors); i++ {
		c := &self.Cursors[i]
		token := self.Lines[c.Line][c.Column].Token
		pos := 0
		for j := c.Column; j > 0; j-- {
			char := self.Lines[c.Line][j]
			if !char.Token.Type.Whitespace() && char.Token.Start != token.Start {
				pos = c.Column - (token.Start - char.Token.Start)
				break
			}
		}
		c.Column = pos
	}
}

func (self *Editor) MoveCursorsToNextToken() {
	for i := 0; i < len(self.Cursors); i++ {
		c := &self.Cursors[i]
		token := self.Lines[c.Line][c.Column].Token
		pos := len(self.Lines[c.Line]) - 1
		for j := c.Column; j < len(self.Lines[c.Line]); j++ {
			char := self.Lines[c.Line][j]
			if !char.Token.Type.Whitespace() && char.Token.Start != token.Start {
				pos = char.Index - self.Lines[c.Line][0].Index
				break
			}
		}
		c.Column = pos
	}
}

// MoveCursorToLineStart moves the cursors to the begining of the line.
func (self *Editor) MoveCursorToLineStart() {
	for i := 0; i < len(self.Cursors); i++ {
		c := &self.Cursors[i]
		if c.Column > 0 {
			c.Column = 0
		}
	}
}

// MoveCursorToLineEnd moves the cursors to the end of the line.
func (self *Editor) MoveCursorToLineEnd() {
	for i := 0; i < len(self.Cursors); i++ {
		c := &self.Cursors[i]
		c.Column = len(self.Lines[c.Line]) - 1
	}
}

// MoveCursorLeft moves the cursor left one space or two the end of
// the previous line if at the bigging of the line.
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

// MoveCursorsRight moves the cursor right one space or to the next
// line if the cursor is at the end of the line.
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

// MoveCursorsUp moves the cursor up to the nearest column position on
// the previous line.
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

// MoveCursorsDown moves the cursor down to the nearest column
// position on the next line.
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

// DeleteCharBeforeCursors will remove the characters preceding the
// cursors.  Like a standard backspace operation.
func (self *Editor) DeleteCharBeforeCursors() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Column > 0 {
			pos := self.Lines[c.Line][c.Column].Index
			self.Text = self.Text[:pos-1] + self.Text[pos:]
			self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
			c.Column--

		} else {
			if c.Line == 0 {
				return
			} else {
				col := len(self.Lines[c.Line-1]) - 1
				pos := self.Lines[c.Line-1][col].Index
				self.Text = self.Text[:pos] + self.Text[pos+1:]
				self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
				c.Line--
				c.Column = col
			}
		}
		self.Cursors[i] = c
	}
}

// DeleteCharBeforeCursors will remove the characters following the
// cursors.  Like a standard delete operation.
func (self *Editor) DeleteCharAfterCursors() {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		if c.Line == len(self.Lines)-1 && c.Column == len(self.Lines[c.Line])-1 {
			return
		}
		pos := self.Lines[c.Line][c.Column].Index
		self.Text = self.Text[:pos] + self.Text[pos+1:]
		self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
		self.Cursors[i] = c
	}
}

// InsertCharAtCursors will place a character at the cursor location
// moving subsequent character right
func (self *Editor) InsertCharAtCursors(r rune) {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		pos := self.Lines[c.Line][c.Column].Index
		self.Text = self.Text[:pos] + string(r) + self.Text[pos:]
		self.Lines = tokenizer.ToLinesOfCharacters(self.Tokenizer.Tokenize(self.Text))
		if r == '\n' {
			c.Column = 0
			c.Line++
		} else {
			c.Column++
		}
		self.Cursors[i] = c
	}
}
