// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"fmt"
	"log"
	"math"
	// "time"
	"github.com/sesteel/go-view"
	"github.com/sesteel/go-view/color"
	"github.com/sesteel/go-view/event"
	"github.com/sesteel/go-view/geometry"
	"github.com/sesteel/go-view/tokenizer"
	"github.com/sesteel/go-view/tokenizer/plaintext"
	"github.com/sesteel/go-view/widget/scroll"
)

const ALIGN = 0.5

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
	Tokenizer      tokenizer.Tokenizer
	text           string // TODO Consider how to optimize this away as it gets created repeatedly
	Lines          []tokenizer.Line
	DrawMargin     bool
	DrawScrollMap  bool
	scrollMap      scrollMap
	DrawGutter     bool
	DrawWhitespace bool
	Keywords       map[string]bool
	Primitives     map[string]bool
	Cursors        []*Cursor
	Selection      *Selection
	Selections     []*Selection
	Errors         []*Error
	vscroll        scroll.Scroll
	linesDrawn     float64
	lineSurfaces   []*view.Surface
	Style          *EditorStyle
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
		tokenizer.ToLines(tknzr.Tokenize(text)),
		true,
		true,
		scrollMap{150, 0.25, color.Gray2.Alpha(0.25), color.Gray2, 1.0, 0, 0},
		true,
		true,
		make(map[string]bool),
		make(map[string]bool),
		make([]*Cursor, 0),
		&Selection{Range{Index{-1, -1}, Index{-1, -1}}},
		make([]*Selection, 0),
		make([]*Error, 0),
		nil,
		1,
		make([]*view.Surface, 0),
		NewEditorStyle(),
	}

	blue := color.Blue4
	e.Cursors = append(e.Cursors, &Cursor{Index{0, 0}, OUTLINE, &blue, 0, 0})
	e.vscroll = scroll.NewVerticalScroll(e)

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
	return e
}

func (self *Editor) Draw(s *view.Surface) {
	s.SetSourceRGBA(self.Style.Background)
	s.Paint()
	self.drawLines(s)
	s.Flush()
}

func (self *Editor) Animate(s *view.Surface) {
	self.drawCursors(s)
	s.Flush()
}

func (self *Editor) drawLines(s *view.Surface) {
	pos := geometry.Point{X: 0, Y: 0}
	for l := int(self.vscroll.Offset()); l < len(self.Lines); l++ {
		line := self.Lines[l]
		if pos.Y < float64(s.Height()) {
			var surf *view.Surface
			if len(self.lineSurfaces) == l {
				surf = self.drawLine(s, line)
				self.lineSurfaces = append(self.lineSurfaces, surf)
			} else if self.lineSurfaces[l] == nil {
				surf = self.drawLine(s, line)
				self.lineSurfaces[l] = surf
			} else if len(line.Bounds) > 0 && line.Bounds[0].X == -1 {
				self.lineSurfaces[l].Destroy()
				surf = self.drawLine(s, line)
				self.lineSurfaces[l] = surf
			} else {
				surf = self.lineSurfaces[l]
			}
			if surf == nil {
				log.Println("Created nil surface on ", line)
			}
			s.SetSourceSurface(surf, 0, pos.Y)
			s.Paint()
			pos.Y += float64(surf.Height()) * self.Style.LineSpace
		} else {
			return
		}
	}
}

func (self *Editor) drawLine(s *view.Surface, line tokenizer.Line) *view.Surface {
	b, extents := self.lineBounds(s, line)
	style := self.Style
	ce := extents.Extents('M')
	se := extents.Extents(' ')
	defaultStyle := &TokenStyle{style.Font.Weight, style.Font.Slant, style.Foreground}
	surf := view.NewSurface(view.FORMAT_ARGB32, int(b.Width), int(ce.Height-ce.Ybearing))
	extents = self.applyTextStyle(surf, style)

	var ts *TokenStyle
	for col := 0; col < len(line.Characters); col++ {
		c := &line.Characters[col]
		var bounds geometry.Bounds
		y := ce.Height + (ce.Height / 3)
		switch c.Token.Type {

		case tokenizer.NEWLINE:
			self.drawWhitespace(surf, 182, b)
			bounds = geometry.Bounds{Point: geometry.Point{b.X, y}, Size: geometry.Size{se.Xadvance, y}}
			b.X += se.Xadvance
			line.Bounds[col] = bounds
			continue

		case tokenizer.SPACE:
			self.drawWhitespace(surf, 183, b)
			bounds = geometry.Bounds{Point: geometry.Point{b.X, y}, Size: geometry.Size{se.Xadvance, y}}
			b.X += se.Xadvance
			line.Bounds[col] = bounds
			continue

		case tokenizer.TAB:
			// TODO - proper non-monospace tab support
			// TODO - mulitple sequential tabs behave strangely
			self.drawWhitespace(s, 166, b)
			advance := float64(style.TabWidth-(col%style.TabWidth)) * se.Xadvance
			bounds = geometry.Bounds{Point: geometry.Point{b.X, y}, Size: geometry.Size{advance, y}}
			b.X = math.Floor(b.X+advance) + ALIGN
			line.Bounds[col] = bounds
			continue

		case tokenizer.IDENTIFIER:
			if self.Keywords[c.Token.Value] {
				ts = &style.KeywordStyle
			} else if self.Primitives[c.Token.Value] {
				ts = &style.PrimitiveStyle
			} else {
				ts = defaultStyle
			}

		case tokenizer.STRING_LITERAL:
			ts = &style.StringStyle

		case tokenizer.COMMENT:
			ts = &style.CommentStyle

		default:
			ts = defaultStyle
		}
		e := extents.Extents(c.Rune)
		surf.SelectFontFace(style.Font.Name, ts.Slant, ts.Weight)
		surf.SetSourceRGBA(ts.Color)
		surf.DrawRune(c.Rune, b.X, y)
		bounds = geometry.Bounds{Point: geometry.Point{b.X, 0}, Size: geometry.Size{e.Xadvance, e.Height}}
		b.X = math.Floor(b.X+e.Xadvance) + ALIGN
		line.Bounds[col] = bounds
	}
	return surf
}

func (self *Editor) lineBounds(s *view.Surface, line tokenizer.Line) (geometry.Bounds, *extents) {
	style := self.Style
	extents := self.applyTextStyle(s, style)
	b := geometry.Bounds{Point: geometry.Point{X: ALIGN, Y: ALIGN}, Size: geometry.Size{Width: 0, Height: 0}}
	se := extents.Extents(' ')
	for col := 0; col < len(line.Characters); col++ {
		c := &line.Characters[col]
		e := extents.Extents(c.Rune)
		if e.Height > b.Height {
			b.Height = (e.Height - e.Ybearing)
		}

		switch c.Token.Type {
		case tokenizer.SPACE, tokenizer.NEWLINE:
			b.Width += se.Xadvance
		case tokenizer.TAB:
			advance := float64(style.TabWidth - (col % style.TabWidth))
			b.Width += se.Xadvance * advance
		default:
			e := extents.Extents(c.Rune)
			b.Width += e.Xadvance
		}
	}
	return b, extents
}

func (self *Editor) applyTextStyle(s *view.Surface, style *EditorStyle) *extents {
	s.SelectFont(style.Font)
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetSourceRGBA(style.Foreground)
	name := fmt.Sprint(style.Font.Name, style.Font.Slant, style.Font.Weight) // TODO- , style.Font.Size?
	em := extentMaps[name]
	if em == nil {
		em = &extents{
			name,
			view.NewSurface(view.FORMAT_ARGB32, 5, 5),
			make(map[rune]*view.TextExtents, 0),
		}
		em.surface.SelectFont(style.Font)
		em.Extents(' ')
		em.Extents('M')
		extentMaps[name] = em
	}
	return em
}

func (self *Editor) drawWhitespace(s *view.Surface, r rune, b geometry.Bounds) {
	if self.DrawWhitespace {
		s.Save()
		s.SetSourceRGBA(self.Style.WhitespaceColor)
		s.DrawRune(r, b.X, b.Y)
		s.Restore()
	}
}

func (self *Editor) drawCursors(s *view.Surface) {
	for _, cursor := range self.Cursors {
		if cursor.Line < len(self.Lines) {
			line := self.Lines[cursor.Line]
			if cursor.Column <= len(line.Characters) {
				b := line.Bounds[cursor.Column]
				// if c.Bounds == nil {
				// 	fmt.Println(c, cursor.Line, cursor.Column, "is nil bounded")
				// }
				cursor.Draw(s, &b, self)
			}
		}
	}
}

// Invalidate causes all editor caches to empty
// and the widget to completely redraw itself.
func (self *Editor) Invalidate() {
	for i, _ := range self.lineSurfaces {
		self.destroyLineSurface(i)
	}
	self.Redraw()
}

func (self *Editor) removeLineSurface(line int) {
	self.destroyLineSurface(line)
	self.lineSurfaces = append(self.lineSurfaces[:line], self.lineSurfaces[line+1:]...)
}

func (self *Editor) destroyLineSurface(line int) {
	if line < 0 || line >= len(self.lineSurfaces) {
		return
	}
	fmt.Println(line, len(self.lineSurfaces))
	s := self.lineSurfaces[line]
	if s != nil {
		s.Destroy()
		self.lineSurfaces[line] = nil
	}
}

func (self *Editor) SetText(text string) {
	lines := tokenizer.ToLines(self.Tokenizer.Tokenize(text))
	self.Lines = lines
	self.text = text
}
