package view

import (
	"ui/tokenizer"
	"ui/view/color"
)

func (self *Surface) drawTextToken(tkn *tokenizer.Token, bounds Bounds, style Style) {
	if tkn.Selected {
		self.SetSourceRGBA(color.Selection)
		self.Rectangle(bounds.X, bounds.Y, bounds.Width, bounds.Height)
		self.Fill()
	}
	self.SelectFontFace(style.FontName(), style.FontSlant(), style.FontWeight())
	self.SetFontSize(style.FontSize())
	self.SetSourceRGBA(style.Foreground())
	self.MoveTo(bounds.X, bounds.Y)
	self.ShowText(tkn.Value)
}

// DrawWrappedPlainText uses the Style's foreground color to draw plain
// ascii formatted text within the bounds.  It stops rendering at last visible
// line, but continues to calculate total height.
// returns height 
func (self *Surface) DrawWrappedPlainText(tokens []*tokenizer.Token, bounds Bounds, offset ScrollOffset, style Style) float64 {
	bounds.X += style.PaddingLeft()
	bounds.Y += style.PaddingTop()
	bounds.Width -= (style.PaddingLeft() + style.PaddingRight())
	bounds.Height -= (style.PaddingTop() + style.PaddingBottom())
	
	var lineHeight, x, y float64 = 0, 0, 0
	
//	self.SetSourceRGBA(style.Background())
//	self.Rectangle(bounds.X, bounds.Y, bounds.Width, bounds.Height)
//	self.Fill()
	
	self.SelectFontFace(style.FontName(), style.FontSlant(), style.FontWeight())
	self.SetFontSize(style.FontSize())
	self.SetSourceRGBA(style.Foreground())
	spaceExtents := self.TextExtents("M")
	spaceExtents.Width /= 1.5
	spaceExtents.Height *= 1.75
	
	//y = bounds.Y 
	y += spaceExtents.Height 	
	self.SetFontOptions(defaultFontOptions)
	line := 0
	
	selected := make([]*tokenizer.Token, 0)
	
	for i := 0; i < len(tokens); i++ {
		t   := tokens[i]	
		e   := self.TextExtents(t.Value)
		
		if t.Type == tokenizer.NEWLINE {
			x = 0
			y += spaceExtents.Height
			line++
			continue
		}	
		
		if t.Type == tokenizer.TAB {
			x += spaceExtents.Width * float64(style.TabWidth())
			continue
		}
	
		if e.Height > lineHeight {
			lineHeight = e.Height
		}
		
		if x + e.Width > bounds.Width {
			y += spaceExtents.Height
			x = 0
			line++
		}
		
		var b Bounds
		b.X = bounds.X + x
		b.Y = bounds.Y + y
		b.Width  = e.Xadvance - b.X
		b.Height = e.Yadvance - b.Y
		
		if t.Selected {
			selected = append(selected, t)
		}
		
		if y < bounds.Height {
			self.drawTextToken(t, b, style)
		}
		
		x += e.Xadvance
		y += e.Yadvance 
	}
	
	
	
//	self.SetSourceRGBA(color.HexRGBA(0xFF000055))
//	self.Rectangle(bounds.X, bounds.Y, bounds.Width, bounds.Height)
//	self.Fill()

	return y + style.PaddingBottom()
}
