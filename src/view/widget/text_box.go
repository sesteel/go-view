package widget

import (
	"fmt"
	"view"
	"view/event"
	"view/theme"
	"view/tokenizer"
)

type TextBox struct {
	view.DefaultComponent
	verticalOffset view.ScrollOffset
	tkns           []*tokenizer.Token
	//	model Document
}

func NewTextBox(parent view.View, text string) *TextBox {
	tb := &TextBox{*view.NewComponent(parent, text), 0, tokenizer.Tokenize(text)}
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
	return tb
}

func (self *TextBox) Draw(s *view.Surface) {

	// resize to draw within outline
	b := view.Bounds{0, 0, view.Size{float64(s.GetWidth()), float64(s.GetHeight())}}
	s.Rectangle(b.X, b.Y, b.Width, b.Height)
	s.SetSourceRGBA(theme.White)
	s.Fill()
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	drawnLines, lines, height := s.DrawWrappedPlainText(self.tkns, b, self.verticalOffset, self.Style())

	d := lines - drawnLines
	if d == 0 {
		d = 1
	}
	percent := float64(self.verticalOffset) / d
	fmt.Println(self.verticalOffset, lines, drawnLines, percent)

	if height > float64(s.GetHeight()) {
		s.DrawVerticalOverflow2(lines, drawnLines, percent, self.Style())
	}

	//	if width > float64(s.GetWidth()) {
	s.DrawHorizontalOverflow(height, self.Style())
	//	}
}
