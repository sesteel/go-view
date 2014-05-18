// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package main

import (
	"fmt"
	"view"
	"view/layout"
	"view/widget/editor"
	"view/tokenizer" 
)

var TEXT = `
package main

import "fmt"

func main() { 
	fmt.Println("Hello World!")
}
`

func main() {
	tkns := tokenizer.Tokenize(TEXT)
	for _, t := range tkns {
		fmt.Println("->", t)
	}
	
	var waitOnExit chan bool
	win := view.NewWindow("Editor Example", 0, 0, 600, 300)
	win.SetSize(600, 300)
	l  := layout.NewFill(win)
	
	mdl := editor.NewModel("main.go", TEXT)
	tb := editor.New(win, "editor", mdl)  
	l.SetChild(tb) 
	win.SetLayout(l)
	win.Start()
	<-waitOnExit
}
