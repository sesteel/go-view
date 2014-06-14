// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"view/event"
)

// Selection represents a character range where text has been
// selected; commonly used for copy and cutting operations.
type Selection struct {
	Range
}

func (self *Editor) addTextSelectionBehavior() {

	// dragging := false
	sel := &Selection{Range{Index{-1, -1}, Index{-1, -1}}}

	// Complex mouse behaviors
	self.AddMouseButtonPressHandler(func(ev event.Mouse) {
		switch ev.Button {
		case event.MOUSE_BUTTON_LEFT:
			self.MoveCursor(float64(ev.X), float64(ev.Y))
			kb := event.LastKeyboardState()
			if kb.CtrlOnly() {
				sel = &Selection{Range{Index{-1, -1}, Index{-1, -1}}}
			} else {
				self.Selections = make([]*Selection, 0)
			}
			self.Redraw()
		}
	})

	self.AddMouseButtonReleaseHandler(func(ev event.Mouse) {
		switch ev.Button {
		case event.MOUSE_BUTTON_LEFT:
			if sel.Start.Line > -1 {
				idx := self.FindClosestIndex(ev.X, ev.Y)
				l, c := idx.Line, idx.Column
				if l >= 0 && c >= 0 {
					sel = &Selection{Range{Index{-1, -1}, Index{-1, -1}}}
				}
			}
			self.Redraw()
		}
	})

	self.AddMousePositionHandler(func(ev event.Mouse) {
		if ev.LeftPressed {
			idx := self.FindClosestIndex(ev.X, ev.Y)
			if idx.Line >= 0 && idx.Column >= 0 && sel.Start.Line == -1 {
				sel.Start = idx
				self.Selections = append(self.Selections, sel)
			} else if idx.Line >= 0 && idx.Column >= 0 && sel.Start.Line > -1 {
				sel.End = idx
				self.MoveCursor(float64(ev.X), float64(ev.Y))
			}
			self.Redraw()
		}
	})
}
