// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package tokenizer

import ()

func ToLinesOfTokens(tkns []*Token) [][]*Token {
	lines := make([][]*Token, 0, 0)
	line := make([]*Token, 0)
	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]
		line = append(line, tkn)
		if tkn.Value == "\n" {
			lines = append(lines, line)
			line = make([]*Token, 0)
		}
	}
	return lines
}

func ToLinesOfCharacters(tkns []*Token) [][]Character {
	lines := make([][]Character, 0, 0)
	line := make([]Character, 0)
	pos := 0
	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]
		for _, r := range tkn.Value {
			line = append(line, Character{pos, tkn, r, nil})
			pos++
		}
		if tkn.Value == "\n" {
			lines = append(lines, line)
			line = make([]Character, 0)
		}
	}
	return lines
}
