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
}

func (self *scrollMap) draw(a, b *view.Surface) {
	a.Save()
	a.SetSourceSurface(b, 0, 0)
	a.Paint()
	a.SetSourceRGBA(self.BoundaryLineColor)
	a.MoveTo(self.Width+1, 0)
	a.SetLineWidth(self.BoundaryLineWidth)
	a.LineTo(self.Width+1, float64(a.Height()))
	a.Stroke()
	a.Restore()
}
