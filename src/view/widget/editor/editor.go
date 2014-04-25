// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package editor

import (
	"view"
	"view/doc"
)	

type Editor struct {
	view.DefaultComponent
	model doc.Document
}

func NewEditor(parent view.View, name string, model doc.Document) {
	e := &Editor{*view.NewComponent(parent, name), model}
	e.SetParent(parent)
	e.SetName(name)
	
}