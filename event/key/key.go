// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package key

import ()

type Key int32

const (
	// here, we abuse the use of some unicode
	// characters to maintain simplicity
	NONE Key = iota
	A
    B
    C
    D
    E
    F
    G
    H
    I
    J
    K
    L
    M
    N
    O
    P
    Q
    R
    S
    T
    U
    V
    W
    X
    Y
    Z
    ZERO
    ONE
    TWO
    THREE
    FOUR
    FIVE
    SIX
    SEVEN
    EIGHT
    NINE
    
    GRAVE_ACCENT
	MINUS_SIGN
	PLUS_SIGN
	EQUAL_SIGN
	OPEN_BRACKET
	CLOSED_BRACKET
	BLACKSLASH
	SEMICOLON
	SINGLE_QUOTE
	COMMA
	PERIOD
	SLASH
	
	SPACE
    BACKSPACE
	TAB
	RETURN
	BREAK
	SCROLL_LOCK
	ESC
	CAPS
	LEFT_SHIFT
	RIGHT_SHIFT
	LEFT_CTRL
	RIGHT_CTRL
	LEFT_ALT
	RIGHT_ALT
	LEFT_CMD
	RIGHT_CMD	
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	ARROW_LEFT
    ARROW_UP
    ARROW_RIGHT
    ARROW_DOWN
    HOME
    END
    PAGE_UP
    PAGE_DOWN
    DELETE
    NUM_LOCK
    NUM_EQUALS
    NUM_DIVIDE
    NUM_MULTIPLY
    NUM_SUBTRACT
    NUM_ADD
    NUM_ENTER
    NUM_DECIMAL
    NUM_ZERO
    NUM_ONE
    NUM_TWO
    NUM_THREE
    NUM_FOUR
    NUM_FIVE
    NUM_SIX
    NUM_SEVEN
    NUM_EIGHT
    NUM_NINE
    
    WEB_HOME
    WEB_BACK
    WEB_FORWARD
    WEB_BOOKMARK
	
)

func (self Key) NumberKey() bool {
	switch self {
		case NUM_ZERO,
    		 NUM_ONE,
             NUM_TWO,
             NUM_THREE,
             NUM_FOUR,
             NUM_FIVE,
             NUM_SIX,
             NUM_SEVEN,
             NUM_EIGHT,
             NUM_NINE:
			return true
		default:
			return false
	}
}
