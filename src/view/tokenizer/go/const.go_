// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package tokenizer

import (

)

var codes map[TokenClass]string = map[TokenClass]string {
  PURE           :"A",
  TYPE           :"T",
  INTERFACE      :"I",
  OP             :"D",
  IF             :"E",
  ELSE           :"F",
  DEFER          :"G",
  WHEN           :"H",
  IS             :"N",
  FOR            :"J",
  SELECT         :"K",
  SHARED         :"L",
  MODULE         :"M",
  MAP_TYPE       :"p",
  LIST_TYPE      :"l",
  TREE_TYPE      :"t",
  TRIE_TYPE      :"y",
  FUNC_TYPE      :"f",
  CHAN_TYPE      :"c",
  MAT_TYPE       :"m",
  VEC_TYPE       :"v",
  UINT_TYPE      :"u",
  INT_TYPE       :"i",
  NUM_TYPE       :"n",
  FLOAT_TYPE     :"f",
  BOOL_TYPE      :"b",
  BYTE_TYPE      :"B",
  TEXT_TYPE      :"a",
  ERROR_TYPE     :"b",
  EXCEPTION_TYPE :"c",
  UNIT_TYPE      :"d", 
  IDENTIFIER     :"e",
  BINARY_LIT     :"z",
  HEX_LIT        :"g",
  OCTAL_LIT      :"h",
  NUMBER_LIT     :"i",
  STRING_LIT     :"j",
  ASSIGN         :"k",
  ASTERISK       :"*",
  AND            :"&",
  AT             :"@",
  BSLASH         :"\\",
  CARAT          :"^",
  COLON          :":",
  COMMA          :",",
  DIVIDE         :"/",
  DOLLAR         :"$",
  EXCLAM         :"!",
  EQUAL          :"=",
  GTHAN          :">",
  LBRACE         :"{",
  LBRACK         :"[",
  LPAREN         :"(",
  LTHAN          :"<",
  MINUS          :"-",
  PERIOD         :".",
  PERCENT        :"%",
  PLUS           :"+",
  POUND          :"#",
  QMARK          :"?",
  RBRACE         :"}",
  RBRACK         :"]",
  RPAREN         :")",
  SEMI           :";",
  DQUOTE         :"\"",
  UNDERSCORE     :"_",
  VBAR           :"|",
  NEWLINE        :"\n",
  CR             :"\r",
  TAB            :"\t",
  SPACE          :" ",
  SQUOTE         :"'",
  CURSOR         :"|",
}

var Names []string = []string {
  "pure",
  "type",
  "interface",
  "op",
  "if",
  "else",
  "defer",
  "when",
  "is",
  "for",
  "select",
  "shared",
  "module",
  "map",
  "list",
  "tree",
  "trie",
  "func",
  "chan",
  "mat",
  "vec",
  "uint",
  "int",
  "num",
  "float",
  "bool",
  "byte",
  "text",
  "error",
  "exception",
  "unit",
  "identifier",
  "binary literal",
  "hex literal",
  "octal literal",
  "numeric literal",
  "string literal",
  ":=",
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
  "'",
  "|"}

type TokenClass int
const (
  PURE TokenClass = iota
  TYPE 
  INTERFACE
  OP
  IF
  ELSE
  DEFER
  WHEN
  IS
  FOR
  SELECT
  SHARED  
  MODULE
  MAP_TYPE
  LIST_TYPE
  TREE_TYPE
  TRIE_TYPE
  FUNC_TYPE
  CHAN_TYPE
  MAT_TYPE
  VEC_TYPE  
  UINT_TYPE
  INT_TYPE
  NUM_TYPE
  FLOAT_TYPE
  BOOL_TYPE
  BYTE_TYPE
  TEXT_TYPE
  ERROR_TYPE
  EXCEPTION_TYPE  
  UNIT_TYPE
  IDENTIFIER
  BINARY_LIT
  HEX_LIT
  OCTAL_LIT
  NUMBER_LIT
  STRING_LIT
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
  CURSOR
)
