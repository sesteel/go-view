package event

import ()

type MouseButton int

const (
	MOUSE_BUTTON_NONE MouseButton = iota
	MOUSE_BUTTON_LEFT
	MOUSE_BUTTON_MIDDLE
	MOUSE_BUTTON_RIGHT
)

type Mouse struct {
	Button MouseButton
	X, Y   int
}

type MouseNotifier interface {
	AddMouseEnterHandler(func(Mouse))
	AddMouseExitHandler(func(Mouse))
	AddMousePositionHandler(func(Mouse))
	AddMouseWheelUpHandler(func(Mouse))
	AddMouseWheelDownHandler(func(Mouse))
	AddMouseButtonReleaseHandler(func(Mouse))
	AddMouseButtonPressHandler(func(Mouse))
}

type MouseHandler interface {
	MouseEnter(Mouse)
	MouseExit(Mouse)
	MousePosition(Mouse)
	MouseWheelUp(Mouse)
	MouseWheelDown(Mouse)
	MouseButtonPress(Mouse)
	MouseButtonRelease(Mouse)
}

type MouseEventDispatcher struct {
	mouseEnterHandlers         []func(Mouse)
	mouseExitHandlers          []func(Mouse)
	mousePositionHandlers      []func(Mouse)
	mouseWheelDownHandlers     []func(Mouse)
	mouseWheelUpHandlers       []func(Mouse)
	mouseButtonPressHandlers   []func(Mouse)
	mouseButtonReleaseHandlers []func(Mouse)
}

// -- ////////////

func (self *MouseEventDispatcher) AddMouseEnterHandler(f func(Mouse)) {
	self.mouseEnterHandlers = append(self.mouseEnterHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMouseEnterHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MouseEnter(me Mouse) {
	for i := 0; i < len(self.mouseEnterHandlers); i++ {
		self.mouseEnterHandlers[i](me)
	}
}

// -- ////////////

func (self *MouseEventDispatcher) AddMouseExitHandler(f func(Mouse)) {
	self.mouseExitHandlers = append(self.mouseExitHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMouseExitHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MouseExit(me Mouse) {
	for i := 0; i < len(self.mouseExitHandlers); i++ {
		self.mouseExitHandlers[i](me)
	}
}

// -- ////////////

func (self *MouseEventDispatcher) AddMousePositionHandler(f func(Mouse)) {
	self.mousePositionHandlers = append(self.mousePositionHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMousePositionHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MousePosition(me Mouse) {
	for i := 0; i < len(self.mousePositionHandlers); i++ {
		self.mousePositionHandlers[i](me)
	}
}

// -- ////////////

func (self *MouseEventDispatcher) AddMouseWheelUpHandler(f func(Mouse)) {
	self.mouseWheelUpHandlers = append(self.mouseWheelUpHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMouseWheelUpHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MouseWheelUp(me Mouse) {
	for i := 0; i < len(self.mouseWheelUpHandlers); i++ {
		self.mouseWheelUpHandlers[i](me)
	}
}

// -- ////////////

func (self *MouseEventDispatcher) AddMouseWheelDownHandler(f func(Mouse)) {
	self.mouseWheelDownHandlers = append(self.mouseWheelDownHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMouseWheelDownHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MouseWheelDown(me Mouse) {
	for i := 0; i < len(self.mouseWheelDownHandlers); i++ {
		self.mouseWheelDownHandlers[i](me)
	}
}

// -- ////////////

func (self *MouseEventDispatcher) AddMouseButtonReleaseHandler(f func(Mouse)) {
	self.mouseButtonReleaseHandlers = append(self.mouseButtonReleaseHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMouseButtonReleaseHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MouseButtonRelease(me Mouse) {
	for i := 0; i < len(self.mouseButtonReleaseHandlers); i++ {
		self.mouseButtonReleaseHandlers[i](me)
	}
}

// -- ////////////

func (self *MouseEventDispatcher) AddMouseButtonPressHandler(f func(Mouse)) {
	self.mouseButtonPressHandlers = append(self.mouseButtonPressHandlers, f)
}

func (self *MouseEventDispatcher) RemoveMouseButtonPressHandler(f func(Mouse)) {
	// TODO - implement
}

func (self *MouseEventDispatcher) MouseButtonPress(me Mouse) {
	for i := 0; i < len(self.mouseButtonPressHandlers); i++ {
		self.mouseButtonPressHandlers[i](me)
	}
}
