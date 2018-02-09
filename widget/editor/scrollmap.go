// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"view"
	"view/color"
)

type scrollMap struct {
	Width             float64
	Scale             float64
	OverlayColor      color.RGBA
	BoundaryLineColor color.RGBA
	BoundaryLineWidth float64
	viewStart         float64
	viewStop          float64
}

func (self *scrollMap) draw(a, b *view.Surface) {
	// pixelAlign aligns to a pixel edge space so we can draw a single pixel line
	const pixelAlign = 0.5
	a.Save()

	a.SetSourceSurface(b, 0, 0)
	a.Paint()

	// Right Line Border
	a.SetSourceRGBA(self.BoundaryLineColor)
	a.MoveTo(self.Width+pixelAlign, 0)
	a.SetLineWidth(self.BoundaryLineWidth)
	a.LineTo(self.Width+pixelAlign, float64(a.Height()))
	a.Stroke()

	// top frame border
	top := float64(int(self.viewStart * self.Scale))
	top += pixelAlign
	a.MoveTo(0, top)
	a.LineTo(self.Width, top)
	a.Stroke()

	// bottom frame border
	bottom := float64(int(self.viewStop * self.Scale))
	bottom += pixelAlign
	a.MoveTo(0, bottom)
	a.LineTo(self.Width, bottom)
	a.Stroke()

	a.SetSourceRGBA(color.Gray10.Alpha(.05))
	a.Rectangle(0, 0, self.Width, self.viewStart*self.Scale)
	a.Fill()
	a.Rectangle(0, self.viewStop*self.Scale, self.Width, float64(a.Height()))
	a.Fill()
	a.Restore()
}
