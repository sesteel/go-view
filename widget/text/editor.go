package text

import (
	// "log"
	"time"
	"view"
	"view/color"
	"view/common"
	"view/tokenizer"
	"view/tokenizer/plaintext"
)

const ALIGN = 0.5

// Index represents a character position in the Lines data structure.
type Index struct {
	Line   int
	Column int
}

// Range represents an group of characters defined by starting and ending (line, column) positions.
type Range struct {
	Start Index
	End   Index
}

type textStyle struct {
	Font    *view.Font
	Color   color.RGBA
	extents map[rune]*view.TextExtents
}

type styler interface {
	style(tokenizer.TokenClass) *textStyle
}

type Editor struct {
	view.DefaultComponent
	Background color.RGBA
	textStyles map[tokenizer.TokenClass]*textStyle
	tokenizer  tokenizer.Tokenizer
	lines      [][]*Character
	lineSpace  float64
	cursor     *Cursor
	Selections []*Selection
	offset     float64
	text       string // TODO Consider how to optimize this away as it gets created repeatedly
	dirty      bool
}

func NewEditor(parent view.View, name string, text string) *Editor {
	if text[len(text)-1] != '\n' {
		text += "\n"
	}
	tknzr := plaintext.New()
	e := &Editor{
		*view.NewComponent(parent, name),
		color.White,
		make(map[tokenizer.TokenClass]*textStyle),
		tknzr,
		nil,
		1.75,
		&Cursor{Index{0, 0}, DEFAULT, &color.RGBA{.05, .05, .05, 1}, 2, time.Now()},
		make([]*Selection, 0),
		0.0,
		text,
		true,
	}
	e.SetStyle(tokenizer.IDENTIFIER, &view.Font{"FreeMono", view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_NORMAL, 15}, color.Gray11)
	e.initDefaultKeyboardHandler()
	e.lines = e.toLines(tknzr.Tokenize(text))
	return e
}

func (self *Editor) SetText(text string) {
	self.lines = self.toLines(self.tokenizer.Tokenize(text))
	self.text = text
}

func (self *Editor) SetStyle(class tokenizer.TokenClass, f *view.Font, c color.RGBA) {
	ts := &textStyle{f, c, make(map[rune]*view.TextExtents)}
	self.textStyles[class] = ts
	s := view.NewSurface(view.FORMAT_ARGB32, int(f.Size)+10, int(f.Size)+10)
	ts.Font.Configure(s)
	defer s.Destroy()
	h := s.TextExtents(string('M'))
	for i := 0; i < 256; i++ {
		r := rune(i)
		e := s.TextExtents(string(r))
		e.Height = h.Height
		ts.extents[r] = e
	}
}

func (self *Editor) style(t tokenizer.TokenClass) *textStyle {
	switch t {
	case tokenizer.STRING_LITERAL:
	case tokenizer.NUMBER_LITERAL:
	default:

	}
	return self.textStyles[tokenizer.IDENTIFIER]
}

func (self *Editor) Text() string {
	return self.text
}

func (self *Editor) Draw(s *view.Surface) {
	// var currFont view.Font
	s.SetSourceRGBA(self.Background)
	s.Paint()
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	style := self.style(tokenizer.IDENTIFIER)
	style.Font.Configure(s)
	s.SetSourceRGBA(style.Color)
	x, y := 0.0, 0.0
	h := float64(s.Height())

	for _, line := range self.lines {
		y += self.drawLine(s, line, x, y).Height * self.lineSpace
		x = 0
		if y > h {
			break
		}
	}
}

func (self *Editor) drawLine(s *view.Surface, line []*Character, x, y float64) *view.TextExtents {
	style := self.style(tokenizer.IDENTIFIER)
	var max *view.TextExtents = style.extents['.']
	xx := 0.0
	// Optimization did not seem to make a difference
	w := float64(s.Width())
	for _, char := range line {
		style = self.style(char.Token.Type)
		e := style.extents[char.Rune]
		if e.Height > max.Height {
			max = e
		}
		xx += e.Xadvance
		// Optimization did not seem to make a meaningful difference
		if xx > w {
			break
		}
	}

	y += max.Height

	for _, char := range line {
		style := self.style(char.Token.Type)
		e := style.extents[char.Rune]
		// style.Font.Configure(s)
		char.Bounds.X = x
		char.Bounds.Y = y
		char.Bounds.Width = e.Width
		char.Bounds.Height = e.Height
		s.DrawRune(char.Rune, x, y)
		x += e.Xadvance
		// Optimization did not seem to make a difference
		if x > w {
			break
		}
	}
	return max
}

func (self *Editor) drawCursors(s *view.Surface) {
	c := self.cursor
	c.Draw(s, self)
}

func (self *Editor) Animate(s *view.Surface) {
	self.drawCursors(s)
	s.Flush()
}

func (self *Editor) initTokenizer() {
	go func() {
		text := self.text
		for {
			if text != self.text {
				text = self.text
				self.lines = self.toLines(self.tokenizer.Tokenize(self.text))
				if text == self.text {
					self.Redraw()
				} else {
					continue
				}
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

func (self *Editor) toLines(tkns []*tokenizer.Token) [][]*Character {
	lines := make([][]*Character, 0, 0)
	line := make([]*Character, 0)
	pos := 0
	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]
		for _, r := range tkn.Value {
			line = append(line, &Character{self, pos, tkn, r, common.Bounds{common.Point{0, 0}, common.Size{0, 0}}})
			pos++
		}

		if tkn.Value == "\n" {
			lines = append(lines, line)
			line = make([]*Character, 0)
		}
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}
	return lines
}
