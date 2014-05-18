
// line 1 "word_tokenizer.rl"
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package tokenizer

//
// The Ragel version should support the -G0 target as G2 creates errors
//  ~/bin/ragel-6.8/ragel/ragel -Z -G0 tokenizer.rl -o tokenizer.go
//

// line 14 "word_tokenizer.rl"

// line 19 "word_tokenizer.go"
var _bindingGenerator_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 1, 15, 
	1, 16, 1, 17, 1, 18, 1, 19, 
	1, 20, 1, 21, 1, 22, 1, 23, 
	1, 24, 1, 25, 1, 26, 1, 27, 
	1, 28, 1, 29, 1, 30, 1, 31, 
	1, 32, 1, 33, 1, 34, 1, 35, 
	1, 36, 1, 37, 1, 38, 1, 39, 
	1, 40, 1, 41, 1, 42, 1, 43, 
	1, 44, 1, 45, 2, 2, 3, 2, 
	2, 4, 2, 2, 5, 2, 2, 6, 
}

var _bindingGenerator_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _bindingGenerator_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 3, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

const bindingGenerator_start int = 9
const bindingGenerator_first_final int = 9
const bindingGenerator_error int = 0

const bindingGenerator_en_main int = 9


// line 15 "word_tokenizer.rl"

func noop(a ... interface{}){
		// do not remove
}

func Tokenize(text string) []*Token {
	data := []byte(text)
	var tokens []*Token
	
    // standard ragel preparedness
    cs, p, pe, eof := 0, 0, len(data), len(data)
    ts, te, act    := 0, 0, 0
    lineCount      := 1
    lineStart      := 0
    var token *Token  
    noop(ts, te, act)
   	
    tkn := func(t int, s string) {
      val := string(data[ts:te])
      code:= codes[t]
      token = &Token{t, val, code, lineCount, ts-lineStart, te-lineStart, false}
      tokens = append(tokens, token)
    }
       
    
// line 80 "word_tokenizer.go"
	{
	cs = bindingGenerator_start
	ts = 0
	te = 0
	act = 0
	}

// line 88 "word_tokenizer.go"
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
	_nacts = uint(_bindingGenerator_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _bindingGenerator_actions[_acts - 1] {
		case 1:
// line 1 "NONE"

ts = p

// line 110 "word_tokenizer.go"
		}
	}

	switch cs {
	case 9:
		switch data[p] {
		case 9:
			goto tr12;
		case 10:
			goto tr14;
		case 13:
			goto tr15;
		case 32:
			goto tr16;
		case 33:
			goto tr17;
		case 34:
			goto tr18;
		case 35:
			goto tr19;
		case 36:
			goto tr20;
		case 37:
			goto tr21;
		case 38:
			goto tr22;
		case 39:
			goto tr23;
		case 40:
			goto tr24;
		case 41:
			goto tr25;
		case 42:
			goto tr26;
		case 43:
			goto tr27;
		case 44:
			goto tr28;
		case 45:
			goto tr29;
		case 46:
			goto tr30;
		case 47:
			goto tr31;
		case 48:
			goto tr32;
		case 58:
			goto tr34;
		case 59:
			goto tr35;
		case 60:
			goto tr36;
		case 61:
			goto tr37;
		case 62:
			goto tr38;
		case 63:
			goto tr39;
		case 64:
			goto tr40;
		case 91:
			goto tr42;
		case 92:
			goto tr43;
		case 93:
			goto tr44;
		case 94:
			goto tr45;
		case 96:
			goto tr13;
		case 123:
			goto tr46;
		case 124:
			goto tr47;
		case 125:
			goto tr48;
		case 160:
			goto tr13;
		}
		switch {
		case data[p] < 126:
			switch {
			case data[p] > 31:
				if 49 <= data[p] && data[p] <= 57 {
					goto tr33;
				}
			default:
				goto tr13;
			}
		case data[p] > 127:
			switch {
			case data[p] < 154:
				if 129 <= data[p] && data[p] <= 152 {
					goto tr13;
				}
			case data[p] > 158:
				if 241 <= data[p] {
					goto tr13;
				}
			default:
				goto tr13;
			}
		default:
			goto tr13;
		}
		goto tr41;
	case 0:
		goto _out
	case 10:
		if data[p] == 34 {
			goto tr2;
		}
		goto tr1;
	case 1:
		if data[p] == 34 {
			goto tr2;
		}
		goto tr1;
	case 11:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr19;
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr19;
			}
		default:
			goto tr19;
		}
		goto tr50;
	case 12:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr33;
		}
		goto tr3;
	case 13:
		switch data[p] {
		case 46:
			goto tr52;
		case 69:
			goto tr53;
		case 101:
			goto tr53;
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr33;
		}
		goto tr51;
	case 2:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr4;
		}
		goto tr3;
	case 14:
		switch data[p] {
		case 69:
			goto tr53;
		case 101:
			goto tr53;
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr4;
		}
		goto tr51;
	case 3:
		switch data[p] {
		case 43:
			goto tr5;
		case 45:
			goto tr5;
		}
		goto tr3;
	case 4:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6;
		}
		goto tr3;
	case 15:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr6;
		}
		goto tr51;
	case 16:
		switch data[p] {
		case 46:
			goto tr52;
		case 66:
			goto tr55;
		case 69:
			goto tr53;
		case 88:
			goto tr19;
		case 98:
			goto tr55;
		case 101:
			goto tr53;
		case 120:
			goto tr19;
		}
		switch {
		case data[p] > 55:
			if 56 <= data[p] && data[p] <= 57 {
				goto tr33;
			}
		case data[p] >= 48:
			goto tr54;
		}
		goto tr51;
	case 17:
		switch data[p] {
		case 46:
			goto tr52;
		case 69:
			goto tr53;
		case 101:
			goto tr53;
		}
		switch {
		case data[p] > 55:
			if 56 <= data[p] && data[p] <= 57 {
				goto tr33;
			}
		case data[p] >= 48:
			goto tr54;
		}
		goto tr56;
	case 18:
		if data[p] == 32 {
			goto tr58;
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto tr55;
		}
		goto tr57;
	case 5:
		if 48 <= data[p] && data[p] <= 49 {
			goto tr8;
		}
		goto tr7;
	case 6:
		if 48 <= data[p] && data[p] <= 49 {
			goto tr9;
		}
		goto tr7;
	case 7:
		if 48 <= data[p] && data[p] <= 49 {
			goto tr10;
		}
		goto tr7;
	case 8:
		if 48 <= data[p] && data[p] <= 49 {
			goto tr11;
		}
		goto tr7;
	case 19:
		if data[p] == 32 {
			goto tr58;
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto tr8;
		}
		goto tr57;
	case 20:
		switch data[p] {
		case 96:
			goto tr59;
		case 160:
			goto tr59;
		}
		switch {
		case data[p] < 123:
			switch {
			case data[p] < 58:
				if data[p] <= 47 {
					goto tr59;
				}
			case data[p] > 64:
				if 91 <= data[p] && data[p] <= 94 {
					goto tr59;
				}
			default:
				goto tr59;
			}
		case data[p] > 127:
			switch {
			case data[p] < 154:
				if 129 <= data[p] && data[p] <= 152 {
					goto tr59;
				}
			case data[p] > 158:
				if 241 <= data[p] {
					goto tr59;
				}
			default:
				goto tr59;
			}
		default:
			goto tr59;
		}
		goto tr41;
	}

	tr13: cs = 0; goto _again
	tr1: cs = 1; goto _again
	tr52: cs = 2; goto _again
	tr53: cs = 3; goto _again
	tr5: cs = 4; goto _again
	tr58: cs = 5; goto _again
	tr8: cs = 6; goto _again
	tr9: cs = 7; goto _again
	tr10: cs = 8; goto _again
	tr0: cs = 9; goto f0
	tr2: cs = 9; goto f1
	tr3: cs = 9; goto f2
	tr7: cs = 9; goto f4
	tr12: cs = 9; goto f8
	tr14: cs = 9; goto f9
	tr15: cs = 9; goto f10
	tr16: cs = 9; goto f11
	tr17: cs = 9; goto f12
	tr20: cs = 9; goto f13
	tr21: cs = 9; goto f14
	tr22: cs = 9; goto f15
	tr23: cs = 9; goto f16
	tr24: cs = 9; goto f17
	tr25: cs = 9; goto f18
	tr26: cs = 9; goto f19
	tr28: cs = 9; goto f21
	tr30: cs = 9; goto f23
	tr31: cs = 9; goto f24
	tr34: cs = 9; goto f25
	tr35: cs = 9; goto f26
	tr36: cs = 9; goto f27
	tr37: cs = 9; goto f28
	tr38: cs = 9; goto f29
	tr39: cs = 9; goto f30
	tr40: cs = 9; goto f31
	tr42: cs = 9; goto f32
	tr43: cs = 9; goto f33
	tr44: cs = 9; goto f34
	tr45: cs = 9; goto f35
	tr46: cs = 9; goto f36
	tr47: cs = 9; goto f37
	tr48: cs = 9; goto f38
	tr49: cs = 9; goto f39
	tr50: cs = 9; goto f40
	tr51: cs = 9; goto f41
	tr56: cs = 9; goto f43
	tr57: cs = 9; goto f44
	tr59: cs = 9; goto f45
	tr18: cs = 10; goto f5
	tr19: cs = 11; goto _again
	tr27: cs = 12; goto f20
	tr29: cs = 12; goto f22
	tr33: cs = 13; goto f3
	tr4: cs = 14; goto f3
	tr6: cs = 15; goto _again
	tr32: cs = 16; goto f3
	tr54: cs = 17; goto f42
	tr55: cs = 18; goto f5
	tr11: cs = 19; goto f5
	tr41: cs = 20; goto _again

	f5: _acts = 5; goto execFuncs
	f9: _acts = 7; goto execFuncs
	f10: _acts = 9; goto execFuncs
	f8: _acts = 11; goto execFuncs
	f11: _acts = 13; goto execFuncs
	f1: _acts = 15; goto execFuncs
	f19: _acts = 17; goto execFuncs
	f15: _acts = 19; goto execFuncs
	f31: _acts = 21; goto execFuncs
	f33: _acts = 23; goto execFuncs
	f35: _acts = 25; goto execFuncs
	f25: _acts = 27; goto execFuncs
	f21: _acts = 29; goto execFuncs
	f24: _acts = 31; goto execFuncs
	f13: _acts = 33; goto execFuncs
	f12: _acts = 35; goto execFuncs
	f28: _acts = 37; goto execFuncs
	f29: _acts = 39; goto execFuncs
	f36: _acts = 41; goto execFuncs
	f32: _acts = 43; goto execFuncs
	f16: _acts = 45; goto execFuncs
	f17: _acts = 47; goto execFuncs
	f27: _acts = 49; goto execFuncs
	f23: _acts = 51; goto execFuncs
	f14: _acts = 53; goto execFuncs
	f30: _acts = 55; goto execFuncs
	f38: _acts = 57; goto execFuncs
	f34: _acts = 59; goto execFuncs
	f18: _acts = 61; goto execFuncs
	f26: _acts = 63; goto execFuncs
	f37: _acts = 65; goto execFuncs
	f45: _acts = 67; goto execFuncs
	f44: _acts = 69; goto execFuncs
	f40: _acts = 71; goto execFuncs
	f43: _acts = 73; goto execFuncs
	f41: _acts = 75; goto execFuncs
	f39: _acts = 77; goto execFuncs
	f4: _acts = 79; goto execFuncs
	f0: _acts = 81; goto execFuncs
	f2: _acts = 83; goto execFuncs
	f42: _acts = 85; goto execFuncs
	f3: _acts = 88; goto execFuncs
	f22: _acts = 91; goto execFuncs
	f20: _acts = 94; goto execFuncs

execFuncs:
	_nacts = uint(_bindingGenerator_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _bindingGenerator_actions[_acts - 1] {
		case 2:
// line 1 "NONE"

te = p+1

		case 3:
// line 95 "word_tokenizer.rl"

act = 8;
		case 4:
// line 96 "word_tokenizer.rl"

act = 9;
		case 5:
// line 116 "word_tokenizer.rl"

act = 29;
		case 6:
// line 120 "word_tokenizer.rl"

act = 33;
		case 7:
// line 84 "word_tokenizer.rl"

te = p+1
{ 
          		tkn(NEWLINE, "\n")
          		lineCount++
                lineStart = ts;
          	}
		case 8:
// line 90 "word_tokenizer.rl"

te = p+1
{ tkn(CR,         "\r")             }
		case 9:
// line 91 "word_tokenizer.rl"

te = p+1
{ tkn(TAB,        "\t")             }
		case 10:
// line 92 "word_tokenizer.rl"

te = p+1
{ tkn(SPACE,      " ")              }
		case 11:
// line 97 "word_tokenizer.rl"

te = p+1
{ tkn(STRING_LIT, "string_literal") }
		case 12:
// line 98 "word_tokenizer.rl"

te = p+1
{ tkn(ASTERISK,   "*")              }
		case 13:
// line 99 "word_tokenizer.rl"

te = p+1
{ tkn(AND,        "&")              }
		case 14:
// line 100 "word_tokenizer.rl"

te = p+1
{ tkn(AT,         "@")              }
		case 15:
// line 101 "word_tokenizer.rl"

te = p+1
{ tkn(BSLASH,     "\\")             }
		case 16:
// line 102 "word_tokenizer.rl"

te = p+1
{ tkn(CARAT,      "^")              }
		case 17:
// line 103 "word_tokenizer.rl"

te = p+1
{ tkn(COLON,      ":")              }
		case 18:
// line 104 "word_tokenizer.rl"

te = p+1
{ tkn(COMMA,      ",")              }
		case 19:
// line 105 "word_tokenizer.rl"

te = p+1
{ tkn(DIVIDE,     "/")              }
		case 20:
// line 106 "word_tokenizer.rl"

te = p+1
{ tkn(DOLLAR,     "/")              }
		case 21:
// line 107 "word_tokenizer.rl"

te = p+1
{ tkn(EXCLAM,     "!")              }
		case 22:
// line 108 "word_tokenizer.rl"

te = p+1
{ tkn(EQUAL,      "=")              }
		case 23:
// line 109 "word_tokenizer.rl"

te = p+1
{ tkn(GTHAN,      ">")              }
		case 24:
// line 111 "word_tokenizer.rl"

te = p+1
{ tkn(LBRACE,     "{")              }
		case 25:
// line 112 "word_tokenizer.rl"

te = p+1
{ tkn(LBRACK,     "[")              }
		case 26:
// line 113 "word_tokenizer.rl"

te = p+1
{ tkn(SQUOTE,     "'")              }
		case 27:
// line 114 "word_tokenizer.rl"

te = p+1
{ tkn(LPAREN,     "(")              }
		case 28:
// line 115 "word_tokenizer.rl"

te = p+1
{ tkn(LTHAN,      "<")              }
		case 29:
// line 118 "word_tokenizer.rl"

te = p+1
{ tkn(PERIOD,     ".")              }
		case 30:
// line 119 "word_tokenizer.rl"

te = p+1
{ tkn(PERCENT,    "%")              }
		case 31:
// line 122 "word_tokenizer.rl"

te = p+1
{ tkn(QMARK,      "?")              }
		case 32:
// line 123 "word_tokenizer.rl"

te = p+1
{ tkn(RBRACE,     "}")              }
		case 33:
// line 124 "word_tokenizer.rl"

te = p+1
{ tkn(RBRACK,     "]")              }
		case 34:
// line 125 "word_tokenizer.rl"

te = p+1
{ tkn(RPAREN,     ")")              }
		case 35:
// line 126 "word_tokenizer.rl"

te = p+1
{ tkn(SEMI,       ";")              }
		case 36:
// line 129 "word_tokenizer.rl"

te = p+1
{ tkn(VBAR,       "|")              }
		case 37:
// line 82 "word_tokenizer.rl"

te = p
p--
{ tkn(IDENTIFIER, "identifier")}
		case 38:
// line 93 "word_tokenizer.rl"

te = p
p--
{ tkn(BINARY_LIT, "binary_literal") }
		case 39:
// line 94 "word_tokenizer.rl"

te = p
p--
{ tkn(HEX_LIT,    "hex_literal")    }
		case 40:
// line 95 "word_tokenizer.rl"

te = p
p--
{ tkn(OCTAL_LIT,  "octal_literal")  }
		case 41:
// line 96 "word_tokenizer.rl"

te = p
p--
{ tkn(NUMBER_LIT, "numeric_literal")}
		case 42:
// line 127 "word_tokenizer.rl"

te = p
p--
{ tkn(DQUOTE,     "\"")             }
		case 43:
// line 93 "word_tokenizer.rl"

p = (te) - 1
{ tkn(BINARY_LIT, "binary_literal") }
		case 44:
// line 127 "word_tokenizer.rl"

p = (te) - 1
{ tkn(DQUOTE,     "\"")             }
		case 45:
// line 1 "NONE"

	switch act {
	case 8:
	{p = (te) - 1
 tkn(OCTAL_LIT,  "octal_literal")  }
	case 9:
	{p = (te) - 1
 tkn(NUMBER_LIT, "numeric_literal")}
	case 29:
	{p = (te) - 1
 tkn(MINUS,      "-")              }
	case 33:
	{p = (te) - 1
 tkn(PLUS,       "+")              }
	}
	
// line 765 "word_tokenizer.go"
		}
	}
	goto _again

_again:
	_acts = int(_bindingGenerator_to_state_actions[cs])
	_nacts = uint(_bindingGenerator_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _bindingGenerator_actions[_acts - 1] {
		case 0:
// line 1 "NONE"

ts = 0

// line 781 "word_tokenizer.go"
		}
	}

	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		switch cs {
		case 10:
			goto tr49
		case 1:
			goto tr0
		case 11:
			goto tr50
		case 12:
			goto tr3
		case 13:
			goto tr51
		case 2:
			goto tr3
		case 14:
			goto tr51
		case 3:
			goto tr3
		case 4:
			goto tr3
		case 15:
			goto tr51
		case 16:
			goto tr51
		case 17:
			goto tr56
		case 18:
			goto tr57
		case 5:
			goto tr7
		case 6:
			goto tr7
		case 7:
			goto tr7
		case 8:
			goto tr7
		case 19:
			goto tr57
		case 20:
			goto tr59
	}
	}

	_out: {}
	}

// line 134 "word_tokenizer.rl"

    
    return tokens
}

