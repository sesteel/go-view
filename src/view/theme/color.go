// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package theme

import ()

var (
	Selection RGBA

	White RGBA
	Gray1 RGBA
	Gray2 RGBA
	Gray3 RGBA
	Gray4 RGBA
	Gray5 RGBA
	Gray6 RGBA
	Gray7 RGBA
	Black RGBA

	Blue1   RGBA
	Blue2   RGBA
	Cyan1   RGBA
	Cyan2   RGBA
	Teal1   RGBA
	Teal2   RGBA
	Green1  RGBA
	Green2  RGBA
	Yellow1 RGBA
	Yellow2 RGBA
	Orange1 RGBA
	Orange2 RGBA
	Red1    RGBA
	Red2    RGBA
	Purple1 RGBA
	Purple2 RGBA
	Pink1   RGBA
	Pink2   RGBA
)

func init() {
	Selection = HexRGBA(0xFF0000FF) 
	White   = HexRGBA(0xFFFFFFFF)
	Gray1   = HexRGBA(0xF5F7FAFF)
	Gray2   = HexRGBA(0xE6E9EDFF)
	Gray3   = HexRGBA(0xCCD1D9FF)
	Gray4   = HexRGBA(0xAAB2BDFF)
	Gray5   = HexRGBA(0x656D78FF)
	Gray6   = HexRGBA(0x434A54FF)
	Gray7   = HexRGBA(0x232A34FF)
	Black   = HexRGBA(0x050505FF)
	Blue1   = HexRGBA(0x5D9CECFF)
	Blue2   = HexRGBA(0x4A89DCFF)
	Cyan1   = HexRGBA(0x4FC1E9FF)
	Cyan2   = HexRGBA(0x3BAFDAFF)
	Teal1   = HexRGBA(0x48CFADFF)
	Teal2   = HexRGBA(0x37BC9BFF)
	Green1  = HexRGBA(0xA0D468FF)
	Green2  = HexRGBA(0x8CC152FF)
	Yellow1 = HexRGBA(0xFFCE54FF)
	Yellow2 = HexRGBA(0xF6BB42FF)
	Orange1 = HexRGBA(0xFC6E51FF)
	Orange2 = HexRGBA(0xE9573FFF)
	Red1    = HexRGBA(0xED5565FF)
	Red2    = HexRGBA(0xDA4453FF)
	Purple1 = HexRGBA(0xAC92ECFF)
	Purple2 = HexRGBA(0x967ADCFF)
	Pink1   = HexRGBA(0xEC87C0FF)
	Pink2   = HexRGBA(0xD770ADFF)
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

// TODO RGBAToHSBA
// TODO RGBAToHSLA
// TODO RGBAToHSVA