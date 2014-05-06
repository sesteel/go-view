// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package checkbox

import (
	"view"
	"view/event"
	"view/color"
	"view/icon"
)

type State uint8

const (
	NORMAL State = 1 << iota
	HOVER 
	FOCUS
	DISABLED
	ERROR
	CHECKED
)

const (
	CHECKBOX_SIZE = 20
)

type checkbox struct {
	view.DefaultComponent
	text           string
	hoverStyle     view.Style
	checkStyle     view.Style
	state          State
}

// New returns a checkbox drawn with the text as a right justified label
func New(parent view.View, name, text string) *checkbox {
	hoverStyle := view.NewStyle() 
	hoverStyle.SetBackground(color.HexRGBA(0xFFFFFF11))
	hoverStyle.SetBorderColor(color.HexRGBA(0xFFFFFF00))
	hoverStyle.SetBorderWidth(1)
	
	checkStyle := view.NewStyle()
	checkStyle.SetFontName("FontAwesome")
	checkStyle.SetFontSize(14)
	
	b := &checkbox{*view.NewComponent(parent, name), text, hoverStyle, checkStyle, NORMAL}
	b.Style().SetBackground(color.Gray2)
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
			b.state ^= CHECKED
			b.Redraw()
		}
	})
	
	return b
}

func (self *checkbox) SetHoverStyle(style view.Style) {
	self.hoverStyle = style
}

func (self *checkbox) HoverStyle() view.Style {
	return self.hoverStyle
}

// Checked returns the state of the checkbox; true == checked
func (self *checkbox) Checked() bool {
	return self.state & CHECKED == CHECKED 
}

func (self *checkbox) Draw(s *view.Surface) {
	style := view.CloneStyle(self.Style())
	
	if self.state & HOVER == HOVER {
		hoverStyle := view.CloneStyle(self.hoverStyle)
		s.DrawFilledBackground(self.hoverStyle)
		s.DrawBackground(float64(style.PaddingLeft()), float64(style.PaddingTop()), CHECKBOX_SIZE, CHECKBOX_SIZE, style)
		pad := hoverStyle.PaddingLeft()
		hoverStyle.SetPaddingLeft(pad + CHECKBOX_SIZE)
		s.DrawTextLeftJustifed(self.text, hoverStyle)
	
	} else {		
		s.DrawBackground(float64(style.PaddingLeft()), float64(style.PaddingTop()), CHECKBOX_SIZE, CHECKBOX_SIZE, style)
		pad := style.PaddingLeft()
		style.SetPaddingLeft(pad + CHECKBOX_SIZE)
		s.DrawTextLeftJustifed(self.text, style)
		style.SetPaddingLeft(pad)
	}
	s.Flush()
	
	
	if self.state & CHECKED == CHECKED {
		s.ConfigureFont(self.checkStyle)
		s.MoveTo(float64(style.PaddingLeft()), float64(style.PaddingTop()) + 12)
		s.SetSourceRGBA(color.Gray10)
		s.ShowText(icon.FA_CHECK)
	} 
}
