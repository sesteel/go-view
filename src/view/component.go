// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	"view/color"
	"view/event"
)

type Component interface {
	View
	// Focus() bool
	// SetFocus(bool)
}

type DefaultComponent struct {
	DefaultView
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
	c.style = NewStyle()
	c.name = name
	// TODO implement component again
	return c
}

func (self *DefaultComponent) Draw(surface *Surface) {
	style := self.style
	parent := self.parent

	// traverse up tree to find style until parent is nil
	for self.style == nil && parent != nil {
		style = self.parent.Style()
		parent = parent.Parent()
	}

	// cannot draw without style... draw fusia and black text
	if style == nil {
		msg := "Error: No style set for component [" + self.name + "]'s hierarchy."
		surface.SelectFontFace("Sans", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
		surface.SetFontSize(16)
		te := surface.TextExtents(msg)
		surface.SetSourceRGB(1, 2, 1)
		surface.RoundedRectangle(0, 0, float64(surface.Width()), float64(surface.Height()), 0, 0, 0, 0)
		surface.StrokePreserve()
		surface.Fill()

		surface.SelectFontFace("Sans", FONT_SLANT_NORMAL, FONT_WEIGHT_NORMAL)
		surface.SetFontSize(16)
		surface.SetSourceRGBA(color.Cyan2)
		surface.MoveTo(0, te.Height)
		surface.ShowText(msg)

		return
	}

	surface.SetAntialias(ANTIALIAS_SUBPIXEL)
	surface.SetSourceRGBA(style.Background())
	surface.RoundedRectangle(0, 0, float64(surface.Width()), float64(surface.Height()), 0, 0, 0, 0)
	surface.Fill()
	surface.SetLineWidth(style.BorderWidthTop())
	surface.RoundedRectangle(0, 0, float64(surface.Width()), float64(surface.Height()), 0, 0, 0, 0)
	surface.StrokePreserve()
}
