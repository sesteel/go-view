// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package tokenizer

import ()

func ToLines(tkns []*Token) [][]*Token {
	lines := make([][]*Token, 0, 0)
	line := make([]*Token, 0)
	lines = append(lines, line)
	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]
		line = append(line, tkn)
		if tkn.Type == NEWLINE {
			line = make([]*Token, 0)
			lines = append(lines, line)
		}
	}
	return lines
}

type Rune struct {
	Type int
	Point rune 
}

//func ToRunes(tkns []*Token) [][]*codepoint {
//	lines := make([][]*Token, 0, 0)
//	line := make([]*Token, 0)
//	lines = append(lines, line)
//	for i := 0; i < len(tkns); i++ {
//		tkn := tkns[i]
//		line = append(line, tkn)
//		if tkn.Type == NEWLINE {
//			line = make([]*Token, 0)
//			lines = append(lines, line)
//		}
//	}
//	return lines
//}
