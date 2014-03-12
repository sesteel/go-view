package event

import (

)

type EventDispatcher struct {
	mouseEnterHandlers []func(MouseEnter)
	mouseExitHandlers  []func(MouseExit)
}

func (self *EventDispatcher) AddMouseEnterHandler(f func(MouseEnter)) {
	self.mouseEnterHandlers = append(self.mouseEnterHandlers, f)
}

func (self *EventDispatcher) DispatchMouseEnter(f MouseEnter) {
	for i := 0; i < len(self.mouseEnterHandlers); i++ {
		self.mouseEnterHandlers[i](f)
	}
}

func (self *EventDispatcher) AddMouseExitHandler(f func(MouseExit)) {
	self.mouseExitHandlers = append(self.mouseExitHandlers, f)
}

func (self *EventDispatcher) DispatchMouseExit(f MouseExit) {
	for i := 0; i < len(self.mouseExitHandlers); i++ {
		self.mouseExitHandlers[i](f)
	}
}