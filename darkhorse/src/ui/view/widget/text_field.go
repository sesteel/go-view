package widget

import (
	"ui/view"
	"ui/view/color"
)

type TextField struct {
	view.DefaultComponent
}

func NewTextField(parent view.View, text string) *TextField {
	return &TextField{*view.NewComponent(parent, text)} 
}

func (self *TextField) Draw(s *view.Surface) {
	x, y := self.Position()
	w, h := self.Size()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetLineWidth(2)
	s.SetSourceRGBA(color.Gray4)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.StrokePreserve()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	p := view.NewLinearPattern(float64(x), float64(y), float64(x), float64(h))
	p.AddColorStop(0, color.Gray1)
	p.AddColorStop(.35, color.White)
	p.AddColorStop(.65, color.White)
	p.AddColorStop(1, color.Gray1)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.SetSource(p)
	s.Fill()
	p.Destroy()
	
//	tkns := text.Tokenize(self.Text())
//	s.SelectFontFace("Sans", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
//	s.SetFontSize(16)
////	te := surface.TextExtents(self.Text())
//	s.SetSourceRGB(1, 2, 1)
	
	s.SetSourceRGB(.4, .4, .4)
	s.SelectFontFace("Sans", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
	s.SetFontSize(16)
	s.MoveTo(0, 15)
	s.ShowText(self.Text())
}
