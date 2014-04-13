package widget

import (
	"view"
	"view/theme"
)

type Button struct {
	view.DefaultComponent
}

func NewButton(parent view.View, text string) *Button {
	return &Button{*view.NewComponent(parent, text)} 
}

func (self *Button) Draw(s *view.Surface) {
	x, y := self.Position()
	w, h := self.Size()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetLineWidth(2)
	s.SetSourceRGBA(theme.Gray4)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.StrokePreserve()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	p := view.NewLinearPattern(float64(x), float64(y), float64(x), float64(h))
	p.AddColorStop(0, theme.Gray3)
	p.AddColorStop(.5, theme.Gray3)
	p.AddColorStop(1, theme.Gray3)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.SetSource(p)
	s.Fill()
	p.Destroy()
	
	s.SelectFontFace("Nimbus Sans L", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
	s.SetFontSize(14)
	s.MoveTo(25, 15)
	s.SetSourceRGBA(theme.Gray5)
	s.ShowText(self.Name())
}

