// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package scroll

import (
	"view"
	"view/color"
)

type verticalScroll struct {
	scroll
}

func NewVerticalScroll(parent view.View) Scroll {
	v := &verticalScroll{scroll{*view.NewComponent(parent, "vertical scroll"), 0, 0, 0}}
	style := v.Style()
	style.SetBackground(color.ScrollTrack)
	style.SetForeground(color.Red2)
	style.SetRadius(1)
	return v
}

func (self *verticalScroll) Draw(s *view.Surface) {
	w, h := float64(s.Width()), float64(s.Height())

	if h <= 0 {
		h = 1
	}

	ratio := self.scope / self.size
	pos := self.offset / self.size * h
	s.SetSourceRGBA(self.Style().Background())
	s.Rectangle(0, 0, w, h)
	s.Fill()
	s.SetSourceRGBA(self.Style().Foreground())
	s.RoundedRectangle(0, pos, w, h*ratio,
		self.Style().RadiusTopLeft(),
		self.Style().RadiusTopRight(),
		self.Style().RadiusBottomRight(),
		self.Style().RadiusBottomLeft())
	s.Fill()
}
