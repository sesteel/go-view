// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package text

import ()

// MoveCursorLeft moves the cursor left one space or two the end of
// the previous line if at the bigging of the line.
func (self *Editor) MoveCursorsLeft() {
	// self.AtEachCursor(func(c *Cursor) {
	c := self.cursor
	if c.Column > 0 {
		c.Column--
	} else {
		if c.Line == 0 {
			return
		} else {
			c.Line--
			c.Column = len(self.lines[c.Line]) - 1
		}
	}
	c.markTime()
	// })
}

// MoveCursorsRight moves the cursor right one space or to the next
// line if the cursor is at the end of the line.
func (self *Editor) MoveCursorsRight() {
	// self.AtEachCursor(func(c *Cursor) {
	c := self.cursor
	// fmt.Println(c)
	if c.Column < len(self.lines[c.Line])-1 {
		c.Column++
	} else {
		if c.Line == len(self.lines)-1 {
			return
		} else {
			c.Line++
			c.Column = 0
		}
	}
	c.markTime()
	// })
}

// MoveCursorsUp moves the cursor up to the nearest column position on
// the previous line.
func (self *Editor) MoveCursorsUp() {
	// self.AtEachCursor(func(c *Cursor) {
	c := self.cursor
	if c.Line > 0 {
		c.Line--
		l := len(self.lines[c.Line])
		if c.Index.Column > l-1 {
			c.Index.Column = l - 1
		}
	}
	c.markTime()
	// })
}

// MoveCursorsDown moves the cursor down to the nearest column
// position on the next line.
func (self *Editor) MoveCursorsDown() {
	// self.AtEachCursor(func(c *Cursor) {
	c := self.cursor
	if c.Line < len(self.lines)-1 {
		c.Line++
		l := len(self.lines[c.Line])
		if c.Column > l-1 {
			c.Column = l - 1
		}
	}
	c.markTime()
	// })
}

// InsertCharAtCursors will place a character at the cursor location
// moving subsequent character right
func (self *Editor) InsertCharAtCursors(r rune) {
	c := self.cursor
	pos := self.lines[c.Line][c.Column].Index
	self.text = self.text[:pos] + string(r) + self.text[pos:]
	self.lines = self.toLines(self.tokenizer.Tokenize(self.text))
	if r == '\n' {
		c.Column = 0
		c.Line++
	} else {
		c.Column++
	}
	c.markTime()
	self.Redraw()
}

// DeleteCharBeforeCursors will remove the characters preceding the
// cursors.  Like a standard backspace operation.
func (self *Editor) DeleteCharBeforeCursors() {
	c := self.cursor
	if c.Column > 0 {
		pos := self.lines[c.Line][c.Column].Index
		self.text = self.text[:pos-1] + self.text[pos:]
		self.lines = self.toLines(self.tokenizer.Tokenize(self.text))
		c.Column--
	} else {
		if c.Line == 0 {
			return
		} else {
			col := len(self.lines[c.Line-1]) - 1
			pos := self.lines[c.Line-1][col].Index
			self.SetText(self.text[:pos] + self.text[pos+1:])
			c.Line--
			c.Column = col
		}
	}
	c.markTime()
	self.Redraw()
}

// DeleteCharBeforeCursors will remove the characters following the
// cursors.  Like a standard delete operation.
func (self *Editor) DeleteCharAfterCursors() {
	c := self.cursor
	if c.Line == len(self.lines)-1 && c.Column == len(self.lines[c.Line])-1 {
		return
	}
	pos := self.lines[c.Line][c.Column].Index
	self.text = self.text[:pos] + self.text[pos+1:]
	self.lines = self.toLines(self.tokenizer.Tokenize(self.text))
	c.markTime()
	self.Redraw()
}
