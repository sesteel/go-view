package view

import (
	"ui/view/icon"
)

type Component interface {
	View
}

type DefaultComponent struct {
	DefaultView
}

// NewComponent creates a new DefaultComponent.  DefaultComponent
// is generally not usable from a user perspective; it is useful 
// for referencing, compositing or embedding in other components,
// however.  
// 
// Aside, the name passed into this function reprsents this component
// programtically and should be unique.  It is intended
// to support  
func NewComponent(parent View, name string) *DefaultComponent {
	c := new(DefaultComponent)
	c.parent  = parent.Parent()
	c.surface = parent.Surface()
	c.text    = name
	c.width, c.height = parent.Size()
	return c
}

func (self *DefaultComponent) Draw() {
	s := self.surface
	
	// draw myself
	s.SetSourceRGB(0.2, 0.2, .2)
	s.Paint()
	
	s.SetAntialias(ANTIALIAS_SUBPIXEL)
	s.SetLineWidth(1)
	s.MoveTo(0, 0)
	s.SetSourceRGB(0.3, 0.3, .3)
	s.RoundedRectangle(float64(self.x), float64(self.y), float64(self.width), float64(self.height), 3, 3, 3, 3)
	s.StrokePreserve()
	s.Fill()
	s.Flush()
	
	s.SetSourceRGB(.7, .7, .7)
	s.RoundedRectangle(float64(self.x), float64(self.y), float64(self.width), float64(self.height), 3, 3, 3, 3)
	s.StrokePreserve()
	
	s.SetSourceRGB(.7, .9, .7)
	s.SelectFontFace("FontAwesome", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
	s.SetFontSize(26)
	s.MoveTo(50, 50)
	s.ShowText(icon.FA_ANDROID)
	s.SetSourceRGB(.9, .9, .9)
	s.SelectFontFace("Sans", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
	s.SetFontSize(26)
	s.MoveTo(0, 70)
	s.ShowText("Stan Was Here @ 1234567890qwertyuiopasdfghjklzxcvbnm, ⊘Δ")
	// !!!!! NOT A COMPOSITE !!!!!!
	// draw my children on myself
//	self.DefaultView.Draw()
}