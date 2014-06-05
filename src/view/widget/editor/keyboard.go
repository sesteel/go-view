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
			self.MoveCursorsLeft()

		case key.ARROW_RIGHT:
			self.MoveCursorsRight()

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

		default:
			self.InsertCharAtCursors(k.Rune())
		}

		self.Redraw()
	})
}
