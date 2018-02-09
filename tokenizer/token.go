// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package tokenizer

import (
	"fmt"
)

type TokenClass int

var Codes map[TokenClass]rune = map[TokenClass]rune{
	IDENTIFIER: 'e',

	NUMBER_LITERAL: 'i',
	STRING_LITERAL: 's',
	BINARY_LITERAL: 'b',
	HEX_LITERAL:    'h',
	OCTAL_LITERAL:  'o',
	COMMENT:        'c',
	ASSIGN:         'a',
	ASTERISK:       '*',
	AND:            '&',
	AT:             '@',
	BSLASH:         '\\',
	CARAT:          '^',
	COLON:          ':',
	COMMA:          ',',
	DIVIDE:         '/',
	DOLLAR:         '$',
	EXCLAM:         '!',
	EQUAL:          '=',
	GTHAN:          '>',
	LBRACE:         '{',
	LBRACK:         '[',
	LPAREN:         '(',
	LTHAN:          '<',
	MINUS:          '-',
	PERIOD:         '.',
	PERCENT:        '%',
	PLUS:           '+',
	POUND:          '#',
	QMARK:          '?',
	RBRACE:         '}',
	RBRACK:         ']',
	RPAREN:         ')',
	SEMI:           ';',
	DQUOTE:         '"',
	UNDERSCORE:     '_',
	VBAR:           '|',
	NEWLINE:        '\n',
	CR:             '\r',
	TAB:            '\t',
	SPACE:          ' ',
	SQUOTE:         '\'',
}

var Names []string = []string{
	"identifier",
	"numeric literal",
	"string literal",
	"binary literal",
	"hex literal",
	"octal literal",
	"comment",
	"assign",
	"*",
	"&",
	"@",
	"\\",
	"^",
	":",
	",",
	"/",
	"$",
	"!",
	"=",
	">",
	"{",
	"[",
	"(",
	"<",
	"-",
	".",
	"%",
	"+",
	"#",
	"?",
	"}",
	"]",
	")",
	";",
	"\"",
	"_",
	"|",
	"\n",
	"\r",
	"\t",
	" ",
	"'"}

const (
	IDENTIFIER TokenClass = iota
	NUMBER_LITERAL
	STRING_LITERAL
	BINARY_LITERAL
	HEX_LITERAL
	OCTAL_LITERAL
	COMMENT
	ASSIGN
	ASTERISK
	AND
	AT
	BSLASH
	CARAT
	COLON
	COMMA
	DIVIDE
	DOLLAR
	EXCLAM
	EQUAL
	GTHAN
	LBRACE
	LBRACK
	LPAREN
	LTHAN
	MINUS
	PERIOD
	PERCENT
	PLUS
	POUND
	QMARK
	RBRACE
	RBRACK
	RPAREN
	SEMI
	DQUOTE
	UNDERSCORE
	VBAR
	NEWLINE
	CR
	TAB
	SPACE
	SQUOTE
)

func (self TokenClass) Whitespace() bool {
	switch self {
	case SPACE, TAB, CR, NEWLINE:
		return true
	default:
		return false
	}
}

type Token struct {
	Type     TokenClass
	Value    string
	Code     rune
	Line     int
	Start    int
	End      int
	Selected bool
}

func (t *Token) String() string {
	return fmt.Sprintf("{%s \"%s\" %d %d %d}", t.Type, t.Value, t.Line, t.Start, t.End)
}
