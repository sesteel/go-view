// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package event

import (
	"github.com/sesteel/go-view/common"
)

type MouseButton int8

const (
	MOUSE_BUTTON_NONE MouseButton = iota
	MOUSE_BUTTON_LEFT
	MOUSE_BUTTON_MIDDLE
	MOUSE_BUTTON_RIGHT
)

var mouse MouseState

func LastMouseState() MouseState {
	return mouse
}

type MouseState struct {
	LeftPressed   bool
	MiddlePressed bool
	RightPressed  bool
	X, Y          float64
}

type Mouse struct {
	Button MouseButton
	MouseState
}

func (self *Mouse) Normalize(offset common.Point) Mouse {
	return Mouse{
		self.Button,
		MouseState{
			self.LeftPressed,
			self.MiddlePressed,
			self.RightPressed,
			self.X - offset.X,
			self.Y - offset.Y},
	}
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
	mouse = me.MouseState
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
	mouse = me.MouseState
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
	mouse = me.MouseState
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
	mouse = me.MouseState
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
	mouse = me.MouseState
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
	mouse = me.MouseState
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
	mouse = me.MouseState
	for i := 0; i < len(self.mouseButtonPressHandlers); i++ {
		self.mouseButtonPressHandlers[i](me)
	}
}
