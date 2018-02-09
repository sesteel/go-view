// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package editor

// Range represents an group of characters defined by starting and ending (line, column) positions.
type Range struct {
	Start Index
	End   Index
}

func (self *Range) Normalize() {
	if self.Start.Line >= self.End.Line {
		self.Start.Line, self.End.Line = self.End.Line, self.Start.Line
		self.Start.Column, self.End.Column = self.End.Column, self.Start.Column
	}
}
