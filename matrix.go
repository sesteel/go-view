// Extended from https://github.com/ungerik/go-cairo
// 
// Copyright © 2002 University of Southern California
// Copyright © 2005 Red Hat, Inc.
// 
// This library is free software; you can redistribute it and/or
// modify it either under the terms of the GNU Lesser General Public
// License version 2.1 as published by the Free Software Foundation
// (the "LGPL") or, at your option, under the terms of the Mozilla
// Public License Version 1.1 (the "MPL"). If you do not alter this
// notice, a recipient may use your version of this file under either
// the MPL or the LGPL.
// 
// You should have received a copy of the LGPL along with this library
// in the file COPYING-LGPL-2.1; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Suite 500, Boston, MA 02110-1335, USA
// You should have received a copy of the MPL along with this library
// in the file COPYING-MPL-1.1
// 
// The contents of this file are subject to the Mozilla Public License
// Version 1.1 (the "License"); you may not use this file except in
// compliance with the License. You may obtain a copy of the License at
// http://www.mozilla.org/MPL/
// 
// This software is distributed on an "AS IS" basis, WITHOUT WARRANTY
// OF ANY KIND, either express or implied. See the LGPL or the MPL for
// the specific language governing rights and limitations.
// 
// The Original Code is the cairo graphics library.
// 
// The Initial Developer of the Original Code is University of Southern
// California.
// 
// Contributor(s):
// 	Carl D. Worth <cworth@cworth.org>

// +build !goci
package view

// #include <cairo/cairo.h>
import "C"

import (
	"unsafe"
)

type Matrix struct {
	Xx, Yx float64
	Xy, Yy float64
	X0, Y0 float64
}

func (self *Matrix) cairo_matrix_t() *C.cairo_matrix_t {
	return (*C.cairo_matrix_t)(unsafe.Pointer(self))
}

func (self *Matrix) InitIdendity() {
	C.cairo_matrix_init_identity(self.cairo_matrix_t())
}

func (self *Matrix) InitTranslate(tx, ty float64) {
	C.cairo_matrix_init_translate(self.cairo_matrix_t(), C.double(tx), C.double(ty))
}

func (self *Matrix) InitScale(sx, sy float64) {
	C.cairo_matrix_init_scale(self.cairo_matrix_t(), C.double(sx), C.double(sy))
}

func (self *Matrix) InitRotate(radians float64) {
	C.cairo_matrix_init_rotate(self.cairo_matrix_t(), C.double(radians))
}

func (self *Matrix) Translate(tx, ty float64) {
	C.cairo_matrix_translate(self.cairo_matrix_t(), C.double(tx), C.double(ty))
}

func (self *Matrix) Scale(sx, sy float64) {
	C.cairo_matrix_scale(self.cairo_matrix_t(), C.double(sx), C.double(sy))
}

func (self *Matrix) Rotate(radians float64) {
	C.cairo_matrix_rotate(self.cairo_matrix_t(), C.double(radians))
}

func (self *Matrix) Invert() {
	C.cairo_matrix_invert(self.cairo_matrix_t())
}

func (self *Matrix) Multiply(a, b Matrix) {
	C.cairo_matrix_multiply(self.cairo_matrix_t(), a.cairo_matrix_t(), b.cairo_matrix_t())
}

func (self *Matrix) TransformDistance(dx, dy float64) (float64, float64) {
	C.cairo_matrix_transform_distance(self.cairo_matrix_t(),
		(*C.double)(unsafe.Pointer(&dx)), (*C.double)(unsafe.Pointer(&dy)))
	return dx, dy
}

func (self *Matrix) TransformPoint(x, y float64) (float64, float64) {
	C.cairo_matrix_transform_point(self.cairo_matrix_t(),
		(*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)))
	return x, y
}
