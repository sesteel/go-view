package view

import (
	"ui/view/color"
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
	c.style   = NewStyle()
	c.text    = name
	c.width, c.height = 100, 20
	return c
}

func (self *DefaultComponent) Draw(surface *Surface) {
	style   := self.style
	parent  := self.parent
	
	// traverse up tree to find style until parent is nil
	for self.style == nil && parent != nil {
		style = self.parent.Style()
		parent = parent.Parent()
	}
	
	// cannot draw without style... draw fusia and black text
	if style == nil {
		msg  := "Error: No style set for component [" + self.text + "]'s hierarchy."
		surface.SelectFontFace("Sans", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
		surface.SetFontSize(16)
		te := surface.TextExtents(msg)
		surface.SetSourceRGB(1, 2, 1)
		surface.RoundedRectangle(float64(self.x), float64(self.y), float64(te.Width), float64(self.height), 0, 0, 0, 0)
		surface.StrokePreserve()
		surface.Fill()
		
		surface.SelectFontFace("Sans", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
		surface.SetFontSize(16)
		surface.SetSourceRGBA(color.Cyan2)
		surface.MoveTo(float64(self.x), float64(self.y) + te.Height)
		surface.ShowText(msg)
		
		return
	}
	
	if style.Antialias() {
		surface.SetAntialias(ANTIALIAS_SUBPIXEL)
	} else {
		surface.SetAntialias(ANTIALIAS_NONE)
	}
	
	surface.SetSourceRGBA(style.Background())
	surface.RoundedRectangle(float64(self.x), float64(self.y), float64(self.width), float64(self.height), 0, 0, 0, 0)
	surface.Fill()
	
	surface.SetLineWidth(style.BorderWidthTop())
	surface.RoundedRectangle(float64(self.x), float64(self.y), float64(self.width), float64(self.height), 0, 0, 0, 0)
	surface.StrokePreserve()
	
//	// draw myself
//	s.SetSourceRGB(0.2, 0.2, .2)
//	s.Paint()
//	
//	s.SetLineWidth(1)
//	s.MoveTo(0, 0)
//	s.SetSourceRGB(0.3, 0.3, .3)
//	s.RoundedRectangle(float64(self.x), float64(self.y), float64(self.width), float64(self.height), 3, 3, 3, 3)
//	
//	s.SetSourceRGB(.7, .7, .7)
//	s.RoundedRectangle(float64(self.x), float64(self.y), float64(self.width), float64(self.height), 3, 3, 3, 3)
//	s.StrokePreserve()
//	
//	s.SetSourceRGB(.7, .9, .7)
//	s.SelectFontFace("FontAwesome", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
//	s.SetFontSize(26)
//	s.MoveTo(50, 50)
//	s.ShowText(icon.FA_ANDROID)
//	s.SetSourceRGB(.9, .9, .9)
//	s.SelectFontFace("Sans", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
//	s.SetFontSize(26)
//	s.MoveTo(0, 70)
//	s.ShowText("Stan Was Here @ 1234567890qwertyuiopasdfghjklzxcvbnm, ⊘Δ")
//	// !!!!! NOT A COMPOSITE !!!!!!
//	// draw my children on myself
////	self.DefaultView.Draw()
}