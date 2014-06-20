// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	// "fmt"
	"view"
	"view/color"
	. "view/common"
	// "view/event"
	"view/tokenizer"
	"view/tokenizer/plaintext"
)

// TokenStyle provides the styling for various token types or subtypes.
type TokenStyle struct {
	Weight int
	Slant  int
	Color  color.RGBA
}

// Lines is an alias to hold lines of characters.
type Lines [][]tokenizer.Character

// Index represents a character position in the Lines data structure.
type Index struct {
	Line   int
	Column int
}

// Cursor is used to store the position of the cursor via a Index.
type Cursor Index

// Error represent a range of characters where an error has occurred.
type Error Range

type Map struct {
	Width   int
	Overlay color.RGBA
}

// Editor is a simple widget by which one can enter and edit text.
type Editor struct {
	view.DefaultComponent
	Tokenizer       tokenizer.Tokenizer
	Text            string // TODO Consider how to optimize this away as it gets created repeatedly
	Lines           Lines
	MarginColumn    int
	DrawMargin      bool
	MarginColor     color.RGBA
	DrawMap         bool
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
	Selections      []*Selection
	Errors          []*Error
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
		make([]*Selection, 0),
		make([]*Error, 0),
	}
	e.addKeyboardHandler()
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

	mapWidth := 125

	// s.Flush()
	s2 := view.NewSurface(view.FORMAT_ARGB32, s.Width()-mapWidth-1, s.Height()*3)
	defer s2.Destroy()

	s3 := view.NewSurface(view.FORMAT_ARGB32, mapWidth, s.Height())
	defer s3.Destroy()

	s3.Scale(.3333, .3333)
	// s3.SetSourceSurface(s2, 0, 0)
	// s3.Paint()

	// Draw Body
	self.drawBody(s2, s3)
	s.SetSourceSurface(s2, float64(mapWidth)+1, 0)
	s.Paint()
	s.Flush()

	// Draw Map
	s.SetSourceSurface(s3, 0, 0)
	s.Paint()
	s.SetSourceRGBA(color.Gray2)
	s.MoveTo(float64(mapWidth)+1, 0)
	s.SetLineWidth(1)
	s.LineTo(float64(mapWidth)+1, float64(s.Height()))
	s.Stroke()

}

func (self *Editor) drawBody(s *view.Surface, m *view.Surface) {
	style := self.Style()
	pl := style.PaddingLeft()
	//	pr := style.PaddingRight()
	pt := style.PaddingTop()
	pb := style.PaddingBottom()

	var se *view.TextExtents
	var ce *view.TextExtents

	// Set Font
	applyTextStyle := func(style view.Style) {
		s.SelectFontFace(style.FontName(), style.FontSlant(), style.FontWeight())
		s.SetFontSize(style.FontSize())
		s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
		s.SetSourceRGBA(style.Foreground())
		se = s.TextExtents(" ")
		ce = s.TextExtents("M")
	}
	applyTextStyle(style)

	self.drawMargin(s, ce.Xadvance, pl, pt, pb)

	PAD := 3.0
	var b Bounds
	b.Y = style.PaddingTop() + ce.Height + PAD
	b.X = style.PaddingLeft() + PAD
	b.Width = ce.Width
	b.Height = ce.Height
	s.SetSourceRGBA(style.Foreground())

	var pos int = 0
	updatePos := func() {
		for i := 0; i < len(self.Cursors); i++ {
			c := self.Cursors[i]
			if pos == self.Lines[c.Line][c.Column].Index {
				self.drawCursor(s, m, b.X, b.Y, ce.Width, ce.Height)
			}
		}
		pos++
	}

	// TODO: Optimize state changes out...
	var tokenClass tokenizer.TokenClass
	var tokenStyle *TokenStyle
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
			if sel := self.selectionAtIndex(idx); sel != nil {
				pad := (ce.Height * self.LineSpace) - (ce.Height)
				x, y := b.X, b.Y-c.Bounds.Height-(ce.Height/2)
				w, h := c.Bounds.Width, c.Bounds.Height+pad
				if c.Token.Type == tokenizer.TAB {
					w = se.Xadvance * float64(self.TabWidth-(col%self.TabWidth))
				}
				sel.drawCharBG(s, self.Lines, idx, x, y, w, h)
			}

			//fmt.Println(b)
			if c.Token.Type == tokenizer.NEWLINE {
				updatePos()
				self.drawWhitespace(s, 182, b)
				// fmt.Println(l, ce.Height, self.LineSpace)
				b.Y += ce.Height * self.LineSpace
				b.X = style.PaddingLeft() + PAD

			} else if c.Token.Type == tokenizer.SPACE {
				updatePos()
				self.drawWhitespace(s, 183, b)
				b.X += se.Xadvance

			} else if c.Token.Type == tokenizer.TAB {
				updatePos()
				self.drawWhitespace(s, 166, b)
				advance := float64(self.TabWidth - (col % self.TabWidth))
				b.X += se.Xadvance * advance

			} else {
				var ts *TokenStyle
				if tokenClass != c.Token.Type {
					tokenClass = c.Token.Type
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
					default:
						ts = defaultStyle
					}

					if ts != tokenStyle {
						s.SelectFontFace(style.FontName(), ts.Slant, ts.Weight)
						s.SetSourceRGBA(ts.Color)
						m.SelectFontFace(style.FontName(), ts.Slant, ts.Weight)
						m.SetSourceRGBA(ts.Color)
						tokenStyle = ts
					}
				}
				updatePos()

				s.DrawRune(c.Rune, b.X, b.Y)
				m.DrawRune(c.Rune, b.X, b.Y)
				// m.Rectangle(b.X, b.Y, ce.Width, ce.Height)
				// m.Fill()
				//ex := s.TextExtents(string(c.Rune))
				b.X += ce.Xadvance
			}
		}
	}
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

func (self *Editor) drawCursor(s, m *view.Surface, x, y, w, h float64) {
	s.Save()
	// TODO Allow different Styles Of Cursors
	s.SetSourceRGBA(color.Red1)
	s.SetLineCap(view.LINE_CAP_ROUND)
	s.SetLineWidth(1)
	s.MoveTo(x+1, y-h-2)
	s.LineTo(x+1, y+2)
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
		s.MoveTo(x, y1)
		s.LineTo(x, y2)
		s.Stroke()
	}
}
