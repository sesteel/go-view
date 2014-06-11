// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"view/event"
	"view/event/key"
)


func (self *Editor) addKeyboardHandler() {
	self.AddKeyPressHandler(func(k event.Keyboard) {
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

		case key.ARROW_RIGHT:
			if k.Ctrl() {
				self.MoveCursorsToNextToken()
			} else {
				self.MoveCursorsRight()
			}

		case key.ARROW_UP:
			self.MoveCursorsUp()

		case key.ARROW_DOWN:
			self.MoveCursorsDown()

		case key.BACKSPACE:
			self.DeleteCharBeforeCursors()
			
		case key.DELETE:
			self.DeleteCharAfterCursors()

		case key.RETURN:
			self.InsertCharAtCursors('\n')
		
		case key.HOME:
			self.MoveCursorToLineStart()
		
		case key.END:
			self.MoveCursorToLineEnd()		

		default:
			self.InsertCharAtCursors(k.Rune())
		}

		self.Redraw()
	})
}
