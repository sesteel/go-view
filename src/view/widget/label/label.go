// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package label

import (
	"view"
	"view/color"
)

type State uint8

const (
	NORMAL State = 1 << iota
	HOVER
	FOCUS
	DISABLED
	ERROR
)

type Label struct {
	view.DefaultComponent
	disabledStyle view.Style
	text          string
	state         State
}

func New(parent view.View, name, text string) *Label {
	style := view.NewStyle()
	style.SetBackground(color.Transparent)
	style.SetBorderColor(color.Transparent)
	style.SetTextAlignment(view.STYLE_TEXT_LEFT)
	ds := view.CloneAsDisabledStyle(style)
	lbl := &Label{*view.NewComponent(parent, name), ds, text, NORMAL}
	lbl.SetStyle(style)
	return lbl
}

func (self *Label) SetEnabled(enabled bool) {
	if enabled {
		self.state |= DISABLED
	} else {
		self.state ^= DISABLED
	}
}

func (self *Label) Draw(s *view.Surface) {
	var style view.Style
	if self.state & DISABLED == DISABLED {
		style = view.CloneStyle(self.disabledStyle)
	} else {
		style = view.CloneStyle(self.Style())
	}
	s.DrawFilledBackground(style)
	switch style.TextAlignment() {
		case view.STYLE_TEXT_LEFT:
			s.DrawTextLeftJustifed(self.text, style)
			
		case view.STYLE_TEXT_CENTERED:
			s.DrawTextCentered(self.text, style)
			
		default:
			s.DrawTextLeftJustifed(self.text, style)
			
	}
	
}
