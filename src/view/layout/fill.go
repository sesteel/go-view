// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package layout

import (
	"view"
)

// Fill is a simple layout that expands a single
// child to the size allotted to the layout by 
// the target.
type Fill struct {
	target view.View
	child  view.Drawer
}

func NewFill(target view.View) *Fill {
	l := new(Fill)
	l.target = target
	return l
}

func (self *Fill) SetChild(d view.Drawer) {
	self.child = d
}

func (self *Fill) Child() view.Drawer {
	return self.child 
}

func (self *Fill) Draw(surface *view.Surface) {
	self.child.Draw(surface)
}
