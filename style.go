// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	"github.com/sesteel/go-view/color"
)

var defaultFontOptions *FontOptions

func init() {
	defaultFontOptions = NewFontOptions()
	defaultFontOptions.SetAntialias(ANTIALIAS_SUBPIXEL)
	defaultFontOptions.SetHintStyle(HINT_STYLE_FULL)
	defaultFontOptions.SetHintMetric(HINT_METRICS_ON)
}

type Style interface {
	SetFontName(string)
	SetFontWeight(int)
	SetFontSlant(int)
	SetFontSize(float64)
	SetTabWidth(int)
	SetBackground(color.RGBA)
	SetForeground(color.RGBA)
	SetBorderColor(color.RGBA)
	SetBorderColorTop(color.RGBA)
	SetBorderColorBottom(color.RGBA)
	SetBorderColorLeft(color.RGBA)
	SetBorderColorRight(color.RGBA)
	SetBorderWidth(float64)
	SetBorderWidthTop(float64)
	SetBorderWidthBottom(float64)
	SetBorderWidthLeft(float64)
	SetBorderWidthRight(float64)
	SetRadius(float64)
	SetRadiusTopLeft(float64)
	SetRadiusTopRight(float64)
	SetRadiusBottomLeft(float64)
	SetRadiusBottomRight(float64)
	SetPadding(float64)
	SetPaddingTop(float64)
	SetPaddingBottom(float64)
	SetPaddingLeft(float64)
	SetPaddingRight(float64)
	// SetOverflowX(OverflowXStrategy)
	// SetOverflowY(OverflowYStrategy)
	SetTextAlignment(TextAlignment)

	FontName() string
	FontWeight() int
	FontSlant() int
	FontSize() float64
	TabWidth() int
	Background() color.RGBA
	Foreground() color.RGBA
	BorderColorTop() color.RGBA
	BorderColorBottom() color.RGBA
	BorderColorLeft() color.RGBA
	BorderColorRight() color.RGBA
	BorderWidthTop() float64
	BorderWidthBottom() float64
	BorderWidthLeft() float64
	BorderWidthRight() float64
	RadiusTopLeft() float64
	RadiusTopRight() float64
	RadiusBottomLeft() float64
	RadiusBottomRight() float64
	Padding() (l, r, t, b float64)
	PaddingTop() float64
	PaddingBottom() float64
	PaddingLeft() float64
	PaddingRight() float64
	// OverflowX() OverflowXStrategy
	// OverflowY() OverflowYStrategy
	TextAlignment() TextAlignment
}

type DefaultStyle struct {
	antialias         bool
	fontName          string
	fontWeight        int
	fontSlant         int
	fontSize          float64
	tabWidth          int
	backgroundColor   color.RGBA
	foregroundColor   color.RGBA
	borderColorTop    color.RGBA
	borderColorBottom color.RGBA
	borderColorLeft   color.RGBA
	borderColorRight  color.RGBA
	borderWidthTop    float64
	borderWidthBottom float64
	borderWidthLeft   float64
	borderWidthRight  float64
	radiusTL          float64
	radiusTR          float64
	radiusBR          float64
	radiusBL          float64
	paddingTop        float64
	paddingBottom     float64
	paddingLeft       float64
	paddingRight      float64
	// overflowX         OverflowXStrategy
	// overflowY         OverflowYStrategy
	textAlignment TextAlignment
}

func NewStyle() Style {
	s := new(DefaultStyle)
	s.SetBorderWidth(1)
	s.SetFontName("Liberation Sans")
	s.SetFontSize(20)
	s.SetFontSlant(FONT_SLANT_NORMAL)
	s.SetFontWeight(FONT_WEIGHT_NORMAL)
	s.SetBackground(color.Gray5)
	s.SetForeground(color.Gray11)
	s.SetPadding(3)
	s.SetBorderColor(color.Gray7)
	s.SetBorderWidth(2)
	s.SetRadius(2)
	s.SetTextAlignment(STYLE_TEXT_CENTERED)
	return s
}

func NewDisabledStyle() Style {
	s := NewStyle()
	s.SetForeground(color.Gray9)
	return s
}

func CloneAsDisabledStyle(style Style) Style {
	s := CloneStyle(style)
	s.SetForeground(color.Gray9)
	return s
}

func CloneStyle(style Style) Style {
	s := new(DefaultStyle)
	s.SetBorderWidthBottom(style.BorderWidthBottom())
	s.SetBorderWidthTop(style.BorderWidthTop())
	s.SetBorderWidthLeft(style.BorderWidthLeft())
	s.SetBorderWidthRight(style.BorderWidthRight())
	s.SetFontName(style.FontName())
	s.SetFontSize(style.FontSize())
	s.SetFontSlant(style.FontSlant())
	s.SetFontWeight(style.FontWeight())
	s.SetTabWidth(style.TabWidth())
	s.SetBackground(style.Background())
	s.SetForeground(style.Foreground())
	s.SetPaddingBottom(style.PaddingBottom())
	s.SetPaddingTop(style.PaddingTop())
	s.SetPaddingLeft(style.PaddingLeft())
	s.SetPaddingRight(style.PaddingRight())
	s.SetBorderColorBottom(style.BorderColorBottom())
	s.SetBorderColorTop(style.BorderColorTop())
	s.SetBorderColorLeft(style.BorderColorLeft())
	s.SetBorderColorRight(style.BorderColorRight())
	s.SetBorderWidthBottom(style.BorderWidthBottom())
	s.SetBorderWidthTop(style.BorderWidthTop())
	s.SetBorderWidthLeft(style.BorderWidthLeft())
	s.SetBorderWidthRight(style.BorderWidthRight())
	s.SetTextAlignment(style.TextAlignment())
	s.SetRadiusTopLeft(style.RadiusTopLeft())
	s.SetRadiusTopRight(style.RadiusTopRight())
	s.SetRadiusBottomLeft(style.RadiusBottomLeft())
	s.SetRadiusBottomRight(style.RadiusBottomRight())
	return s
}

func (self *DefaultStyle) SetFontName(name string)               { self.fontName = name }
func (self *DefaultStyle) SetFontWeight(weight int)              { self.fontWeight = weight }
func (self *DefaultStyle) SetFontSlant(slant int)                { self.fontSlant = slant }
func (self *DefaultStyle) SetFontSize(size float64)              { self.fontSize = size }
func (self *DefaultStyle) SetTabWidth(width int)                 { self.tabWidth = width }
func (self *DefaultStyle) SetBackground(color color.RGBA)        { self.backgroundColor = color }
func (self *DefaultStyle) SetForeground(color color.RGBA)        { self.foregroundColor = color }
func (self *DefaultStyle) SetBorderColorTop(color color.RGBA)    { self.borderColorTop = color }
func (self *DefaultStyle) SetBorderColorBottom(color color.RGBA) { self.borderColorBottom = color }
func (self *DefaultStyle) SetBorderColorLeft(color color.RGBA)   { self.borderColorLeft = color }
func (self *DefaultStyle) SetBorderColorRight(color color.RGBA)  { self.borderColorRight = color }
func (self *DefaultStyle) SetBorderWidthTop(width float64)       { self.borderWidthTop = width }
func (self *DefaultStyle) SetBorderWidthBottom(width float64)    { self.borderWidthBottom = width }
func (self *DefaultStyle) SetBorderWidthLeft(width float64)      { self.borderWidthLeft = width }
func (self *DefaultStyle) SetBorderWidthRight(width float64)     { self.borderWidthRight = width }
func (self *DefaultStyle) SetPaddingTop(padding float64)         { self.paddingTop = padding }
func (self *DefaultStyle) SetPaddingBottom(padding float64)      { self.paddingBottom = padding }
func (self *DefaultStyle) SetPaddingLeft(padding float64)        { self.paddingLeft = padding }
func (self *DefaultStyle) SetPaddingRight(padding float64)       { self.paddingRight = padding }
func (self *DefaultStyle) SetRadiusTopLeft(radius float64)       { self.radiusTL = radius }
func (self *DefaultStyle) SetRadiusTopRight(radius float64)      { self.radiusTR = radius }
func (self *DefaultStyle) SetRadiusBottomLeft(radius float64)    { self.radiusBL = radius }
func (self *DefaultStyle) SetRadiusBottomRight(radius float64)   { self.radiusBR = radius }

func (self *DefaultStyle) SetRadius(radius float64) {
	self.radiusTL = radius
	self.radiusTR = radius
	self.radiusBL = radius
	self.radiusBR = radius
}

func (self *DefaultStyle) SetTextAlignment(alignment TextAlignment) {
	self.textAlignment = alignment
}

func (self *DefaultStyle) SetBorderColor(color color.RGBA) {
	self.borderColorTop = color
	self.borderColorBottom = color
	self.borderColorLeft = color
	self.borderColorRight = color
}

func (self *DefaultStyle) SetBorderWidth(width float64) {
	self.borderWidthTop = width
	self.borderWidthBottom = width
	self.borderWidthLeft = width
	self.borderWidthRight = width
}

func (self *DefaultStyle) SetPadding(padding float64) {
	self.paddingTop = padding
	self.paddingBottom = padding
	self.paddingLeft = padding
	self.paddingRight = padding
}

// func (self *DefaultStyle) SetOverflowX(overflowX OverflowXStrategy) {
// 	self.overflowX = overflowX
// }

// func (self *DefaultStyle) SetOverflowY(overflowY OverflowYStrategy) {
// 	self.overflowY = overflowY
// }

func (self *DefaultStyle) FontName() string              { return self.fontName }
func (self *DefaultStyle) FontWeight() int               { return self.fontWeight }
func (self *DefaultStyle) FontSlant() int                { return self.fontSlant }
func (self *DefaultStyle) FontSize() float64             { return self.fontSize }
func (self *DefaultStyle) TabWidth() int                 { return self.tabWidth }
func (self *DefaultStyle) Background() color.RGBA        { return self.backgroundColor }
func (self *DefaultStyle) Foreground() color.RGBA        { return self.foregroundColor }
func (self *DefaultStyle) BorderColorTop() color.RGBA    { return self.borderColorTop }
func (self *DefaultStyle) BorderColorBottom() color.RGBA { return self.borderColorBottom }
func (self *DefaultStyle) BorderColorLeft() color.RGBA   { return self.borderColorLeft }
func (self *DefaultStyle) BorderColorRight() color.RGBA  { return self.borderColorRight }
func (self *DefaultStyle) BorderWidthTop() float64       { return self.borderWidthTop }
func (self *DefaultStyle) BorderWidthBottom() float64    { return self.borderWidthBottom }
func (self *DefaultStyle) BorderWidthLeft() float64      { return self.borderWidthLeft }
func (self *DefaultStyle) BorderWidthRight() float64     { return self.borderWidthRight }

func (self *DefaultStyle) Padding() (l, r, t, b float64) {
	return self.paddingLeft, self.paddingRight, self.paddingTop, self.paddingBottom
}

func (self *DefaultStyle) PaddingTop() float64    { return self.paddingTop }
func (self *DefaultStyle) PaddingBottom() float64 { return self.paddingBottom }
func (self *DefaultStyle) PaddingLeft() float64   { return self.paddingLeft }
func (self *DefaultStyle) PaddingRight() float64  { return self.paddingRight }

// func (self *DefaultStyle) OverflowX() OverflowXStrategy { return self.overflowX }
// func (self *DefaultStyle) OverflowY() OverflowYStrategy { return self.overflowY }
func (self *DefaultStyle) RadiusTopLeft() float64     { return self.radiusTL }
func (self *DefaultStyle) RadiusTopRight() float64    { return self.radiusTR }
func (self *DefaultStyle) RadiusBottomLeft() float64  { return self.radiusBL }
func (self *DefaultStyle) RadiusBottomRight() float64 { return self.radiusBR }

func (self *DefaultStyle) TextAlignment() TextAlignment {
	return self.textAlignment
}
