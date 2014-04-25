// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import ()

type ScrollOffset float64 

type Size struct {
	Width, Height float64
}

type Bounds struct {
	X, Y float64
	Size
}

func (b Bounds) Contains(x, y float64) bool {
	return x >= b.X && 
	       x <= (b.X + b.Width) && 
	       y >= b.Y && 
	       y <= (b.Y + b.Height)
}

