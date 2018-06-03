// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package layout

import (
	"github.com/sesteel/go-view"
	"github.com/sesteel/go-view/event"
)

// Fill is a simple layout that expands a single
// child to the size allotted to the layout by
// the target.
type Fill struct {
	target view.View
	child  view.View
}

func NewFill(target view.View) *Fill {
	l := new(Fill)
	l.target = target
	return l
}

func (self *Fill) SetChild(d view.View) {
	self.child = d
}

func (self *Fill) Child() view.View {
	return self.child
}

func (self *Fill) Draw(surface *view.Surface) {
	self.child.Draw(surface)
}

func (self *Fill) Animate(surface *view.Surface) {
	if anim, ok := self.child.(view.Animator); ok {
		anim.Animate(surface)
	}
}

func (self *Fill) Redraw() {
	self.target.Redraw()
}

func (self *Fill) MousePosition(ev event.Mouse) {
	self.child.MousePosition(ev)
}

func (self *Fill) MouseButtonPress(ev event.Mouse) {
	self.child.MouseButtonPress(ev)
}

func (self *Fill) MouseButtonRelease(ev event.Mouse) {
	self.child.MouseButtonRelease(ev)
}

func (self *Fill) MouseEnter(ev event.Mouse) {
	self.child.MouseEnter(ev)
}

func (self *Fill) MouseExit(ev event.Mouse) {
	self.child.MouseExit(ev)
}

func (self *Fill) MouseWheelUp(ev event.Mouse) {
	self.child.MouseWheelUp(ev)
}

func (self *Fill) MouseWheelDown(ev event.Mouse) {
	self.child.MouseWheelDown(ev)
}
