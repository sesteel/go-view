// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package util

import (
	"view"
	"view/color"
)

type Scroll interface {
	view.View
	SetPosition(float64)
	SetRatio(float64)
}

type scroll struct {
	view.DefaultComponent
	position float64
	ratio    float64
}

// SetPosition sets the location of the handle
// given as a percentage between 0 and 1.
func (self *scroll) SetPosition(pos float64) {
	self.position = pos
}

// SetRatio sets the ratio which represents the
// height of the handle.
func (self *scroll) SetRatio(ratio float64) {
	self.ratio = ratio
}

type verticalScroll struct {
	scroll
}

func NewVerticalScroll(parent view.View) Scroll {
	v := &verticalScroll{scroll{*view.NewComponent(parent, "vertical scroll"), 0, 0}}
	style := v.Style()
	style.SetBackground(color.ScrollTrack)
	style.SetForeground(color.ScrollHandle)
	style.SetRadius(5)
	return v
}

func (self *verticalScroll) Draw(s *view.Surface) {
	w, h := float64(s.Width()), float64(s.Height())
	s.SetSourceRGBA(self.Style().Background())
	s.Rectangle(0, 0, w, h)
	s.Fill()

	s.SetSourceRGBA(self.Style().Foreground())
	s.RoundedRectangle(0, h*self.position, w, h*self.ratio,
		self.Style().RadiusTopLeft(),
		self.Style().RadiusTopRight(),
		self.Style().RadiusBottomRight(),
		self.Style().RadiusBottomLeft())
	s.Fill()
}
