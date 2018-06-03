// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	"github.com/sesteel/go-view/event"
)

type CompositeView struct {
	DefaultView
	layout Layout
	event.EventDispatcher
}

func (self *CompositeView) Parent() View {
	return self.parent
}

func (self *CompositeView) SetLayout(layout Layout) {
	self.layout = layout
}

func (self *CompositeView) Layout() Layout {
	return self.layout
}

func (self *CompositeView) Draw(surface *Surface) {
	self.layout.Draw(surface)
}

func (self *CompositeView) MouseEnter(me event.Mouse) {
	self.MouseEventDispatcher.MouseEnter(me)
	self.layout.MouseEnter(me)
}

func (self *CompositeView) MouseExit(me event.Mouse) {
	self.MouseEventDispatcher.MouseExit(me)
	self.layout.MouseExit(me)
}

func (self *CompositeView) MousePosition(me event.Mouse) {
	self.MouseEventDispatcher.MousePosition(me)
	self.layout.MousePosition(me)
}

func (self *CompositeView) MouseWheelUp(me event.Mouse) {
	self.MouseEventDispatcher.MouseWheelUp(me)
	self.layout.MouseWheelUp(me)
}

func (self *CompositeView) MouseWheelDown(me event.Mouse) {
	self.MouseEventDispatcher.MouseWheelDown(me)
	self.layout.MouseWheelDown(me)
}

func (self *CompositeView) MouseButtonPress(me event.Mouse) {
	self.MouseEventDispatcher.MouseButtonPress(me)
	self.layout.MouseButtonPress(me)
}

func (self *CompositeView) MouseButtonRelease(me event.Mouse) {
	self.MouseEventDispatcher.MouseButtonRelease(me)
	self.layout.MouseButtonRelease(me)
}
