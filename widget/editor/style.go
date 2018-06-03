// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import (
	"github.com/sesteel/go-view"
	"github.com/sesteel/go-view/color"
)

type EditorStyle struct {
	Background      color.RGBA
	Foreground      color.RGBA
	StringStyle     TokenStyle
	PrimitiveStyle  TokenStyle
	KeywordStyle    TokenStyle
	CommentStyle    TokenStyle
	WhitespaceColor color.RGBA
	Paddings        view.Paddings
	Font            view.Font
	LineSpace       float64
	TabWidth        int
	MarginColumn    int
	MarginColor     color.RGBA
}

func NewEditorStyle() *EditorStyle {
	return &EditorStyle{
		color.TextBackground,
		color.TextForeground,
		TokenStyle{view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_NORMAL, color.Black},
		TokenStyle{view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_ITALIC, color.Black},
		TokenStyle{view.FONT_WEIGHT_BOLD, view.FONT_SLANT_NORMAL, color.Black},
		TokenStyle{view.FONT_WEIGHT_NORMAL, view.FONT_SLANT_ITALIC, color.Gray8},
		color.RGBA{color.Cyan1.R, color.Cyan1.G, color.Cyan1.B, 0.5},
		view.Paddings{0, 0, 0, 0},
		view.Font{
			"Monospace",
			view.FONT_WEIGHT_NORMAL,
			view.FONT_SLANT_NORMAL,
			14,
		},
		1,
		4,
		80,
		color.RGBA{color.Pink1.R, color.Pink1.G, color.Pink1.B, 0.25},
	}
}
