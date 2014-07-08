// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"fmt"
	"view"
	"view/color"
	. "view/common"
	"view/event"
	"view/tokenizer"
	"view/tokenizer/plaintext"
	"view/widget/scroll"
)

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
	Cursors         []Cursor
	Selection       *Selection
	Selections      []*Selection
	Errors          []*Error
	vscroll         scroll.Scroll
	linesDrawn      float64
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
		1.7,
		[]Cursor{Cursor{0, 0}},
		&Selection{Range{Index{-1, -1}, Index{-1, -1}}},
		make([]*Selection, 0),
		make([]*Error, 0),
		nil,
		1,
	}

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
	e.Style().SetFontSize(14)
	return e
}

func (self *Editor) Draw(s *view.Surface) {
	s.SetSourceRGBA(self.Style().Background())
	s.Paint()
	s2 := view.NewSurface(view.FORMAT_ARGB32, s.Width()-int(self.scrollMap.Width)-1, s.Height()*int(1/self.scrollMap.Scale))
	defer s2.Destroy()

	s3 := view.NewSurface(view.FORMAT_ARGB32, int(self.scrollMap.Width), s.Height())
	defer s3.Destroy()
	if self.DrawScrollMap {
		defer self.scrollMap.draw(s, s3)
	}
	s3.Scale(self.scrollMap.Scale, self.scrollMap.Scale)

	// Draw Body
	self.drawBody(s2, s3)
	if self.DrawScrollMap {
		s.SetSourceSurface(s2, self.scrollMap.Width+1, 0)
	} else {
		s.SetSourceSurface(s2, 0, 0)
	}
	s.Paint()
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

func (self *Editor) drawBody(s *view.Surface, m *view.Surface) {

	const PAD = 3.0
	style := self.Style()

	// Configure Style and Get Metrics
	extents := self.applyTextStyle(s, style)
	ce := extents.Extents('M')
	se := extents.Extents(' ')
	self.drawMargin(s, ce.Xadvance, style.PaddingLeft(), style.PaddingTop(), style.PaddingBottom())
	s.SetSourceRGBA(style.Foreground())

	lineHeight := ce.Height * self.LineSpace
	self.scrollMap.viewStart = float64(s.Height())
	self.scrollMap.viewStop = 0

	var b Bounds
	b.Y = style.PaddingTop() + ce.Height + PAD
	b.X = style.PaddingLeft() + PAD
	b.Width = ce.Width
	b.Height = ce.Height

	poff := self.vscroll.Offset() * lineHeight
	var pos int = 0
	updateCursorPosition := func(x, y float64) {
		for i := 0; i < len(self.Cursors); i++ {
			c := self.Cursors[i]
			if pos == self.Lines[c.Line][c.Column].Index {
				self.drawCursor(s, m, x, y, ce.Width, ce.Height)
			}
		}
		pos++
	}

	// TODO: Optimize state changes out...
	// var tokenClass tokenizer.TokenClass
	// var tokenStyle *TokenStyle
	defaultStyle := &TokenStyle{style.FontWeight(), style.FontSlant(), style.Foreground()}

	for l := 0; l < len(self.Lines); l++ {
		line := self.Lines[l]
		for col := 0; col < len(line); col++ {
			idx := Index{l, col}
			c := &line[col]

			// Set the character bounds subtracting the
			// character height to change coord space.
			c.Bounds = b
			c.Bounds.Y -= ce.Height

			// Draw Text Selection if Present
			self.drawTextSelection(s, idx, ce, se, c, b)

			if b.Y < (float64(s.Height())*self.scrollMap.Scale)+poff && poff <= b.Y {
				updateCursorPosition(b.X, b.Y)
			}

			if c.Token.Type == tokenizer.NEWLINE {
				self.drawWhitespace(s, 182, b)
				b.Y += lineHeight
				b.X = style.PaddingLeft() + PAD

			} else if c.Token.Type == tokenizer.SPACE {
				self.drawWhitespace(s, 183, b)
				b.X += se.Xadvance

			} else if c.Token.Type == tokenizer.TAB {
				self.drawWhitespace(s, 166, b)
				advance := float64(self.TabWidth - (col % self.TabWidth))
				b.X += se.Xadvance * advance

			} else {
				var ts *TokenStyle
				switch c.Token.Type {
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

				if b.Y < (float64(s.Height())*self.scrollMap.Scale)+poff && poff <= b.Y {
					s.SelectFontFace(style.FontName(), ts.Slant, ts.Weight)
					s.SetSourceRGBA(ts.Color)
					s.DrawRune(c.Rune, b.X, b.Y-poff)

					if self.scrollMap.viewStart > b.Y {
						self.scrollMap.viewStart = b.Y - lineHeight
					}

					if self.scrollMap.viewStop < b.Y {
						self.scrollMap.viewStop = b.Y
					}
				}

				if self.DrawScrollMap && b.Y < float64(m.Height())/self.scrollMap.Scale {
					m.SelectFontFace(style.FontName(), ts.Slant, ts.Weight)
					m.SetSourceRGBA(ts.Color)
					m.DrawRune(c.Rune, b.X, b.Y)
				}

				ad := extents.Extents(c.Rune)
				b.X += ad.Xadvance
			}
		}
	}
	fmt.Println(self.vscroll.Offset(), lineHeight)
}

func (self *Editor) selectionAtIndex(i Index) *Selection {
	if len(self.Selections) > 0 {
		for j := 0; j < len(self.Selections); j++ {
			s := *self.Selections[j]
			if s.IndexInSelection(i) {
				return &s
			}
		}
	}
	return nil
}

func (self *Editor) drawTextSelection(surface *view.Surface, idx Index, ce, se *view.TextExtents, c *tokenizer.Character, b Bounds) {
	if sel := self.selectionAtIndex(idx); sel != nil {
		pad := (ce.Height * self.LineSpace) - (ce.Height)
		x, y := b.X, b.Y-c.Bounds.Height-(ce.Height/2)
		w, h := c.Bounds.Width, c.Bounds.Height+pad
		if c.Token.Type == tokenizer.TAB {
			w = se.Xadvance * float64(self.TabWidth-(idx.Column%self.TabWidth))
		}
		sel.drawCharBG(surface, self.Lines, idx, x, y, w, h)
	}
}

func (self *Editor) drawCursor(s, m *view.Surface, x, y, w, h float64) {
	s.Save()
	// TODO Allow different Styles Of Cursors
	s.SetSourceRGBA(color.Red1)
	s.SetLineCap(view.LINE_CAP_ROUND)
	s.SetLineWidth(1)
	s.MoveTo(x+1-0.5, y-h-2-0.5)
	s.LineTo(x+1-0.5, y+2-0.5)
	s.Stroke()
	s.Restore()
}

func (self *Editor) drawWhitespace(s *view.Surface, r rune, b Bounds) {
	if self.DrawWhitespace {
		s.Save()
		s.SetSourceRGBA(self.WhitespaceColor)
		s.DrawRune(r, b.X, b.Y)
		s.Restore()
	}
}

func (self *Editor) drawMargin(s *view.Surface, xAdvance, padL, padT, padB float64) {
	// Draw Margin (Vertical Line)
	if self.DrawMargin {
		x := padL + (float64(self.MarginColumn) * xAdvance)
		y1 := padT
		y2 := float64(s.Height()) - padB
		s.SetLineWidth(1)
		s.SetSourceRGBA(self.MarginColor)
		s.MoveTo(x-0.5, y1-0.5)
		s.LineTo(x-0.5, y2-0.5)
		s.Stroke()
	}
}
