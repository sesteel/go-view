// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package event

import (
)

type EventDispatcher struct {
	MouseEventDispatcher
	FocusEventDispatcher
	KeyboardEventDispatcher
}

func (self *EventDispatcher) MouseButtonPress(me Mouse) {
	SetFocus(self)
	for i := 0; i < len(self.mouseButtonPressHandlers); i++ {
		self.mouseButtonPressHandlers[i](me)
	}
}
