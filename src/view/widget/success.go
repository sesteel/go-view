// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package widget

import (
	"view"
	"view/color"
)

type SuccessButton struct {
	view.DefaultComponent
}

func NewSuccessButton(parent view.View, text string) *SuccessButton {
	return &SuccessButton{*view.NewComponent(parent, text)} 
}

func (self *SuccessButton) Draw(s *view.Surface) {
	x, y := self.Position()
	w, h := self.Size()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	s.SetLineWidth(2)
	s.SetSourceRGBA(color.Green2)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.StrokePreserve()
	
	s.SetAntialias(view.ANTIALIAS_SUBPIXEL)
	p := view.NewLinearPattern(float64(x), float64(y), float64(x), float64(h))
	p.AddColorStop(0, color.Green1)
	p.AddColorStop(.5, color.Green1)
	p.AddColorStop(1, color.Green1)
	s.RoundedRectangle(float64(x), float64(y), float64(w), float64(h), 2, 2, 2, 2)
	s.SetSource(p)
	s.Fill()
	p.Destroy()
		
	s.SelectFontFace("Nimbus Sans L", view.FONT_SLANT_NORMAL, view.FONT_WEIGHT_NORMAL)
	s.SetFontSize(14)
	s.MoveTo(25, 15)
	s.SetSourceRGBA(color.White)
	s.ShowText(self.Name())
}

