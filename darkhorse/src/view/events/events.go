package events

import ()

type EventHandler interface {
	getMouseUpChan()
}

type ChanEventHandler struct {
	MouseButtonUp        chan MouseButtonUp
	MouseLeftButtonUp    chan MouseLeftButtonUp
	MouseRightButtonUp   chan MouseRightButtonUp
	MouseButtonDown      chan MouseButtonDown
	MouseLeftButtonDown  chan MouseLeftButtonDown
	MouseRightButtonDown chan MouseRightButtonDown
	MousePosition        chan MousePosition
	KeyDown              chan KeyDown
	KeyUp                chan KeyUp
}

type KeyUp struct {
	Key int
}

type KeyDown struct {KeyUp}

type MousePosition struct {
	X, Y   int
}

type MouseButtonUp struct {
	MousePosition
	Button int
}

type MouseLeftButtonUp    struct{ MouseButtonUp }
type MouseRightButtonUp   struct{ MouseButtonUp }
type MouseButtonDown      struct{ MouseButtonUp }
type MouseLeftButtonDown  struct{ MouseButtonUp }
type MouseRightButtonDown struct{ MouseButtonUp }
