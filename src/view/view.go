// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

// +build linux,!goci
package view

import (
	"view/event"
)

type Drawer interface {
	// Draw traverses the view heirarchy drawing dirty views.
	Draw(*Surface)

	// Redraw marks the dirty path up the view heirarchy.
	Redraw()
}

type View interface {
	Drawer
	event.FocusNotifier
	event.FocusHandler
	event.MouseNotifier
	event.MouseHandler
	SetParent(View)
	Parent() View
	Name() string
	SetStyle(Style)
	Style() Style
}

type DefaultView struct {
	parent  View
	name    string
	style   Style
	current bool
}

func (self *DefaultView) SetParent(parent View) {
	self.parent = parent
}

func (self *DefaultView) Parent() View {
	return self.parent
}

func (self *DefaultView) SetName(name string) {
	self.name = name
}

func (self *DefaultView) Name() string {
	return self.name
}

func (self *DefaultView) SetStyle(style Style) {
	self.style = style
}

func (self *DefaultView) Style() Style {
	return self.style
}

func (self *DefaultView) Draw(surface *Surface) {
	// default drawing goes here
}

func (self *DefaultView) Redraw() {
	if self.parent != nil {
		self.parent.Redraw()
	}
}
