// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package editor

import (
	"view"
	"view/color"
	"view/tokenizer"
	"fmt"
)	

type Editor struct {
	view.DefaultComponent
	model Model
}

func New(parent view.View, name string, model Model) *Editor {
	e := &Editor{*view.NewComponent(parent, name), model}
	e.Style().SetPadding(5)
//	e.Style().SetForeground(color.Gray13)
	e.Style().SetBackground(color.Gray2)
	e.Style().SetFontName("Consolas")
	e.Style().SetFontSize(15)
	e.SetParent(parent)
	e.SetName(name)
	return e
}

func (self *Editor) Draw(s *view.Surface) {
	s.DrawFilledBackground(self.Style())
	
//	start position
	var b view.Bounds
	b.Y = 15
	b.X = 10
	for _, l := range self.model.Lines() {
		for _, t := range l {
			if t.Type != tokenizer.NEWLINE && t.Type != tokenizer.SPACE && t.Type != tokenizer.TAB {
				b.Y += 15		
				fmt.Println(">>>>>>", t.Type, t.Value)
				s.DrawTextToken(t, b, self.Style())
			}
		}
	}
}