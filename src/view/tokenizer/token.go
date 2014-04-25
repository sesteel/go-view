// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package tokenizer

import (
	"fmt"
)

type Token struct {
  Type     int
  Value    string
  Code     string
  Line     int
  Start    int 
  End      int
  Selected bool
}

func (t *Token) String() string {
   return fmt.Sprintf("{%s \"%s\" %d %d %d}", 
   	 Names[t.Type], t.Value, t.Line, t.Start, t.End)
}

