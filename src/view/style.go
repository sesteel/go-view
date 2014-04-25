// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	"view/theme"
)

type OverflowXStrategy int

const (
	STYLE_OVERFLOW_X_NONE OverflowXStrategy = iota
	STYLE_OVERFLOW_X_SCROLL
	STYLE_OVERFLOW_X_WRAP

//	STYLE_OVERFLOW_X_SHADE
//	STYLE_OVERFLOW_X_FADE
)

type OverflowYStrategy int

const (
	STYLE_OVERFLOW_Y_NONE OverflowYStrategy = iota
	STYLE_OVERFLOW_Y_SCROLL

//	STYLE_OVERFLOW_Y_SHADE
//	STYLE_OVERFLOW_Y_FADE
)

var defaultFontOptions *FontOptions

func init() {
	defaultFontOptions = NewFontOptions()
	defaultFontOptions.SetAntialias(ANTIALIAS_SUBPIXEL)
	defaultFontOptions.SetHintStyle(HINT_STYLE_FULL)
	defaultFontOptions.SetHintMetric(HINT_METRICS_ON)
}

type Style interface {
	SetAntialias(bool)
	SetFontName(string)
	SetFontWeight(int)
	SetFontSlant(int)
	SetFontSize(float64)
	SetTabWidth(int)
	SetBackground(theme.RGBA)
	SetForeground(theme.RGBA)
	SetBorderColor(theme.RGBA)
	SetBorderColorTop(theme.RGBA)
	SetBorderColorBottom(theme.RGBA)
	SetBorderColorLeft(theme.RGBA)
	SetBorderColorRight(theme.RGBA)
	SetBorderWidth(float64)
	SetBorderWidthTop(float64)
	SetBorderWidthBottom(float64)
	SetBorderWidthLeft(float64)
	SetBorderWidthRight(float64)
	SetPadding(float64)
	SetPaddingTop(float64)
	SetPaddingBottom(float64)
	SetPaddingLeft(float64)
	SetPaddingRight(float64)
	SetOverflowX(OverflowXStrategy)
	SetOverflowY(OverflowYStrategy)

	Antialias() bool
	FontName() string
	FontWeight() int
	FontSlant() int
	FontSize() float64
	TabWidth() int
	Background() theme.RGBA
	Foreground() theme.RGBA
	BorderColorTop() theme.RGBA
	BorderColorBottom() theme.RGBA
	BorderColorLeft() theme.RGBA
	BorderColorRight() theme.RGBA
	BorderWidthTop() float64
	BorderWidthBottom() float64
	BorderWidthLeft() float64
	BorderWidthRight() float64
	PaddingTop() float64
	PaddingBottom() float64
	PaddingLeft() float64
	PaddingRight() float64
	OverflowX() OverflowXStrategy
	OverflowY() OverflowYStrategy
}

type defaultStyle struct {
	antialias         bool
	fontName          string
	fontWeight        int
	fontSlant         int
	fontSize          float64
	tabWidth          int
	backgroundColor   theme.RGBA
	foregroundColor   theme.RGBA
	borderColorTop    theme.RGBA
	borderColorBottom theme.RGBA
	borderColorLeft   theme.RGBA
	borderColorRight  theme.RGBA
	borderWidthTop    float64
	borderWidthBottom float64
	borderWidthLeft   float64
	borderWidthRight  float64
	paddingTop        float64
	paddingBottom     float64
	paddingLeft       float64
	paddingRight      float64
	overflowX         OverflowXStrategy
	overflowY         OverflowYStrategy
}

func NewStyle() Style {
	s := new(defaultStyle)
	s.SetBorderWidth(1)
	s.SetAntialias(true)
	s.SetFontName("Clear Sans")
	s.SetFontSize(16)
	s.SetTabWidth(4)
	s.SetFontSlant(FONT_SLANT_NORMAL)
	s.SetFontWeight(FONT_WEIGHT_NORMAL)
	s.SetBackground(theme.Gray3)
	s.SetForeground(theme.Gray7)
	s.SetPadding(3)
	return s
}

func (self *defaultStyle) SetAntialias(a bool)                   { self.antialias = a }
func (self *defaultStyle) SetFontName(name string)               { self.fontName = name }
func (self *defaultStyle) SetFontWeight(weight int)              { self.fontWeight = weight }
func (self *defaultStyle) SetFontSlant(slant int)                { self.fontSlant = slant }
func (self *defaultStyle) SetFontSize(size float64)              { self.fontSize = size }
func (self *defaultStyle) SetTabWidth(width int)                 { self.tabWidth = width }
func (self *defaultStyle) SetBackground(color theme.RGBA)        { self.backgroundColor = color }
func (self *defaultStyle) SetForeground(color theme.RGBA)        { self.foregroundColor = color }
func (self *defaultStyle) SetBorderColorTop(color theme.RGBA)    { self.borderColorTop = color }
func (self *defaultStyle) SetBorderColorBottom(color theme.RGBA) { self.borderColorBottom = color }
func (self *defaultStyle) SetBorderColorLeft(color theme.RGBA)   { self.borderColorLeft = color }
func (self *defaultStyle) SetBorderColorRight(color theme.RGBA)  { self.borderColorRight = color }
func (self *defaultStyle) SetBorderWidthTop(width float64)       { self.borderWidthTop = width }
func (self *defaultStyle) SetBorderWidthBottom(width float64)    { self.borderWidthBottom = width }
func (self *defaultStyle) SetBorderWidthLeft(width float64)      { self.borderWidthLeft = width }
func (self *defaultStyle) SetBorderWidthRight(width float64)     { self.borderWidthRight = width }
func (self *defaultStyle) SetPaddingTop(padding float64)         { self.paddingTop = padding }
func (self *defaultStyle) SetPaddingBottom(padding float64)      { self.paddingBottom = padding }
func (self *defaultStyle) SetPaddingLeft(padding float64)        { self.paddingLeft = padding }
func (self *defaultStyle) SetPaddingRight(padding float64)       { self.paddingRight = padding }

func (self *defaultStyle) SetBorderColor(color theme.RGBA) {
	self.borderColorTop = color
	self.borderColorBottom = color
	self.borderColorLeft = color
	self.borderColorRight = color
}

func (self *defaultStyle) SetBorderWidth(width float64) {
	self.borderWidthTop = width
	self.borderWidthBottom = width
	self.borderWidthLeft = width
	self.borderWidthRight = width
}

func (self *defaultStyle) SetPadding(padding float64) {
	self.paddingTop = padding
	self.paddingBottom = padding
	self.paddingLeft = padding
	self.paddingRight = padding
}

func (self *defaultStyle) SetOverflowX(overflowX OverflowXStrategy) {
	self.overflowX = overflowX
}

func (self *defaultStyle) SetOverflowY(overflowY OverflowYStrategy) {
	self.overflowY = overflowY
}

func (self *defaultStyle) Antialias() bool               { return self.antialias }
func (self *defaultStyle) FontName() string              { return self.fontName }
func (self *defaultStyle) FontWeight() int               { return self.fontWeight }
func (self *defaultStyle) FontSlant() int                { return self.fontSlant }
func (self *defaultStyle) FontSize() float64             { return self.fontSize }
func (self *defaultStyle) TabWidth() int                 { return self.tabWidth }
func (self *defaultStyle) Background() theme.RGBA        { return self.backgroundColor }
func (self *defaultStyle) Foreground() theme.RGBA        { return self.foregroundColor }
func (self *defaultStyle) BorderColorTop() theme.RGBA    { return self.borderColorTop }
func (self *defaultStyle) BorderColorBottom() theme.RGBA { return self.borderColorBottom }
func (self *defaultStyle) BorderColorLeft() theme.RGBA   { return self.borderColorLeft }
func (self *defaultStyle) BorderColorRight() theme.RGBA  { return self.borderColorRight }
func (self *defaultStyle) BorderWidthTop() float64       { return self.borderWidthTop }
func (self *defaultStyle) BorderWidthBottom() float64    { return self.borderWidthBottom }
func (self *defaultStyle) BorderWidthLeft() float64      { return self.borderWidthLeft }
func (self *defaultStyle) BorderWidthRight() float64     { return self.borderWidthRight }
func (self *defaultStyle) PaddingTop() float64           { return self.paddingTop }
func (self *defaultStyle) PaddingBottom() float64        { return self.paddingBottom }
func (self *defaultStyle) PaddingLeft() float64          { return self.paddingLeft }
func (self *defaultStyle) PaddingRight() float64         { return self.paddingRight }
func (self *defaultStyle) OverflowX() OverflowXStrategy  { return self.overflowX }
func (self *defaultStyle) OverflowY() OverflowYStrategy  { return self.overflowY }
