// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

// +build linux,!goci
package view

import (
	"fmt"
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
	Surface() *Surface
	Name() string
	SetStyle(Style)
	Style() Style
}

type DefaultView struct {
	parent  View
	surface *Surface
	name    string
	width   float64
	height  float64
	x, y    float64
	focus   bool
	style   Style
	current bool
}

func (self *DefaultView) SetParent(parent View) {
	self.parent = parent
}

func (self *DefaultView) Parent() View {
	return self.parent
}

func (self *DefaultView) Position() (float64, float64) {
	return self.x, self.y
}

func (self *DefaultView) Surface() *Surface {
	return self.surface
}

func (self *DefaultView) SetName(name string) {
	self.name = name
}

func (self *DefaultView) Name() string {
	return self.name
}

func (self *DefaultView) SetSize(width, height float64) {
	self.width = width
	self.height = height
}

func (self *DefaultView) Size() (float64, float64) {
	return self.width, self.height
}

func (self *DefaultView) SetStyle(style Style) {
	self.style = style
}

func (self *DefaultView) Style() Style {
	return self.style
}

func (self *DefaultView) Draw(surface *Surface) {
	// default drawing does here
}

func (self *DefaultView) Redraw() {
	if DEBUG {
		fmt.Println("View.Redraw()")
	}

	if self.parent != nil {
		self.parent.Redraw()
	}
}
