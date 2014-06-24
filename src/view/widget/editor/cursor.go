// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

import ()

// Cursor is used to store the position of the cursor via a Index.
type Cursor Index

func (self Cursor) PreviousPos(lines Lines) Index {
	if len(lines) < self.Line-1 || self.Line < 0 {
		return Index(self)
	}

	if self.Column == 0 {
		return Index{self.Line - 1, len(lines[self.Line-1])}
	} else {
		return Index{self.Line, self.Column - 1}
	}
}
