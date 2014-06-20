// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package main

import (
	"view"
	"view/color"
	"view/layout"
	"view/widget/editor"
)

func main() {
	var waitOnExit chan bool
	win := view.NewWindow("Editor Example", 0, 0, 600, 300)
	win.SetSize(800, 1000)
	l := layout.NewFill(win)

	tb := editor.New(win, "editor", TEXT)
	tb.KeywordStyle = editor.TokenStyle{view.FONT_WEIGHT_BOLD, view.FONT_SLANT_NORMAL, color.Blue1}
	tb.StringStyle = editor.TokenStyle{view.FONT_WEIGHT_BOLD, view.FONT_SLANT_NORMAL, color.Green2}
	tb.PrimitiveStyle = editor.TokenStyle{view.FONT_WEIGHT_BOLD, view.FONT_SLANT_ITALIC, color.Purple1}

	/*
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var
	*/
	tb.Keywords["func"] = true
	tb.Keywords["type"] = true
	tb.Keywords["break"] = true
	tb.Keywords["default"] = true
	tb.Keywords["case"] = true
	tb.Keywords["chan"] = true
	tb.Keywords["const"] = true
	tb.Keywords["continue"] = true
	tb.Keywords["defer"] = true
	tb.Keywords["else"] = true
	tb.Keywords["fallthrough"] = true
	tb.Keywords["for"] = true
	tb.Keywords["go"] = true
	tb.Keywords["goto"] = true
	tb.Keywords["if"] = true
	tb.Keywords["import"] = true
	tb.Keywords["interface"] = true
	tb.Keywords["map"] = true
	tb.Keywords["package"] = true
	tb.Keywords["range"] = true
	tb.Keywords["return"] = true
	tb.Keywords["select"] = true
	tb.Keywords["struct"] = true
	tb.Keywords["switch"] = true
	tb.Keywords["type"] = true
	tb.Keywords["var"] = true

	tb.Primitives["int"] = true
	tb.Primitives["int8"] = true
	tb.Primitives["int16"] = true
	tb.Primitives["int32"] = true
	tb.Primitives["int64"] = true
	tb.Primitives["uint"] = true
	tb.Primitives["uint8"] = true
	tb.Primitives["uint16"] = true
	tb.Primitives["uint32"] = true
	tb.Primitives["uint64"] = true
	tb.Primitives["float32"] = true
	tb.Primitives["float64"] = true
	tb.Primitives["string"] = true
	tb.Primitives["rune"] = true
	tb.Primitives["bool"] = true
	tb.Primitives["byte"] = true
	tb.Primitives["complex64"] = true
	tb.Primitives["complex128"] = true
	tb.Primitives["uintptr"] = true

	tb.DrawWhitespace = false
	tb.SetFocus(true)
	l.SetChild(tb)
	win.SetLayout(l)
	win.Start()
	<-waitOnExit
}

// TODO Fix the fact that an empty string cannot be edited or display a cursor
var TEXT = `one two three
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

// +build linux,!goci
package view

// #cgo pkg-config: cairo x11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xresource.h>
// #include <X11/keysymdef.h>
// #include <cairo/cairo-xlib.h>
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	//	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
	"unsafe"
	"view/event"
	"view/color"
)
`
