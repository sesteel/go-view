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
	. "view/tokenizer"
)
%% machine bindingGenerator;
%% write data;

func noop(a ... interface{}){
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
    ts, te, act    := 0, 0, 0
    lineCount      := 1
    lineStart      := 0
    var token *Token  
    noop(ts, te, act)
   	
    tkn := func(t TokenClass, s string) {
      val := string(data[ts:te])
      code:= Codes[t]
      token = &Token{t, val, code, lineCount, ts-lineStart, te-lineStart, false}
      tokens = append(tokens, token)
    }
       
    %%{
        identifier         = [_a-zA-Z¡-🙀][_a-zA-Z0-9¡-🙀]*;
        newline            = '\n';
        carriage           = '\r';
        tab                = '\t';
        spacechar          = ' ';
        number             = ('+'|'-')? [0-9]+ ('.' [0-9]+)? (('e'|'E') ('+'|'-')[0-9]+)?;
        and                = '&';
        ast                = '*';
        at                 = '@';
		dollar             = '$';
        percent            = '%';
        pound              = '#';
        carat              = '^';
        exclam             = '!';
        lparen             = '(';
        rparen             = ')';
        underscore         = '_';
        minus              = '-';
        plus               = '+';
    	equal              = '=';
        lbrace             = '{';
        lbrack             = '[';
        rbrace             = '}';
        rbrack             = ']';
        bslash             = '\\';
        colon              = ':';
        comma              = ',';
        divide             = '/';
        qmark              = "?";
        gthan              = '>';
        squote             = '\'';
        lthan              = '<';
        period             = '.';
        semi               = ';';
        dquote             = '"';
        vbar               = '|'; 
               
        main := |*
          	identifier         => { tkn(IDENTIFIER, "identifier")};  
          	
          	newline            => { 
          		tkn(NEWLINE, "\n")
          		lineCount++
                lineStart = ts;
          	};
          	
        	carriage           => { tkn(CR,         "\r")             };
        	tab                => { tkn(TAB,        "\t")             };
        	spacechar          => { tkn(SPACE,      " ")              };   
          	number             => { tkn(NUMBER_LITERAL, "numeric_literal")};   
        	ast                => { tkn(ASTERISK,   "*")              };
            and                => { tkn(AND,        "&")              };
            at                 => { tkn(AT,         "@")              };           
            bslash             => { tkn(BSLASH,     "\\")             };
            carat              => { tkn(CARAT,      "^")              };
            colon              => { tkn(COLON,      ":")              };
            comma              => { tkn(COMMA,      ",")              };
            divide             => { tkn(DIVIDE,     "/")              };
            dollar             => { tkn(DOLLAR,     "/")              };
            exclam             => { tkn(EXCLAM,     "!")              };
            equal              => { tkn(EQUAL,      "=")              };
            gthan              => { tkn(GTHAN,      ">")              };
            lbrace             => { tkn(LBRACE,     "{")              };
            lbrack             => { tkn(LBRACK,     "[")              };
            squote             => { tkn(SQUOTE,     "'")              };
            lparen             => { tkn(LPAREN,     "(")              };
            lthan              => { tkn(LTHAN,      "<")              };
            minus              => { tkn(MINUS,      "-")              };
            number             => { tkn(NUMBER,     "n")              };
            period             => { tkn(PERIOD,     ".")              };
            percent            => { tkn(PERCENT,    "%")              };
            plus               => { tkn(PLUS,       "+")              };
            pound              => { tkn(POUND,      "#")              };
            qmark              => { tkn(QMARK,      "?")              };
            rbrace             => { tkn(RBRACE,     "}")              };
            rbrack             => { tkn(RBRACK,     "]")              };
            rparen             => { tkn(RPAREN,     ")")              };
            semi               => { tkn(SEMI,       ";")              };
            dquote             => { tkn(DQUOTE,     "\"")             };
            underscore         => { tkn(UNDERSCORE, "_")              };
            vbar               => { tkn(VBAR,       "|")              };   
        *|;
       
        write init;
        write exec;
    }%%
    
    return tokens
}

