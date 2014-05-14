// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package main

import (
	"view"
	"view/layout"
	"view/widget/editor"
)

var TEXT = `
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`

func main() {
	mdl := editor.NewModel("main.go", TEXT)
	var waitOnExit chan bool
	win := view.NewWindow("Text Box Example", 0, 0, 600, 300)
	win.SetSize(600, 300)
	l  := layout.NewFill(win)
	tb := editor.New(win, "editor", mdl)
	l.SetChild(tb) 
	win.SetLayout(l)
	win.Start()
	<-waitOnExit
}
