// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	"github.com/sesteel/go-view/color"
	"github.com/sesteel/go-view/event"
)

type Component interface {
	View
}

type DefaultComponent struct {
	DefaultView
	Background  color.RGBA
	BorderColor color.RGBA
	BorderWidth float64
	event.EventDispatcher
}

// NewComponent creates a new DefaultComponent.  DefaultComponent
// is generally not usable from a user perspective; it is useful
// for referencing, compositing or embedding in other components,
// however.
//
// Aside, the name passed into this function reprsents this component
// programtically and should be unique in cases where context may be
// ambiguous.  It is intended to support testing and accessibility
// frameworks.
func NewComponent(parent View, name string) *DefaultComponent {
	c := new(DefaultComponent)
	c.parent = parent
	c.name = name
	c.Background = color.WidgetBackground
	c.BorderColor = color.WidgetBorder
	c.BorderWidth = 1
	return c
	// return &DefaultComponent{
	// 	DefaultView{
	// 		parent,
	// 		name},
	// 	color.WidgetBackground,
	// 	color.WidgetBorder,
	// 	1,
	// }
}

func (self *DefaultComponent) Draw(surface *Surface) {
	parent := self.parent

	for parent != nil {
		parent = parent.Parent()
	}

	surface.SetAntialias(ANTIALIAS_SUBPIXEL)
	surface.SetSourceRGBA(self.Background)
	surface.RoundedRectangle(0.5, 0.5, float64(surface.Width()), float64(surface.Height()), 0, 0, 0, 0)
	surface.Fill()
	surface.SetLineWidth(self.BorderWidth)
	surface.RoundedRectangle(0.5, 0.5, float64(surface.Width()), float64(surface.Height()), 0, 0, 0, 0)
	surface.StrokePreserve()
}
