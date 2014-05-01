package progressbar

import (
	"view"
	"view/color"
)

type ProgressBar struct {
	view.DefaultComponent
	progressStyle view.Style
	text          string
	value         float64
	max           float64
}

func New(parent view.View, name string, max float64) *ProgressBar {
	progressStyle := view.NewStyle()
	progressStyle.SetBackground(color.Green1)
	progressStyle.SetBorderColor(color.Green2)
	bar := &ProgressBar{*view.NewComponent(parent, name), progressStyle, "", 0, max}
	return bar
}

func (self *ProgressBar) SetProgressStyle(style view.Style) {
	self.progressStyle = style
}

func (self *ProgressBar) ProgressStyle() view.Style {
	return self.progressStyle
}

func (self *ProgressBar) SetText(text string) {
	self.text = text
}

func (self *ProgressBar) Text() string {
	return self.text
}

func (self *ProgressBar) SetValue(v float64) {
	self.value = v
}

func (self *ProgressBar) Value() float64 {
	return self.value
}

func (self *ProgressBar) SetMax(m float64) {
	self.max = m
}

func (self *ProgressBar) Max() float64 {
	return self.max
}

func (self *ProgressBar) Draw(s *view.Surface) {
	style := self.Style()
	s.DrawBackgroundStyle(style)
	if self.max != 0 || self.value != 0 {
		
		// TODO fix this patch
		max := (self.max * .01) + self.max
		
		percent := self.value / max
		pl := style.PaddingLeft() + 1
		pr := style.PaddingRight() + 1
		maxw := float64(s.Width()) - pl - pr
		self.progressStyle.SetPaddingRight(maxw - (maxw * percent))
		self.progressStyle.SetPaddingLeft(pl)
		self.progressStyle.SetPaddingTop(style.PaddingTop() + 1)
		self.progressStyle.SetPaddingBottom(style.PaddingBottom() + 1)
		s.DrawBackgroundStyle(self.progressStyle)
	}
}
