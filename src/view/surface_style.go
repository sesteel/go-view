package view

import ()

// DrawBackgroundStyle will draw a background as defined by the style
func (self *Surface) DrawBackgroundStyle(style Style) {
	x := style.PaddingLeft()
	y := style.PaddingTop()
	w := float64(self.Width()) - style.PaddingRight() - x 
	h := float64(self.Height()) - style.PaddingBottom() - y 
	self.DrawTopBorder(style)
	self.DrawRightBorder(style)
	self.DrawBottomBorder(style)
	self.DrawLeftBorder(style)
//	self.drawBackground(style)
	
	self.SetSourceRGBA(style.Background())
	self.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	self.Fill()
}

func (self *Surface) getBorderConstraints(style Style) (x, y, w, h, radius_x, radius_y, c1, c2 float64){
	x = style.PaddingLeft()
	y = style.PaddingTop()
	w = float64(self.Width()) - style.PaddingRight() - x 
	h = float64(self.Height()) - style.PaddingBottom() - y 
	ARC_TO_BEZIER := 0.55228475
    radius_x = 2.0
    radius_y = 2.0
    
    if radius_x > w - radius_x {
        radius_x = w / 2
    }
    
    if radius_y > h - radius_y {
        radius_y = h / 2
    }
		
    //approximate (quite close) the arc using a bezier curve
    c1 = ARC_TO_BEZIER * radius_x
    c2 = ARC_TO_BEZIER * radius_y
    return
}

func (self *Surface) DrawTopBorder(style Style) {
	x, y, w, _, radius_x, radius_y, c1, c2  := self.getBorderConstraints(style)
    self.SetSourceRGBA(style.BorderColorTop())
    self.NewPath()
    self.SetLineWidth(style.BorderWidthTop())
    self.MoveTo( x + radius_x, y)
    self.RelLineTo( w - 2 * radius_x, 0.0)
    self.RelCurveTo( c1, 0.0, radius_x, c2, radius_x, radius_y)
    self.StrokePreserve()	
}

func (self *Surface) DrawRightBorder(style Style) {
	x, y, w, h, radius_x, radius_y, c1, c2  := self.getBorderConstraints(style)
	self.SetSourceRGBA(style.BorderColorRight())
    self.NewPath()
    self.SetLineWidth(style.BorderWidthRight()) 
    self.MoveTo( x + w , y + radius_y )
    self.RelLineTo( 0, h - 2 * radius_y)
    self.RelCurveTo( 0.0, c2, c1 - radius_x, radius_y, -radius_x, radius_y)
    self.StrokePreserve()
}

func (self *Surface) DrawBottomBorder(style Style) {
	x, y, w, h, radius_x, radius_y, c1, c2  := self.getBorderConstraints(style)
	self.SetSourceRGBA(style.BorderColorBottom())
    self.NewPath()
    self.SetLineWidth(style.BorderWidthBottom()) 
    self.MoveTo( x + w - radius_x, y + h)
    self.RelLineTo ( -w + 2 * radius_x, 0)
    self.RelCurveTo ( -c1, 0, -radius_x, -c2, -radius_x, -radius_y)
    self.StrokePreserve()
}

func (self *Surface) DrawLeftBorder(style Style) {
	x, y, _, h, radius_x, radius_y, c1, c2  := self.getBorderConstraints(style)
	self.SetSourceRGBA(style.BorderColorLeft())
	self.NewPath()
	self.SetLineWidth(style.BorderWidthLeft()) 
    self.MoveTo( x, y + h - radius_y)
    self.RelLineTo (0, -h + 2 * radius_y)
    self.RelCurveTo (0.0, -c2, radius_x - c1, -radius_y, radius_x, -radius_y)
    self.StrokePreserve()
}

func (self *Surface) DrawTextCentered(text string, style Style) {
	self.SelectFontFace(style.FontName(), style.FontSlant(), style.FontWeight())
	self.SetFontSize(style.FontSize())
	self.SetSourceRGBA(style.Foreground())
	
	extents := self.TextExtents(text)
	x := float64(self.Width())/2 + style.PaddingLeft() - style.PaddingRight() - (extents.Width/2) 
	y := float64(self.Height()/2) + style.PaddingTop() - style.PaddingBottom() + (extents.Height/2) 
	self.MoveTo(x, y)
	self.ShowText(text)
}
