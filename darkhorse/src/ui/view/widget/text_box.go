package widget

import (
	"ui/view"
	"ui/tokenizer"
	"ui/view/color"
)

type TextBox struct {
	view.DefaultComponent
}

func NewTextBox(parent view.View, text string) *TextBox {
	return &TextBox{*view.NewComponent(parent, text)} 
}

func (self *TextBox) Draw(s *view.Surface) {
	tkns := tokenizer.Tokenize(self.Text())
	
	// draw outline
	s.SetAntialias(view.ANTIALIAS_NONE)
	s.SetLineWidth(1)
	s.Rectangle(0, 0, float64(s.GetWidth()), float64(s.GetHeight()))
	s.SetSourceRGBA(color.Gray4)
	s.Stroke()
	
	// resize to draw within outline
	b := view.Bounds{0, 0, view.Size{float64(s.GetWidth()), float64(s.GetHeight())}}
	b.Width -= 1
	b.Height -= 1
	s.Rectangle(b.X, b.Y, b.Width, b.Height)
	s.SetSourceRGBA(color.White)
	s.Fill()
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	height := s.DrawWrappedPlainText(tkns, b, 0, self.Style())
	
	if height > float64(s.GetHeight()) {
		ratio := float64(s.GetHeight()) / height
		s.SetSourceRGBA(color.HexRGBA(0x00000010))
		s.Rectangle(float64(s.GetWidth()) - 10, 0, 10, float64(s.GetHeight()))
		s.Fill()
		s.SetSourceRGBA(color.HexRGBA(0x00000022))
		s.RoundedRectangle(float64(s.GetWidth()) - 8, 50, 6, float64(s.GetHeight()) * ratio, 1, 1, 1, 1)
		s.Fill()
	}
}
