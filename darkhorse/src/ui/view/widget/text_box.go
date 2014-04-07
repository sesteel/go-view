package widget

import (
	"ui/view"
	"ui/tokenizer"
	"ui/view/color"
	"ui/view/event"
	"fmt"
)

type TextBox struct {
	view.DefaultComponent
	verticalOffset view.ScrollOffset
	tkns []*tokenizer.Token
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
	return tb
}

func (self *TextBox) Draw(s *view.Surface) {
	
	
	// draw outline
//	s.SetAntialias(view.ANTIALIAS_NONE)
//	s.SetLineWidth(1)
//	s.Rectangle(0, 0, float64(s.GetWidth()), float64(s.GetHeight()))
//	s.SetSourceRGBA(color.Gray4)
//	s.Stroke()
	
	// resize to draw within outline
	b := view.Bounds{0, 0, view.Size{float64(s.GetWidth()), float64(s.GetHeight())}}
	s.Rectangle(b.X, b.Y, b.Width, b.Height)
	s.SetSourceRGBA(color.White)
	s.Fill()
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	drawnLines, lines, height := s.DrawWrappedPlainText(self.tkns, b, self.verticalOffset, self.Style())
	
	fmt.Println("dl:", drawnLines, "l:", lines, "h:", height, float64(s.GetHeight()), height, float64(s.GetHeight()) * (drawnLines/lines))
	
	if height > float64(s.GetHeight()) {
		s.DrawVerticalOverflow(height, float64(s.GetHeight()) * (drawnLines/lines), float64(self.verticalOffset)/height, self.Style())
	}
	
//	if width > float64(s.GetWidth()) {
		s.DrawHorizontalOverflow(height, self.Style())
//	}
}
