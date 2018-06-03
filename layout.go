// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	"github.com/sesteel/go-view/event"
)

// Layouts are special types that provide strategies
// for dividing the view space among components.
// They are charged with several important responsibilities:
//
// - They mask or otherwise protect the rest of the
//   target view surface from unintended drawing.
//
// - They define the boundries from which a component
//   or other composite view can be rendered
//
// - They distribute events to their child components
//
type Layout interface {
	Drawer
	Animator
	event.MouseHandler
}
