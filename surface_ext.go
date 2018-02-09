// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package view

import (
	"github.com/sesteel/go-view/color"
	. "github.com/sesteel/go-view/common"
	"github.com/sesteel/go-view/tokenizer"
)

func (self *Surface) DrawTextToken(t *tokenizer.Token, f *Font, b Bounds, c color.RGBA) {
	if t.Selected {
		self.SetSourceRGBA(color.Selection)
		self.Rectangle(b.X, b.Y, b.Width, b.Height)
		self.Fill()
	}

	if f != nil {
		f.Configure(self)
	}
	self.SetSourceRGBA(c)
	self.MoveTo(b.X, b.Y)
	self.ShowText(t.Value)
}

var EXTENTS map[string]*TextExtents = make(map[string]*TextExtents)

// DrawWrappedPlainText uses the Style's foreground color.to draw plain
// ascii formatted text within the bounds.  It stops rendering at last visible
// line, but continues to calculate total height.
// returns height
func (self *Surface) DrawWrappedPlainText(ts []*tokenizer.Token, b Bounds, o ScrollOffset, p Paddings, tabwidth int, f Font, c color.RGBA) (linesDrawn, lines, height float64) {
	b.X += p.Left
	b.Y += p.Top
	b.Width -= (p.Left + p.Right)
	b.Height -= (p.Top + p.Bottom)

	var lineHeight, x, y float64 = 0, 0, 0

	f.Configure(self)
	self.SetSourceRGBA(c)

	spaceExtents := self.TextExtents("M")
	y += spaceExtents.Height
	spaceExtents.Width *= 1
	spaceExtents.Height *= 1.75

	selected := make([]*tokenizer.Token, 0)

	for i := 0; i < len(ts); i++ {
		t := ts[i]
		e := EXTENTS[t.Value]
		if e == nil {
			e = self.TextExtents(t.Value)
			EXTENTS[t.Value] = e
		}

		if t.Value == "\n" {
			x = 0
			if ScrollOffset(lines) >= o {
				y += spaceExtents.Height
			}
			lines++
			continue
		}

		if t.Value == "\t" {
			x += spaceExtents.Width * float64(tabwidth)
			continue
		}

		if e.Height > lineHeight {
			lineHeight = e.Height
		}

		if x+e.Width > b.Width {

			if ScrollOffset(lines) >= o {
				y += spaceExtents.Height
			}
			x = 0
			lines++
		}

		var b2 Bounds
		b2.X = b.X + x
		b2.Y = b.Y + y
		b2.Width = e.Xadvance - b2.X
		b2.Height = e.Yadvance - b2.Y

		if t.Selected {
			selected = append(selected, t)
		}

		if y < b.Height && ScrollOffset(lines) >= o {
			linesDrawn = float64(lines) - float64(o)
			self.DrawTextToken(t, nil, b2, c)
		}

		x += e.Xadvance
		y += e.Yadvance
	}

	return linesDrawn, lines, y + p.Bottom
}

func (self *Surface) DrawVerticalOverflow2(rows, shown, percent float64, oy OverflowY) {
	if rows <= 0 || shown <= 0 {
		rows = 1
		shown = 1
	}

	switch oy {
	case OVERFLOW_Y_SCROLL:
		self.SetSourceRGBA(color.HexRGBA(0x00000007))
		self.Rectangle(float64(self.Width())-10, 0, float64(self.Width()), float64(self.Height()))
		self.Fill()
		ratio := float64(shown) / float64(rows)
		height := ratio * float64(self.Height())
		if height < 5 {
			height = 15
		}
		self.SetSourceRGBA(color.HexRGBA(0x00000011))
		self.RoundedRectangle(float64(self.Width())-10, percent*(float64(self.Height())-height), 10, height, 2, 2, 2, 2)
		self.Fill()
	}
}

func (self *Surface) DrawVerticalOverflow(boundsHeight, height, percent float64, oy OverflowY) {
	if height < 5 {
		height = 15
	}
	switch oy {
	case OVERFLOW_Y_SCROLL:
		self.SetSourceRGBA(color.HexRGBA(0x00000025))
		self.RoundedRectangle(float64(self.Width())-11, percent, 10, height, 2, 2, 2, 2)
		self.Fill()
	}
}

func (self *Surface) DrawHorizontalOverflow(boundsWidth float64) {
	ratio := float64(self.Width()) / boundsWidth
	self.SetSourceRGBA(color.HexRGBA(0x00000025))
	self.RoundedRectangle(float64(self.Height())-10, 50, 10, float64(self.Width())*ratio, 2, 2, 2, 2)
	self.Fill()
}
