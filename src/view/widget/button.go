// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package widget

import (
	"view"
	"view/event"
	"view/color"
	"view/widget/state"
)

type Button struct {
	view.DefaultComponent
	hoverStyle     view.Style
	activatedStyle view.Style
	state          state.State
}

func NewButton(parent view.View, text string) *Button {
	b := new(Button)
	b.DefaultComponent = *view.NewComponent(parent, text)
	b.state = state.NORMAL
	
	b.hoverStyle = view.NewStyle() 
	b.hoverStyle.SetBackground(color.Gray3)
	
	b.activatedStyle = view.NewStyle()
	b.activatedStyle.SetBackground(color.Gray9)
	b.activatedStyle.SetBorderColor(color.Gray10)
	b.activatedStyle.SetForeground(color.Gray3)
	
	b.state = state.NORMAL

	b.AddFocusGainedHandler(func() {
		b.state |= state.FOCUS
		b.Redraw()
	})

	b.AddFocusLostHandler(func() {
		b.state ^= state.FOCUS
		b.Redraw()
	})

	b.AddMouseEnterHandler(func(mp event.Mouse) {
		b.state |= state.HOVER
		b.Redraw()
	})

	b.AddMouseExitHandler(func(mp event.Mouse) {
		b.state ^= state.HOVER
		b.Redraw()
	})
	
	b.AddMouseButtonPressHandler(func(mp event.Mouse) {
		if mp.Button == event.MOUSE_BUTTON_LEFT {
			b.state |= state.ACTIVATED
			b.Redraw()
		}
	})
	
	b.AddMouseButtonReleaseHandler(func(mp event.Mouse) {
		if mp.Button == event.MOUSE_BUTTON_LEFT {
			b.state ^= state.ACTIVATED
			b.Redraw()
		}
	})
	
	return b
}

func (self *Button) SetHoverStyle(style view.Style) {
	self.hoverStyle = style
}

func (self *Button) Draw(s *view.Surface) {
	
	if self.state & state.ACTIVATED == state.ACTIVATED {
		s.DrawBackgroundStyle(self.activatedStyle)
		s.DrawTextCentered(self.DefaultComponent.Name(), self.activatedStyle)
	} else if self.state & state.HOVER == state.HOVER {
		s.DrawBackgroundStyle(self.hoverStyle)
		s.DrawTextCentered(self.DefaultComponent.Name(), self.hoverStyle)
	} else {
		s.DrawBackgroundStyle(self.Style())
		s.DrawTextCentered(self.DefaultComponent.Name(), self.Style())
	}
	
	
}
