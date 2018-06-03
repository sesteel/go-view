// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package scroll

import (
	"github.com/sesteel/go-view"
)

type verticalScroll struct {
	scroll
}

func NewVerticalScroll(parent view.View) Scroll {
	return &verticalScroll{newScroll(parent, "vertical scroll")}
}

func (self *verticalScroll) Draw(s *view.Surface) {
	w, h := float64(s.Width()), float64(s.Height())

	if h <= 0 {
		h = 1
	}

	ratio := self.scope / self.size
	pos := self.offset / self.size * h
	s.SetSourceRGBA(self.Style.TrackColor)
	s.Rectangle(0, 0, w, h)
	s.Fill()
	s.SetSourceRGBA(self.Style.HandleColor)
	r := self.Style.HandleRadius
	s.RoundedRectangle(0, pos, w, h*ratio, r, r, r, r)
	s.Fill()
}
