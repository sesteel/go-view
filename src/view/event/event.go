package event

import (
)

type EventDispatcher struct {
	MouseEventDispatcher
	FocusEventDispatcher
	KeyboardEventDispatcher
}

func (self *EventDispatcher) MouseButtonPress(me Mouse) {
	SetFocus(&self.FocusEventDispatcher)
	for i := 0; i < len(self.mouseButtonPressHandlers); i++ {
		self.mouseButtonPressHandlers[i](me)
	}
}
