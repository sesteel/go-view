// line 1 "tokenizer.rl"
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package plaintext

//
// The Ragel version should support the -G0 target as G2 creates errors
//  ~/bin/ragel-6.8/ragel/ragel -Z -G0 tokenizer.rl -o tokenizer.go
//
import (
	. "github.com/sesteel/go-view/tokenizer"
)

// line 17 "tokenizer.rl"

// line 22 "tokenizer.go"
var _bindingGenerator_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 5,
	1, 6, 1, 7, 1, 8, 1, 9,
	1, 10, 1, 11, 1, 12, 1, 13,
	1, 14, 1, 15, 1, 16, 1, 17,
	1, 18, 1, 19, 1, 20, 1, 21,
	1, 22, 1, 23, 1, 24, 1, 25,
	1, 26, 1, 27, 1, 28, 1, 29,
	1, 30, 1, 31, 1, 32, 1, 33,
	1, 34, 1, 35, 1, 36, 1, 37,
	1, 38, 1, 39, 2, 2, 3, 2,
	2, 4,
}

var _bindingGenerator_to_state_actions []byte = []byte{
	0, 0, 0, 0, 1, 0, 0, 0,
	0, 0,
}

var _bindingGenerator_from_state_actions []byte = []byte{
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0,
}

const bindingGenerator_start int = 4
const bindingGenerator_first_final int = 4
const bindingGenerator_error int = 0

const bindingGenerator_en_main int = 4

// line 18 "tokenizer.rl"

func noop(a ...interface{}) {
	// do not remove
}

type _PlainTextTokenizer struct{}

func New() Tokenizer {
	return new(_PlainTextTokenizer)
}

func (self *_PlainTextTokenizer) Tokenize(text string) []*Token {
	data := []byte(text)
	var tokens []*Token

	// standard ragel preparedness
	cs, p, pe, eof := 0, 0, len(data), len(data)
	ts, te, act := 0, 0, 0
	lineCount := 1
	lineStart := 0
	var token *Token
	noop(ts, te, act)

	tkn := func(t TokenClass, s string) {
		val := string(data[ts:te])
		code := Codes[t]
		token = &Token{t, val, code, lineCount, ts - lineStart, te - lineStart, false}
		tokens = append(tokens, token)
	}

	// line 82 "tokenizer.go"
	{
		cs = bindingGenerator_start
		ts = 0
		te = 0
		act = 0
	}

	// line 90 "tokenizer.go"
	{
		var _acts int
		var _nacts uint

		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		_acts = int(_bindingGenerator_from_state_actions[cs])
		_nacts = uint(_bindingGenerator_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _bindingGenerator_actions[_acts-1] {
			case 1:
				// line 1 "NONE"

				ts = p

				// line 112 "tokenizer.go"
			}
		}

		switch cs {
		case 4:
			switch data[p] {
			case 9:
				goto tr4
			case 10:
				goto tr6
			case 13:
				goto tr7
			case 32:
				goto tr8
			case 33:
				goto tr9
			case 34:
				goto tr10
			case 35:
				goto tr11
			case 36:
				goto tr12
			case 37:
				goto tr13
			case 38:
				goto tr14
			case 39:
				goto tr15
			case 40:
				goto tr16
			case 41:
				goto tr17
			case 42:
				goto tr18
			case 43:
				goto tr19
			case 44:
				goto tr20
			case 45:
				goto tr21
			case 46:
				goto tr22
			case 47:
				goto tr23
			case 58:
				goto tr25
			case 59:
				goto tr26
			case 60:
				goto tr27
			case 61:
				goto tr28
			case 62:
				goto tr29
			case 63:
				goto tr30
			case 64:
				goto tr31
			case 91:
				goto tr33
			case 92:
				goto tr34
			case 93:
				goto tr35
			case 94:
				goto tr36
			case 96:
				goto tr5
			case 123:
				goto tr37
			case 124:
				goto tr38
			case 125:
				goto tr39
			case 160:
				goto tr5
			}
			switch {
			case data[p] < 126:
				switch {
				case data[p] > 31:
					if 48 <= data[p] && data[p] <= 57 {
						goto tr24
					}
				default:
					goto tr5
				}
			case data[p] > 127:
				switch {
				case data[p] < 154:
					if 129 <= data[p] && data[p] <= 152 {
						goto tr5
					}
				case data[p] > 158:
					if 241 <= data[p] {
						goto tr5
					}
				default:
					goto tr5
				}
			default:
				goto tr5
			}
			goto tr32
		case 0:
			goto _out
		case 5:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
			goto tr40
		case 6:
			switch data[p] {
			case 46:
				goto tr42
			case 69:
				goto tr43
			case 101:
				goto tr43
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
			goto tr41
		case 1:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1
			}
			goto tr0
		case 7:
			switch data[p] {
			case 69:
				goto tr43
			case 101:
				goto tr43
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1
			}
			goto tr41
		case 2:
			switch data[p] {
			case 43:
				goto tr2
			case 45:
				goto tr2
			}
			goto tr0
		case 3:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr3
			}
			goto tr0
		case 8:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr3
			}
			goto tr41
		case 9:
			switch data[p] {
			case 96:
				goto tr44
			case 160:
				goto tr44
			}
			switch {
			case data[p] < 123:
				switch {
				case data[p] < 58:
					if data[p] <= 47 {
						goto tr44
					}
				case data[p] > 64:
					if 91 <= data[p] && data[p] <= 94 {
						goto tr44
					}
				default:
					goto tr44
				}
			case data[p] > 127:
				switch {
				case data[p] < 154:
					if 129 <= data[p] && data[p] <= 152 {
						goto tr44
					}
				case data[p] > 158:
					if 241 <= data[p] {
						goto tr44
					}
				default:
					goto tr44
				}
			default:
				goto tr44
			}
			goto tr32
		}

	tr5:
		cs = 0
		goto _again
	tr42:
		cs = 1
		goto _again
	tr43:
		cs = 2
		goto _again
	tr2:
		cs = 3
		goto _again
	tr0:
		cs = 4
		goto f0
	tr4:
		cs = 4
		goto f4
	tr6:
		cs = 4
		goto f5
	tr7:
		cs = 4
		goto f6
	tr8:
		cs = 4
		goto f7
	tr9:
		cs = 4
		goto f8
	tr10:
		cs = 4
		goto f9
	tr11:
		cs = 4
		goto f10
	tr12:
		cs = 4
		goto f11
	tr13:
		cs = 4
		goto f12
	tr14:
		cs = 4
		goto f13
	tr15:
		cs = 4
		goto f14
	tr16:
		cs = 4
		goto f15
	tr17:
		cs = 4
		goto f16
	tr18:
		cs = 4
		goto f17
	tr20:
		cs = 4
		goto f19
	tr22:
		cs = 4
		goto f21
	tr23:
		cs = 4
		goto f22
	tr25:
		cs = 4
		goto f23
	tr26:
		cs = 4
		goto f24
	tr27:
		cs = 4
		goto f25
	tr28:
		cs = 4
		goto f26
	tr29:
		cs = 4
		goto f27
	tr30:
		cs = 4
		goto f28
	tr31:
		cs = 4
		goto f29
	tr33:
		cs = 4
		goto f30
	tr34:
		cs = 4
		goto f31
	tr35:
		cs = 4
		goto f32
	tr36:
		cs = 4
		goto f33
	tr37:
		cs = 4
		goto f34
	tr38:
		cs = 4
		goto f35
	tr39:
		cs = 4
		goto f36
	tr40:
		cs = 4
		goto f37
	tr41:
		cs = 4
		goto f38
	tr44:
		cs = 4
		goto f39
	tr19:
		cs = 5
		goto f18
	tr21:
		cs = 5
		goto f20
	tr24:
		cs = 6
		goto f1
	tr1:
		cs = 7
		goto f1
	tr3:
		cs = 8
		goto _again
	tr32:
		cs = 9
		goto _again

	f1:
		_acts = 5
		goto execFuncs
	f5:
		_acts = 7
		goto execFuncs
	f6:
		_acts = 9
		goto execFuncs
	f4:
		_acts = 11
		goto execFuncs
	f7:
		_acts = 13
		goto execFuncs
	f17:
		_acts = 15
		goto execFuncs
	f13:
		_acts = 17
		goto execFuncs
	f29:
		_acts = 19
		goto execFuncs
	f31:
		_acts = 21
		goto execFuncs
	f33:
		_acts = 23
		goto execFuncs
	f23:
		_acts = 25
		goto execFuncs
	f19:
		_acts = 27
		goto execFuncs
	f22:
		_acts = 29
		goto execFuncs
	f11:
		_acts = 31
		goto execFuncs
	f8:
		_acts = 33
		goto execFuncs
	f26:
		_acts = 35
		goto execFuncs
	f27:
		_acts = 37
		goto execFuncs
	f34:
		_acts = 39
		goto execFuncs
	f30:
		_acts = 41
		goto execFuncs
	f14:
		_acts = 43
		goto execFuncs
	f15:
		_acts = 45
		goto execFuncs
	f25:
		_acts = 47
		goto execFuncs
	f21:
		_acts = 49
		goto execFuncs
	f12:
		_acts = 51
		goto execFuncs
	f10:
		_acts = 53
		goto execFuncs
	f28:
		_acts = 55
		goto execFuncs
	f36:
		_acts = 57
		goto execFuncs
	f32:
		_acts = 59
		goto execFuncs
	f16:
		_acts = 61
		goto execFuncs
	f24:
		_acts = 63
		goto execFuncs
	f9:
		_acts = 65
		goto execFuncs
	f35:
		_acts = 67
		goto execFuncs
	f39:
		_acts = 69
		goto execFuncs
	f38:
		_acts = 71
		goto execFuncs
	f0:
		_acts = 73
		goto execFuncs
	f37:
		_acts = 75
		goto execFuncs
	f20:
		_acts = 77
		goto execFuncs
	f18:
		_acts = 80
		goto execFuncs

	execFuncs:
		_nacts = uint(_bindingGenerator_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _bindingGenerator_actions[_acts-1] {
			case 2:
				// line 1 "NONE"

				te = p + 1

			case 3:
				// line 112 "tokenizer.rl"

				act = 24
			case 4:
				// line 116 "tokenizer.rl"

				act = 28
			case 5:
				// line 85 "tokenizer.rl"

				te = p + 1
				{
					tkn(NEWLINE, "\n")
					lineCount++
					lineStart = ts
				}
			case 6:
				// line 91 "tokenizer.rl"

				te = p + 1
				{
					tkn(CR, "\r")
				}
			case 7:
				// line 92 "tokenizer.rl"

				te = p + 1
				{
					tkn(TAB, "\t")
				}
			case 8:
				// line 93 "tokenizer.rl"

				te = p + 1
				{
					tkn(SPACE, " ")
				}
			case 9:
				// line 95 "tokenizer.rl"

				te = p + 1
				{
					tkn(ASTERISK, "*")
				}
			case 10:
				// line 96 "tokenizer.rl"

				te = p + 1
				{
					tkn(AND, "&")
				}
			case 11:
				// line 97 "tokenizer.rl"

				te = p + 1
				{
					tkn(AT, "@")
				}
			case 12:
				// line 98 "tokenizer.rl"

				te = p + 1
				{
					tkn(BSLASH, "\\")
				}
			case 13:
				// line 99 "tokenizer.rl"

				te = p + 1
				{
					tkn(CARAT, "^")
				}
			case 14:
				// line 100 "tokenizer.rl"

				te = p + 1
				{
					tkn(COLON, ":")
				}
			case 15:
				// line 101 "tokenizer.rl"

				te = p + 1
				{
					tkn(COMMA, ",")
				}
			case 16:
				// line 102 "tokenizer.rl"

				te = p + 1
				{
					tkn(DIVIDE, "/")
				}
			case 17:
				// line 103 "tokenizer.rl"

				te = p + 1
				{
					tkn(DOLLAR, "/")
				}
			case 18:
				// line 104 "tokenizer.rl"

				te = p + 1
				{
					tkn(EXCLAM, "!")
				}
			case 19:
				// line 105 "tokenizer.rl"

				te = p + 1
				{
					tkn(EQUAL, "=")
				}
			case 20:
				// line 106 "tokenizer.rl"

				te = p + 1
				{
					tkn(GTHAN, ">")
				}
			case 21:
				// line 107 "tokenizer.rl"

				te = p + 1
				{
					tkn(LBRACE, "{")
				}
			case 22:
				// line 108 "tokenizer.rl"

				te = p + 1
				{
					tkn(LBRACK, "[")
				}
			case 23:
				// line 109 "tokenizer.rl"

				te = p + 1
				{
					tkn(SQUOTE, "'")
				}
			case 24:
				// line 110 "tokenizer.rl"

				te = p + 1
				{
					tkn(LPAREN, "(")
				}
			case 25:
				// line 111 "tokenizer.rl"

				te = p + 1
				{
					tkn(LTHAN, "<")
				}
			case 26:
				// line 114 "tokenizer.rl"

				te = p + 1
				{
					tkn(PERIOD, ".")
				}
			case 27:
				// line 115 "tokenizer.rl"

				te = p + 1
				{
					tkn(PERCENT, "%")
				}
			case 28:
				// line 117 "tokenizer.rl"

				te = p + 1
				{
					tkn(POUND, "#")
				}
			case 29:
				// line 118 "tokenizer.rl"

				te = p + 1
				{
					tkn(QMARK, "?")
				}
			case 30:
				// line 119 "tokenizer.rl"

				te = p + 1
				{
					tkn(RBRACE, "}")
				}
			case 31:
				// line 120 "tokenizer.rl"

				te = p + 1
				{
					tkn(RBRACK, "]")
				}
			case 32:
				// line 121 "tokenizer.rl"

				te = p + 1
				{
					tkn(RPAREN, ")")
				}
			case 33:
				// line 122 "tokenizer.rl"

				te = p + 1
				{
					tkn(SEMI, ";")
				}
			case 34:
				// line 123 "tokenizer.rl"

				te = p + 1
				{
					tkn(DQUOTE, "\"")
				}
			case 35:
				// line 125 "tokenizer.rl"

				te = p + 1
				{
					tkn(VBAR, "|")
				}
			case 36:
				// line 83 "tokenizer.rl"

				te = p
				p--
				{
					tkn(IDENTIFIER, "identifier")
				}
			case 37:
				// line 94 "tokenizer.rl"

				te = p
				p--
				{
					tkn(NUMBER_LITERAL, "numeric_literal")
				}
			case 38:
				// line 94 "tokenizer.rl"

				p = (te) - 1
				{
					tkn(NUMBER_LITERAL, "numeric_literal")
				}
			case 39:
				// line 1 "NONE"

				switch act {
				case 24:
					{
						p = (te) - 1
						tkn(MINUS, "-")
					}
				case 28:
					{
						p = (te) - 1
						tkn(PLUS, "+")
					}
				}

				// line 602 "tokenizer.go"
			}
		}
		goto _again

	_again:
		_acts = int(_bindingGenerator_to_state_actions[cs])
		_nacts = uint(_bindingGenerator_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _bindingGenerator_actions[_acts-1] {
			case 0:
				// line 1 "NONE"

				ts = 0

				// line 618 "tokenizer.go"
			}
		}

		if cs == 0 {
			goto _out
		}
		if p++; p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 5:
				goto tr40
			case 6:
				goto tr41
			case 1:
				goto tr0
			case 7:
				goto tr41
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 8:
				goto tr41
			case 9:
				goto tr44
			}
		}

	_out:
		{
		}
	}

	// line 130 "tokenizer.rl"

	return tokens
}
