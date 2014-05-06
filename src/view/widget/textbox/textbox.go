// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package textbox

import (
	//	"fmt"
	"view"
	"view/event"
	"view/color"
	"view/tokenizer"
)

type TextBox struct {
	view.DefaultComponent
	verticalOffset view.ScrollOffset
	tkns           []*tokenizer.Token
	cursor         *tokenizer.Token
	//	model Document
}

func New(parent view.View, text string) *TextBox {
	tb := &TextBox{*view.NewComponent(parent, text), 0, tokenizer.Tokenize(text), new(tokenizer.Token)}
	tb.AddMouseWheelDownHandler(func(event.Mouse) {
		tb.verticalOffset++
		tb.Redraw()
	})

	tb.AddMouseWheelUpHandler(func(event.Mouse) {
		if tb.verticalOffset > 0 {
			tb.verticalOffset--
		}
		tb.Redraw()
	})

	tb.AddMouseWheelUpHandler(func(event.Mouse) {
		if tb.verticalOffset > 0 {
			tb.verticalOffset--
		}
		tb.Redraw()
	})

	tb.AddKeyPressHandler(func(k event.Keyboard) {
		text = k.String() + text
		tb.tkns = tokenizer.Tokenize(text)
		tb.Redraw()
	})
	
	cur := new(tokenizer.Token)
	cur.Type = tokenizer.CURSOR
	return tb
}

//func (self *TextBox) MoveCursor(position int) {
//	self.tkns = append(self.tkns, self.cursor)
//	copy(self.tkns[self.cursor+1:], self.tkns[self.cursor:])
//	self.tkns[tb.cursor] = cur
//}

func (self *TextBox) Draw(s *view.Surface) {

	// resize to draw within outline
	b := view.Bounds{0, 0, view.Size{float64(s.Width()), float64(s.Height())}}
	s.Rectangle(b.X, b.Y, b.Width, b.Height)
	s.SetSourceRGBA(color.White)
	s.Fill()
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	drawnLines, lines, height := s.DrawWrappedPlainText(self.tkns, b, self.verticalOffset, self.Style())

	d := lines - drawnLines
	if d == 0 {
		d = 1
	}
	percent := float64(self.verticalOffset) / d

	if height > float64(s.Height()) {
		s.DrawVerticalOverflow2(lines, drawnLines, percent, self.Style())
	}

	//	if width > float64(s.GetWidth()) {
	//	s.DrawHorizontalOverflow(height, self.Style())
	//	}
}
