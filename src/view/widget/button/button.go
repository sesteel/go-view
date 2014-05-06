// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package button

import (
	"view"
	"view/event"
	"view/color"
)

type State uint8

const (
	NORMAL State = 1 << iota
	HOVER 
	FOCUS
	DISABLED
	ERROR
	ACTIVATED
)

type button struct {
	view.DefaultComponent
	hoverStyle     view.Style
	activatedStyle view.Style
	state          State
}

func New(parent view.View, text string) *button {
	b := new(button)
	b.DefaultComponent = *view.NewComponent(parent, text)
	b.state = NORMAL
	
	b.hoverStyle = view.NewStyle() 
	b.hoverStyle.SetBackground(color.Gray3)
	
	b.activatedStyle = view.NewStyle()
	b.activatedStyle.SetBackground(color.Gray9)
	b.activatedStyle.SetBorderColor(color.Gray10)
	b.activatedStyle.SetForeground(color.Gray3)
	
	b.AddFocusGainedHandler(func() {
		b.state |= FOCUS
		b.Redraw()
	})

	b.AddFocusLostHandler(func() {
		b.state ^= FOCUS
		b.Redraw()
	})

	b.AddMouseEnterHandler(func(mp event.Mouse) {
		b.state |= HOVER
		b.Redraw()
	})

	b.AddMouseExitHandler(func(mp event.Mouse) {
		b.state ^= HOVER
		b.Redraw()
	})
	
	b.AddMouseButtonPressHandler(func(mp event.Mouse) {
		if mp.Button == event.MOUSE_BUTTON_LEFT {
			b.state |= ACTIVATED
			b.Redraw()
		}
	})
	
	b.AddMouseButtonReleaseHandler(func(mp event.Mouse) {
		if mp.Button == event.MOUSE_BUTTON_LEFT {
			b.state ^= ACTIVATED
			b.Redraw()
		}
	})
	
	return b
}

func (self *button) SetHoverStyle(style view.Style) {
	self.hoverStyle = style
}

func (self *button) HoverStyle() view.Style {
	return self.hoverStyle
}

func (self *button) Draw(s *view.Surface) {
	
	if self.state & ACTIVATED == ACTIVATED {
		s.DrawFilledBackground(self.activatedStyle)
		s.DrawTextCentered(self.DefaultComponent.Name(), self.activatedStyle)
	} else if self.state & HOVER == HOVER {
		s.DrawFilledBackground(self.hoverStyle)
		s.DrawTextCentered(self.DefaultComponent.Name(), self.hoverStyle)
	} else {
		s.DrawFilledBackground(self.Style())
		s.DrawTextCentered(self.DefaultComponent.Name(), self.Style())
	}
	
	
}
