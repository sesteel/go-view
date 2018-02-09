// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package tokenizer

import (
	"github.com/sesteel/go-view/common"
)

type Line struct {
	Characters []Character
	Bounds     []common.Bounds
}

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
			line = append(line, Character{pos, tkn, r})
			pos++
		}
		if tkn.Value == "\n" {
			lines = append(lines, line)
			line = make([]Character, 0)
		}
	}
	return lines
}

func ToLines(tkns []*Token) []Line {
	lines := make([]Line, 0, 0)
	line := Line{make([]Character, 0), make([]common.Bounds, 0)}
	pos := 0

	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]

		for _, r := range tkn.Value {
			line.Characters = append(line.Characters, Character{pos, tkn, r})
			line.Bounds = append(line.Bounds, common.Bounds{common.Point{-1, -1}, common.Size{-1, -1}})
			pos++
		}

		if tkn.Value == "\n" {
			lines = append(lines, line)
			line = Line{make([]Character, 0), make([]common.Bounds, 0)}
		}
	}
	lines = append(lines, line)
	return lines
}
