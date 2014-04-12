package event

import (

)

type FocusNotifier interface {
	AddFocusGainedHandler(func())
	AddFocusLostHandler(func())
}

type FocusHandler interface {
	FocusGained()
	FocusLost()
}

type FocusEventDispatcher struct {
	focusGainedHandlers []func()
	focusLostHandlers   []func()
}


func (self *FocusEventDispatcher) AddFocusGainedHandler(f func()) {
	self.focusGainedHandlers = append(self.focusGainedHandlers, f)
}

func (self *FocusEventDispatcher) RemoveFocusGainedHandler(f func()) {
	// TODO - implement
}

func (self *FocusEventDispatcher) FocusGained() {
	for i := 0; i < len(self.focusGainedHandlers); i++ {
		self.focusGainedHandlers[i]()
	}
}

func (self *FocusEventDispatcher) AddFocusLostHandler(f func()) {
	self.focusGainedHandlers = append(self.focusGainedHandlers, f)
}

func (self *FocusEventDispatcher) RemoveFocusLostHandler(f func()) {
	// TODO - implement
}

func (self *FocusEventDispatcher) FocusLost() {
	for i := 0; i < len(self.focusLostHandlers); i++ {
		self.focusLostHandlers[i]()
	}
}