// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"log"

	"github.com/sesteel/go-view/tokenizer"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func (self *Editor) ScrollTo(offset float64) {
	if offset < 0 || int(offset) > len(self.Lines) {
		return
	}
	self.vscroll.SetOffset(offset)

}

func (self *Editor) AddSelection(sel *Selection) {
	self.Selections = append(self.Selections, sel)
}

func (self *Editor) ClearSelections() {
	self.Selection = &Selection{Range{Index{-1, -1}, Index{-1, -1}}}
	self.Selections = make([]*Selection, 0)
}

func (self *Editor) Select(start, end Index) {
	if self.Selection.Start.Line < 0 {
		self.Selections = append(self.Selections, self.Selection)
	}
	self.Selection.Start = start
	self.Selection.End = end
	self.Selection.Normalize()
}

func (self *Editor) MoveCursor(x, y float64) {
	idx := self.FindClosestIndex(x, y)

	if idx.Line > -1 {
		self.Cursors[0].Line = idx.Line
		self.Cursors[0].Column = idx.Column
	}
}

func (self *Editor) FindClosestIndex(x, y float64) Index {
	idx := Index{-1, -1}

	findChar := func(x, y float64) bool {

		yoff := 0.0

		for l := int(self.vscroll.Offset()); l < len(self.lineSurfaces); l++ {
			// get the last char for sampling
			linelen := len(self.Lines[l].Characters)
			surface := self.lineSurfaces[l]
			off := yoff + float64(surface.Height())

			if off < y {
				yoff = off
				continue
			}

			// last := self.Lines[l].Characters[linelen-1]
			lastBounds := self.Lines[l].Bounds[linelen-1]
			log.Println("::::", x, y, lastBounds, off, yoff)
			// if lastBounds.X != -1 {
			// lx := lastBounds.X

			if x >= lastBounds.X+lastBounds.Width {
				idx.Column = linelen - 1
				idx.Line = l
				return true
			}

			for c := 0; c < linelen; c++ {
				char := self.Lines[l].Bounds[c]
				log.Println(char, y, y-yoff, yoff)
				if char.Contains(x, y-yoff) {
					idx.Column = c
					idx.Line = l
					return true
				}
			}
			// }
		}
		return false
	}

	// LIMITATION
	// Clicking in between lines of text will only seek
	// 1000px upwards or downwards for a character to
	// place the cursor at.
	findChar(x, y)
	// for i := 0; !findChar(x, y) && i < 200; i++ {
	// 	if y <= self.Style.Paddings.Top {
	// 		y += 5
	// 	} else {
	// 		y -= 5
	// 	}
	// }

	return idx
}

func (self *Editor) MoveCursorsToPreviousToken() {
	self.AtEachCursor(func(c *Cursor) {
		token := self.Lines[c.Line].Characters[c.Column].Token
		pos := 0
		for j := c.Column; j > 0; j-- {
			char := self.Lines[c.Line].Characters[j]
			if !char.Token.Type.Whitespace() && char.Token.Start != token.Start {
				pos = c.Column - (token.Start - char.Token.Start)
				break
			}
		}
		c.Column = pos
	})
}

func (self *Editor) MoveCursorsToNextToken() {
	self.AtEachCursor(func(c *Cursor) {
		token := self.Lines[c.Line].Characters[c.Column].Token
		pos := len(self.Lines[c.Line].Characters) - 1
		for j := c.Column; j < len(self.Lines[c.Line].Characters); j++ {
			char := self.Lines[c.Line].Characters[j]
			if !char.Token.Type.Whitespace() && char.Token.Start != token.Start {
				pos = char.Index - self.Lines[c.Line].Characters[0].Index
				break
			}
		}
		c.Column = pos
	})
}

// MoveCursorToLineStart moves the cursors to the begining of the line.
func (self *Editor) MoveCursorToLineStart() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Column > 0 {
			c.Column = 0
		}
	})
}

// MoveCursorToLineEnd moves the cursors to the end of the line.
func (self *Editor) MoveCursorToLineEnd() {
	self.AtEachCursor(func(c *Cursor) {
		c.Column = len(self.Lines[c.Line].Characters) - 1
	})
}

// MoveCursorLeft moves the cursor left one space or two the end of
// the previous line if at the bigging of the line.
func (self *Editor) MoveCursorsLeft() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Column > 0 {
			c.Column--
		} else {
			if c.Line == 0 {
				return
			} else {
				c.Line--
				c.Column = len(self.Lines[c.Line].Characters) - 1
			}
		}
	})
}

// MoveCursorsRight moves the cursor right one space or to the next
// line if the cursor is at the end of the line.
func (self *Editor) MoveCursorsRight() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Column < len(self.Lines[c.Line].Characters)-1 {
			c.Column++
		} else {
			if c.Line == len(self.Lines)-1 {
				return
			} else {
				c.Line++
				c.Column = 0
			}
		}
	})
}

// MoveCursorsUp moves the cursor up to the nearest column position on
// the previous line.
func (self *Editor) MoveCursorsUp() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Line > 0 {
			c.Line--
			l := len(self.Lines[c.Line].Characters)
			if c.Column > l-1 {
				c.Column = l - 1
			}
		}
	})
}

// MoveCursorsDown moves the cursor down to the nearest column
// position on the next line.
func (self *Editor) MoveCursorsDown() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Line < len(self.Lines)-1 {
			c.Line++
			l := len(self.Lines[c.Line].Characters)
			if c.Column > l-1 {
				c.Column = l - 1
			}
		}
	})
}

// DeleteCharBeforeCursors will remove the characters preceding the
// cursors.  Like a standard backspace operation.
func (self *Editor) DeleteCharBeforeCursors() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Column > 0 {
			pos := self.Lines[c.Line].Characters[c.Column].Index
			self.text = self.text[:pos-1] + self.text[pos:]
			self.Lines = tokenizer.ToLines(self.Tokenizer.Tokenize(self.text))
			c.Column--
			self.destroyLineSurface(c.Line)
		} else {
			if c.Line == 0 {
				return
			} else {
				col := len(self.Lines[c.Line-1].Characters) - 1
				pos := self.Lines[c.Line-1].Characters[col].Index
				self.SetText(self.text[:pos] + self.text[pos+1:])
				self.removeLineSurface(c.Line)
				c.Line--
				self.destroyLineSurface(c.Line)
				c.Column = col
			}
		}
	})
}

// DeleteCharBeforeCursors will remove the characters following the
// cursors.  Like a standard delete operation.
func (self *Editor) DeleteCharAfterCursors() {
	self.AtEachCursor(func(c *Cursor) {
		if c.Line == len(self.Lines)-1 && c.Column == len(self.Lines[c.Line].Characters)-1 {
			return
		}
		pos := self.Lines[c.Line].Characters[c.Column].Index
		self.text = self.text[:pos] + self.text[pos+1:]
		self.Lines = tokenizer.ToLines(self.Tokenizer.Tokenize(self.text))
		self.destroyLineSurface(c.Line)
	})
}

// InsertCharAtCursors will place a character at the cursor location
// moving subsequent character right
func (self *Editor) InsertCharAtCursors(r rune) {
	self.AtEachCursor(func(c *Cursor) {
		// log.Println(string(r))
		pos := self.Lines[c.Line].Characters[c.Column].Index
		self.text = self.text[:pos] + string(r) + self.text[pos:]
		self.Lines = tokenizer.ToLines(self.Tokenizer.Tokenize(self.text))
		if r == '\n' {
			c.Column = 0
			c.Line++
		} else {
			c.Column++
		}
		self.destroyLineSurface(c.Line)
	})
}

func (self *Editor) AtEachCursor(f func(*Cursor)) {
	for i := 0; i < len(self.Cursors); i++ {
		c := self.Cursors[i]
		f(c)
	}
}
