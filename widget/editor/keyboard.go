// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	// "fmt"
	"view/event"
	"view/event/key"
)

func (self *Editor) initDefaultKeyboardHandler() {
	self.AddKeyPressHandler(func(k event.Keyboard) {
		start := self.Cursors[0]
		shiftMove := func() {
			end := self.Cursors[0]
			if k.Shift() {
				if self.Selection.Start.Line >= 0 {
					self.Select(Index(self.Selection.Start), Index(end.PreviousPos(self.Lines)))
				} else {
					self.Select(Index(start.Index), Index(end.PreviousPos(self.Lines)))
				}
			} else {
				self.ClearSelections()
			}
		}

		switch k.Value {
		case key.NONE,
			key.CAPS,
			key.LEFT_SHIFT,
			key.RIGHT_SHIFT,
			key.LEFT_ALT,
			key.RIGHT_ALT,
			key.LEFT_CMD,
			key.RIGHT_CMD,
			key.LEFT_CTRL,
			key.RIGHT_CTRL:
			return

		case key.ARROW_LEFT:
			if k.Ctrl() {
				self.MoveCursorsToPreviousToken()
			} else {
				self.MoveCursorsLeft()
			}
			shiftMove()

		case key.ARROW_RIGHT:
			if k.Ctrl() {
				self.MoveCursorsToNextToken()
			} else {
				self.MoveCursorsRight()
			}
			shiftMove()

		case key.ARROW_UP:
			self.MoveCursorsUp()
			shiftMove()

		case key.ARROW_DOWN:
			self.MoveCursorsDown()
			shiftMove()

		case key.BACKSPACE:
			self.DeleteCharBeforeCursors()

		case key.DELETE:
			self.DeleteCharAfterCursors()

		case key.RETURN:
			self.InsertCharAtCursors('\n')

		case key.HOME:
			self.MoveCursorToLineStart()
			shiftMove()

		case key.END:
			self.MoveCursorToLineEnd()
			shiftMove()

		case key.ESC:
			self.Selection = &Selection{Range{Index{-1, -1}, Index{-1, -1}}}

		default:
			self.InsertCharAtCursors(k.Rune())
		}
		self.Redraw()
	})
}
