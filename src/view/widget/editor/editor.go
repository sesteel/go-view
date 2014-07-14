// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"fmt"
	"math"
	// "time"
	"view"
	"view/color"
	. "view/common"
	"view/event"
	"view/tokenizer"
	"view/tokenizer/plaintext"
	"view/widget/scroll"
)

const ALIGN = 0.5

// Lines is an alias to hold lines of characters.
type Lines [][]tokenizer.Character

// Index represents a character position in the Lines data structure.
type Index struct {
	Line   int
	Column int
}

// Error represent a range of characters where an error has occurred.
type Error Range

// Editor is a simple widget by which one can enter and edit text.
type Editor struct {
	view.DefaultComponent
	Tokenizer       tokenizer.Tokenizer
	Text            string // TODO Consider how to optimize this away as it gets created repeatedly
	Lines           Lines
	MarginColumn    int
	DrawMargin      bool
	MarginColor     color.RGBA
	DrawScrollMap   bool
	scrollMap       scrollMap
	DrawGutter      bool
	DrawWhitespace  bool
	WhitespaceColor color.RGBA
	Keywords        map[string]bool
	Primitives      map[string]bool
	StringStyle     TokenStyle
	PrimitiveStyle  TokenStyle
	KeywordStyle    TokenStyle
	CommentStyle    TokenStyle
	TabWidth        int
	LineSpace       float64
	Cursors         []*Cursor
	Selection       *Selection
	Selections      []*Selection
	Errors          []*Error
	vscroll         scroll.Scroll
	linesDrawn      float64
	lineSurfaces    []*view.Surface
}

// New creates and returns a simple Editor object with its defaults set.
func New(parent view.View, name string, text string) *Editor {
	if len(text) == 0 {
		text = " "
	}

	tknzr := plaintext.New()

	e := &Editor{
		*view.NewComponent(parent, name),
		tknzr,
		text,
		tokenizer.ToLinesOfCharacters(tknzr.Tokenize(text)),
		80,
		true,
		color.RGBA{color.Pink1.R, color.Pink1.G, color.Pink1.B, 0.25},
		true,
		scrollMap{150, 0.25, color.Gray2.Alpha(0.25), color.Gray2, 1.0, 0, 0},
		true,
		true,
		color.RGBA{color.Cyan1.R, color.Cyan1.G, color.Cyan1.B, 0.5},
		make(map[string]bool),
		make(map[string]bool),
		TokenStyle{view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_NORMAL, color.Black},
		TokenStyle{view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_ITALIC, color.Black},
		TokenStyle{view.FONT_WEIGHT_BOLD, view.FONT_SLANT_NORMAL, color.Black},
		TokenStyle{view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_ITALIC, color.Gray8},
		4,
		1,
		make([]*Cursor, 0),
		&Selection{Range{Index{-1, -1}, Index{-1, -1}}},
		make([]*Selection, 0),
		make([]*Error, 0),
		nil,
		1,
		make([]*view.Surface, 0),
	}

	blue := color.Blue4
	e.Cursors = append(e.Cursors, &Cursor{Index{0, 0}, OUTLINE, &blue})

	e.vscroll = scroll.NewVerticalScroll(e)
	e.vscroll.Style().SetForeground(color.Blue5.Alpha(.15))
	e.vscroll.Style().SetBackground(color.Gray10.Alpha(.05))
	e.AddMouseWheelDownHandler(func(event.Mouse) {
		offset := e.vscroll.Offset()
		offset++
		e.ScrollTo(offset)
		e.Redraw()
	})

	e.AddMouseWheelUpHandler(func(event.Mouse) {
		offset := e.vscroll.Offset()
		offset--
		e.ScrollTo(offset)
		e.Redraw()
	})

	e.initDefaultKeyboardHandler()
	e.addTextSelectionBehavior()
	e.Style().SetPadding(0)
	e.Style().SetRadius(0)
	e.Style().SetForeground(color.Gray13)
	e.Style().SetBackground(color.Gray1)
	e.Style().SetFontName("Monospace")
	e.Style().SetFontSize(13)
	return e
}

func (self *Editor) Draw(s *view.Surface) {
	self.drawLines(s)
	s.Flush()
}

func (self *Editor) Animate(s *view.Surface) {
	self.drawCursors(s)
	s.Flush()
}

func (self *Editor) drawLines(s *view.Surface) {
	pos := Point{X: 0, Y: 0}
	for l := int(self.vscroll.Offset()); l < len(self.Lines); l++ {
		line := self.Lines[l]
		if pos.Y < float64(s.Height()) {
			var surf *view.Surface
			if len(self.lineSurfaces) == l {
				surf = self.drawLine(s, line)
				self.lineSurfaces = append(self.lineSurfaces, surf)
			} else {
				surf = self.lineSurfaces[l]
			}
			s.SetSourceSurface(surf, 0, pos.Y)
			s.Paint()
			pos.Y += float64(surf.Height()) * self.LineSpace
		} else {
			return
		}
	}
}

func (self *Editor) drawLine(s *view.Surface, line []tokenizer.Character) *view.Surface {
	b, extents := self.lineBounds(s, line)
	style := self.Style()
	ce := extents.Extents('M')
	se := extents.Extents(' ')
	defaultStyle := &TokenStyle{style.FontWeight(), style.FontSlant(), style.Foreground()}
	surf := view.NewSurface(view.FORMAT_ARGB32, int(b.Width), int(ce.Height-ce.Ybearing))
	extents = self.applyTextStyle(surf, style)

	var ts *TokenStyle
	for col := 0; col < len(line); col++ {
		c := &line[col]
		y := ce.Height + (ce.Height / 3)
		switch c.Token.Type {

		case tokenizer.NEWLINE:
			self.drawWhitespace(surf, 182, b)
			c.Bounds = &Bounds{Point: Point{b.X, y}, Size: Size{se.Xadvance, y}}
			b.X += se.Xadvance
			continue

		case tokenizer.SPACE:
			self.drawWhitespace(surf, 183, b)
			c.Bounds = &Bounds{Point: Point{b.X, y}, Size: Size{se.Xadvance, y}}
			b.X += se.Xadvance
			continue

		case tokenizer.TAB:
			// TODO - proper non-monospace tab support
			self.drawWhitespace(s, 166, b)
			advance := float64(self.TabWidth-(col%self.TabWidth)) * se.Xadvance
			c.Bounds = &Bounds{Point: Point{b.X, y}, Size: Size{advance, y}}
			b.X = math.Floor(b.X+advance) + ALIGN
			continue

		case tokenizer.IDENTIFIER:
			if self.Keywords[c.Token.Value] {
				ts = &self.KeywordStyle
			} else if self.Primitives[c.Token.Value] {
				ts = &self.PrimitiveStyle
			} else {
				ts = defaultStyle
			}

		case tokenizer.STRING_LITERAL:
			ts = &self.StringStyle

		case tokenizer.COMMENT:
			ts = &self.CommentStyle

		default:
			ts = defaultStyle
		}

		e := extents.Extents(c.Rune)
		surf.SelectFontFace(style.FontName(), ts.Slant, ts.Weight)
		surf.SetSourceRGBA(ts.Color)

		surf.DrawRune(c.Rune, b.X, y)
		c.Bounds = &Bounds{Point: Point{b.X, y}, Size: Size{e.Xadvance, y}}
		b.X = math.Floor(b.X+e.Xadvance) + ALIGN
	}
	return surf
}

func (self *Editor) lineBounds(s *view.Surface, line []tokenizer.Character) (Bounds, *extents) {
	style := self.Style()
	extents := self.applyTextStyle(s, style)
	b := Bounds{Point: Point{X: ALIGN, Y: ALIGN}, Size: Size{Width: 0, Height: 0}}
	se := extents.Extents(' ')
	for col := 0; col < len(line); col++ {
		c := &line[col]
		e := extents.Extents(c.Rune)
		if e.Height > b.Height {
			b.Height = (e.Height - e.Ybearing)
		}

		switch c.Token.Type {
		case tokenizer.SPACE, tokenizer.NEWLINE:
			b.Width += se.Xadvance
		case tokenizer.TAB:
			advance := float64(self.TabWidth - (col % self.TabWidth))
			b.Width += se.Xadvance * advance
		default:
			e := extents.Extents(c.Rune)
			b.Width += e.Xadvance
		}
	}
	return b, extents
}

func (self *Editor) applyTextStyle(s *view.Surface, style view.Style) *extents {
	s.SelectFontFace(style.FontName(), style.FontSlant(), style.FontWeight())
	s.SetFontSize(style.FontSize())
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetSourceRGBA(style.Foreground())
	name := fmt.Sprint(style.FontName(), style.FontSlant(), style.FontWeight())
	em := extentMaps[name]
	if em == nil {
		em = &extents{
			name,
			view.NewSurface(view.FORMAT_ARGB32, 5, 5),
			make(map[rune]*view.TextExtents, 0),
		}
		em.surface.SelectFontFace(style.FontName(), style.FontSlant(), style.FontWeight())
		em.surface.SetFontSize(style.FontSize())
		em.Extents(' ')
		em.Extents('M')
		extentMaps[name] = em
	}
	return em
}

func (self *Editor) drawWhitespace(s *view.Surface, r rune, b Bounds) {
	if self.DrawWhitespace {
		s.Save()
		s.SetSourceRGBA(self.WhitespaceColor)
		s.DrawRune(r, b.X, b.Y)
		s.Restore()
	}
}

func (self *Editor) drawCursors(s *view.Surface) {
	for _, cursor := range self.Cursors {
		if cursor.Line < len(self.Lines) {
			line := self.Lines[cursor.Line]
			if cursor.Column < len(line) {
				c := line[cursor.Column]
				cursor.Draw(s, c.Bounds, self)
			}
		}
	}
}
