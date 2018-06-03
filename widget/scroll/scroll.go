// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package scroll

import (
	"github.com/sesteel/go-view"
	"github.com/sesteel/go-view/color"
	"github.com/sesteel/go-view/event"
)

type Scroll interface {
	view.View
	SetOffset(float64)
	Offset() float64
	SetScope(float64)
	Scope() float64
	SetSize(float64)
	Size() float64
	Increment()
	Decrement()
}

type ScrollStyle struct {
	TrackColor   color.RGBA
	HandleColor  color.RGBA
	HandleRadius float64
}

func NewScrollStyle() *ScrollStyle {
	return &ScrollStyle{
		color.ScrollTrack,
		color.ScrollHandle,
		1,
	}
}

type scroll struct {
	view.DefaultView
	event.EventDispatcher
	Style  *ScrollStyle
	offset float64
	scope  float64
	size   float64
}

func newScroll(parent view.View, name string) scroll {
	var s scroll
	s.DefaultView = view.NewDefaultView(parent, name)
	s.Style = NewScrollStyle()
	return s
}

// SetOffset sets the offset but does not enforce  0 > offset > size.
func (self *scroll) SetOffset(offset float64) {
	self.offset = offset
}

func (self *scroll) Offset() float64 {
	return self.offset
}

// Increment increases the offset by 1 unless the scope extends beyond the set size.
func (self *scroll) Increment() {
	self.offset++
}

// Decrement reduces the offset by 1 unless the offset is 0.
func (self *scroll) Decrement() {
	if self.offset > 0 {
		self.offset--
	}
}

// SetSize sets the number of elements from which the visibility ratio is derived from.
func (self *scroll) SetSize(size float64) {
	self.size = size
}

// Size returns the number of elements from which the visibility ratio is derived from.
func (self *scroll) Size() float64 {
	return self.size
}

// SetScope sets the number of elements that can be viewed by the user.
func (self *scroll) SetScope(scope float64) {
	self.scope = scope
}

// Scope returns the number of elements that can be viewed by the user.
func (self *scroll) Scope() float64 {
	return self.scope
}
