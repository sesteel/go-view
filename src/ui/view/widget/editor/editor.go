package editor

import (
	"ui/view"
	"ui/doc"
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