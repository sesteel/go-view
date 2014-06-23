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
