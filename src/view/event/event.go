package event

import (
)

type KeyUp struct {
	Key int
}

type KeyDown struct {
	KeyUp
}

type EventDispatcher struct {
	MouseEventDispatcher
	FocusEventDispatcher
}

func (self *EventDispatcher) MouseButtonPress(me Mouse) {
	SetFocus(&self.FocusEventDispatcher)
	for i := 0; i < len(self.mouseButtonPressHandlers); i++ {
		self.mouseButtonPressHandlers[i](me)
	}
}