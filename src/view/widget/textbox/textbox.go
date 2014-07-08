// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package textbox

import (
	//	"fmt"
	"view"
	"view/color"
	"view/common"
	"view/event"
	"view/tokenizer"
	"view/tokenizer/plaintext"
)

type TextBox struct {
	view.DefaultComponent
	verticalOffset view.ScrollOffset
	tkns           []*tokenizer.Token
	tknr           tokenizer.Tokenizer
	cursor         *tokenizer.Token
	//	model Document
}

func New(parent view.View, text string) *TextBox {
	tknr := plaintext.New()
	tb := &TextBox{*view.NewComponent(parent, text), 0, tknr.Tokenize(text), tknr, new(tokenizer.Token)}
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

	tb.AddKeyPressHandler(func(k event.Keyboard) {
		text = k.String() + text
		tb.tkns = tknr.Tokenize(text)
		tb.Redraw()
	})

	return tb
}

func (self *TextBox) Draw(s *view.Surface) {

	// resize to draw within outline
	b := common.Bounds{common.Point{0, 0}, common.Size{float64(s.Width()), float64(s.Height())}}
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

}
