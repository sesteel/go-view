package widget

import (
	"fmt"
	"ui/view/icon"
	"ui/view"
)

type Button struct {
	view.DefaultComponent
}

func NewButton(parent view.View, text string) *Button {
	return &Button{*view.NewComponent(parent, text)} 
}

func (self *Button) Draw() {
	fmt.Println("X")
	s := self.Surface()
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetLineWidth(1)
	s.MoveTo(0, 0)
	s.SetSourceRGB(0.3, 0.3, .3)
	x, y := self.Position()
	w, h := self.Size()
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 5, 5, 5, 5)
	s.StrokePreserve()
	s.Fill()
//	s.Flush()
	
//	s.SetSourceRGB(.7, .7, .7)
//	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 30, 30, 30, 30)
//	s.StrokePreserve()
	
	s.SetSourceRGB(.7, .9, .7)
	s.SelectFontFace("FontAwesome", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
	s.SetFontSize(16)
	s.MoveTo(50, 50)
	s.ShowText(icon.FA_ALIGN_JUSTIFY)
	
	s.SelectFontFace("Sans", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
	s.SetFontSize(26)
	s.MoveTo(70, 50)
//	s.ShowText("Stan Was Here @ 1234567890qwertyuiopasdfghjklzxcvbnm, ⊘Δ")
	s.ShowText(self.Text())
}

