// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package color

import ()

var (
	Selection RGBA

	White       RGBA
	Gray1       RGBA
	Gray2       RGBA
	Gray3       RGBA
	Gray4       RGBA
	Gray5       RGBA
	Gray6       RGBA
	Gray7       RGBA
	Gray8       RGBA
	Gray9       RGBA
	Gray10      RGBA
	Gray11      RGBA
	Gray12      RGBA
	Gray13      RGBA
	Black       RGBA
	Blue1       RGBA
	Blue2       RGBA
	Blue3       RGBA
	Blue4       RGBA
	Blue5       RGBA
	Cyan1       RGBA
	Cyan2       RGBA
	Teal1       RGBA
	Teal2       RGBA
	Green1      RGBA
	Green2      RGBA
	Yellow1     RGBA
	Yellow2     RGBA
	Orange1     RGBA
	Orange2     RGBA
	Red1        RGBA
	Red2        RGBA
	Purple1     RGBA
	Purple2     RGBA
	Pink1       RGBA
	Pink2       RGBA
	Transparent RGBA

	WindowBegin      RGBA
	WindowEnd        RGBA
	WidgetBorder     RGBA
	WidgetBackground RGBA
	WidgetForeground RGBA
	ProgressBar      RGBA
	Check            RGBA
)

func init() {
	Selection = HexRGBA(0xFF0000FF)
	White = HexRGBA(0xFFFFFFFF)
	Gray1 = HexRGBA(0xF5F7FAFF)
	Gray2 = HexRGBA(0xE4E6EAFF)
	Gray3 = HexRGBA(0xD2D5D9FF)
	Gray4 = HexRGBA(0xC1C4C8FF)
	Gray5 = HexRGBA(0xAFB3B8FF)
	Gray6 = HexRGBA(0x9EA2A7FF)
	Gray7 = HexRGBA(0x8C9197FF)
	Gray8 = HexRGBA(0x7B7F87FF)
	Gray9 = HexRGBA(0x696E76FF)
	Gray10 = HexRGBA(0x585D65FF)
	Gray11 = HexRGBA(0x464C55FF)
	Gray12 = HexRGBA(0x353B45FF)
	Gray13 = HexRGBA(0x232A34FF)
	Black = HexRGBA(0x050505FF)
	Blue1 = HexRGBA(0xCDDBECFF)
	Blue2 = HexRGBA(0xA3C3ECFF)
	Blue3 = HexRGBA(0x5D9CECFF)
	Blue4 = HexRGBA(0x4A89DCFF)
	Blue5 = HexRGBA(0x2B486CFF)
	Cyan1 = HexRGBA(0x4FC1E9FF)
	Cyan2 = HexRGBA(0x3BAFDAFF)
	Teal1 = HexRGBA(0x48CFADFF)
	Teal2 = HexRGBA(0x37BC9BFF)
	Green1 = HexRGBA(0xA0D468FF)
	Green2 = HexRGBA(0x8CC152FF)
	Yellow1 = HexRGBA(0xFFCE54FF)
	Yellow2 = HexRGBA(0xF6BB42FF)
	Orange1 = HexRGBA(0xFC6E51FF)
	Orange2 = HexRGBA(0xE9573FFF)
	Red1 = HexRGBA(0xED5565FF)
	Red2 = HexRGBA(0xDA4453FF)
	Purple1 = HexRGBA(0xAC92ECFF)
	Purple2 = HexRGBA(0x967ADCFF)
	Pink1 = HexRGBA(0xEC87C0FF)
	Pink2 = HexRGBA(0xD770ADFF)

	WindowBegin = HexRGBA(0xD770ADFF)
	WindowEnd = HexRGBA(0xD770ADFF)
	WidgetBorder = Gray7
	WidgetBackground = Gray5
	WidgetForeground = Gray11
	ProgressBar = HexRGBA(0xD770ADFF)
	Check = HexRGBA(0xD770ADFF)

}

type RGBA struct {
	R, G, B, A float64
}

func clritof(c uint8) float64 {
	if c == 0 {
		return 0
	} else {
		return float64(c) / 255
	}
}

func UintRGBA(r, g, b, a uint8) RGBA {
	return RGBA{clritof(r), clritof(g), clritof(b), clritof(a)}
}

func HexRGBA(hex uint32) RGBA {
	return RGBA{clritof(uint8(hex >> 24 & 0xFF)),
		clritof(uint8(hex >> 16 & 0xFF)),
		clritof(uint8(hex >> 8 & 0xFF)),
		clritof(uint8(hex & 0xFF))}
}

func (self RGBA) Shade(pct float64) RGBA {
	if pct < -1 {
		return RGBA{0, 0, 0, self.A}
	} else if pct > 1 {
		return RGBA{1, 1, 1, self.A}
	} else {
		a := 1.0
		b := pct
		if pct < 0.0 {
			a = 0.0
			b = pct * -1.0
		}
		self.R = ((a - self.R) * b) + self.R
		self.G = ((a - self.G) * b) + self.G
		self.B = ((a - self.B) * b) + self.B
	}
	return self
}

func (self RGBA) Alpha(alpha float64) RGBA {
	self.A = alpha
	return self
}
