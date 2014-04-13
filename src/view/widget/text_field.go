package widget

import (
	"view"
	"view/theme"
)

type TextField struct {
	view.DefaultComponent
}

func NewTextField(parent view.View, name string) *TextField {
	return &TextField{*view.NewComponent(parent, name)} 
}

func (self *TextField) Draw(s *view.Surface) {
	x, y := self.Position()
	w, h := self.Size()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetLineWidth(2)
	s.SetSourceRGBA(theme.Gray4)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.StrokePreserve()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	p := view.NewLinearPattern(float64(x), float64(y), float64(x), float64(h))
	p.AddColorStop(0, theme.Gray1)
	p.AddColorStop(.35, theme.White)
	p.AddColorStop(.65, theme.White)
	p.AddColorStop(1, theme.Gray1)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.SetSource(p)
	s.Fill()
	p.Destroy()
	
	s.SetSourceRGB(.4, .4, .4)
	s.SelectFontFace("Sans", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
	s.SetFontSize(16)
	s.MoveTo(0, 15)
	s.ShowText(self.Name())
}
