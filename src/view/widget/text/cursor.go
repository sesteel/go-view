// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package text

import (
	// "log"
	"time"
	"view"
	"view/color"
	"view/tokenizer"
)

type CursorType int

const (
	DEFAULT CursorType = iota
	BAR
	BLINK
	// BLOCK
	OUTLINE
	UNDERLINE
)

// Cursor is used to store the position of the cursor via a Index.
type Cursor struct {
	Index
	Type         CursorType
	Color        *color.RGBA
	LineWidth    float64
	LastMovement time.Time
}

func (self *Cursor) PreviousPos(lines []tokenizer.Line) Index {
	if len(lines) < self.Line-1 || self.Line < 0 {
		return Index(self.Index)
	}

	if self.Column == 0 {
		return Index{self.Line - 1, len(lines[self.Line-1].Characters)}
	} else {
		return Index{self.Line, self.Column - 1}
	}
}

func (self *Cursor) Draw(s *view.Surface, e *Editor) {
	offset := int(e.offset)
	now := time.Now()
	char := e.lines[self.Index.Line][self.Index.Column]
	extents := e.style(char.Token.Type).extents['M']
	b := char.Bounds
	b.Height = extents.Height
	b.Width = extents.Width

	yoff := 0.0
	y := b.Y + yoff
	// log.Println(b)
	// the cursor is not visible
	if offset > self.Line {
		return
	}

	recentlyMoved := time.Since(self.LastMovement) < time.Second/2

	blink := func(max float64) {
		if recentlyMoved {
			self.Color.A = 1
		} else if now.Nanosecond() < 450000000 {
			self.Color.A = 0
		} else {
			self.Color.A = max
		}
	}

	prepare := func() {
		s.SetSourceRGBA(*self.Color)
		s.SetLineWidth(self.LineWidth)
	}

	// fade
	switch self.Type {

	case BAR:
		self.Color.A = 1
		prepare()
		s.MoveTo(b.X, y+ALIGN-b.Height/.75)
		s.LineTo(b.X, y+ALIGN+b.Height/2)
		s.Stroke()

	case BLINK:
		blink(1)
		prepare()
		s.MoveTo(b.X, y+ALIGN-b.Height/.75)
		s.LineTo(b.X, y+ALIGN+b.Height/2)
		s.Stroke()

	// case BLOCK:
	// 	blink(0.5)
	// 	prepare()
	// 	s.RoundedRectangle(b.X+ALIGN, y-b.Height+ALIGN, b.Width, b.Height+b.Height/2, 1, 1, 1, 1)
	// 	s.Fill()
	// 	s.RoundedRectangle(b.X+ALIGN, y-b.Height+ALIGN, b.Width, b.Height+b.Height/2, 1, 1, 1, 1)
	// 	s.Stroke()

	case OUTLINE:
		blink(0.7)
		prepare()
		s.RoundedRectangle(b.X+ALIGN, y-ALIGN-b.Height*1.25, b.Width, b.Height/.50, 1, 1, 1, 1)
		s.Stroke()

	case UNDERLINE:
		blink(1)
		prepare()
		s.MoveTo(b.X, y+ALIGN+b.Height/2)
		s.LineTo(b.X+b.Width, y+ALIGN+b.Height/2)
		s.Stroke()

	case DEFAULT:
		fallthrough
	default:
		if recentlyMoved {
			self.Color.A = 1
		} else if now.Nanosecond() < 300000000 {
			self.Color.A += 0.075
		} else if now.Nanosecond() < 600000000 {
			self.Color.A = 1
		} else if self.Color.A > 0 {
			self.Color.A -= 0.075
		}
		prepare()
		s.MoveTo(b.X+1, y-b.Height/.75)
		s.LineTo(b.X+1, y-ALIGN+b.Height/2)
		s.Stroke()
	}
	s.Flush()
}

func (self *Cursor) markTime() {
	self.LastMovement = time.Now()
}
