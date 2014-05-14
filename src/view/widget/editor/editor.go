// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package editor

import (
	"view"
)	

type Editor struct {
	view.DefaultComponent
	model Model
}

func New(parent view.View, name string, model Model) *Editor {
	e := &Editor{*view.NewComponent(parent, name), model}
	e.Style().SetPadding(0)
	e.SetParent(parent)
	e.SetName(name)
	return e
}

func (self *Editor) Draw(s *view.Surface) {
	s.DrawFilledBackground(self.Style())
	
	
}