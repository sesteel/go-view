// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

package view

import (
	. "view/event/key"
)

var keymap map[uint]Key

func init() {
	keymap = make(map[uint]Key)
	keymap[0xFF08] = BACKSPACE
	keymap[0x20]   = SPACE
	keymap[0xFF09] = TAB
	keymap[0xFF0D] = RETURN
	keymap[0xFF13] = BREAK
	keymap[0xFF14] = SCROLL_LOCK
	keymap[0xFF1B] = ESC
	keymap[0xFFE5] = CAPS
	keymap[0xFFE1] = LEFT_SHIFT
	keymap[0xFFE2] = RIGHT_SHIFT
	keymap[0xFFE3] = LEFT_CTRL
	keymap[0xFFE4] = RIGHT_CTRL
	keymap[0xFFE9] = LEFT_ALT
	keymap[0xFFEA] = RIGHT_ALT
	keymap[0xFFEB] = LEFT_CMD
	keymap[0xFFEC] = RIGHT_CMD

	keymap[0xFF50] = HOME
	keymap[0xFF51] = ARROW_LEFT
	keymap[0xFF52] = ARROW_UP
	keymap[0xFF53] = ARROW_RIGHT
	keymap[0xFF54] = ARROW_DOWN
	keymap[0xFF55] = PAGE_UP
	keymap[0xFF56] = PAGE_DOWN
	keymap[0xFF57] = END
	keymap[0xFFFF] = DELETE

	keymap[0xFF7F] = NUM_LOCK
	keymap[0xFFBD] = NUM_EQUALS
	keymap[0xFFAF] = NUM_DIVIDE
	keymap[0xFFAA] = NUM_MULTIPLY
	keymap[0xFFAD] = NUM_SUBTRACT
	keymap[0xFFAB] = NUM_ADD
	keymap[0xFF8D] = NUM_ENTER
	keymap[0xFF9F] = NUM_DECIMAL
	keymap[0xFF9E] = NUM_ZERO
	keymap[0xFF9C] = NUM_ONE
	keymap[0xFF99] = NUM_TWO
	keymap[0xFF9B] = NUM_THREE
	keymap[0xFF96] = NUM_FOUR
	keymap[0xFF9D] = NUM_FIVE
	keymap[0xFF98] = NUM_SIX
	keymap[0xFF95] = NUM_SEVEN
	keymap[0xFF97] = NUM_EIGHT
	keymap[0xFF9A] = NUM_NINE

	keymap[0xFFBE] = F1
	keymap[0xFFBF] = F2
	keymap[0xFFC0] = F3
	keymap[0xFFC1] = F4
	keymap[0xFFC2] = F5
	keymap[0xFFC3] = F6
	keymap[0xFFC4] = F7
	keymap[0xFFC5] = F8
	keymap[0xFFC6] = F9
	keymap[0xFFC7] = F10
	keymap[0xFFC8] = F11
	keymap[0xFFC9] = F12

	keymap[0x1008ff81] = F13
	keymap[0x1008ff45] = F14
	keymap[0x1008ff46] = F15
	keymap[0x1008ff47] = F16
	keymap[0x1008ff48] = F17
	keymap[0x1008ff49] = F18

	keymap[0x1008ff18] = WEB_HOME
	keymap[0x1008ff26] = WEB_BACK
	keymap[0x1008ff27] = WEB_FORWARD
	keymap[0x1008ff30] = WEB_BOOKMARK

	keymap[0x61] = A
	keymap[0x62] = B
	keymap[0x63] = C
	keymap[0x64] = D
	keymap[0x65] = E
	keymap[0x66] = F
	keymap[0x67] = G
	keymap[0x68] = H
	keymap[0x69] = I
	keymap[0x6A] = J
	keymap[0x6B] = K
	keymap[0x6C] = L
	keymap[0x6D] = M
	keymap[0x6E] = N
	keymap[0x6F] = O
	keymap[0x70] = P
	keymap[0x71] = Q
	keymap[0x72] = R
	keymap[0x73] = S
	keymap[0x74] = T
	keymap[0x75] = U
	keymap[0x76] = V
	keymap[0x77] = W
	keymap[0x78] = X
	keymap[0x79] = Y
	keymap[0x7A] = Z

	keymap[0x30] = ZERO
	keymap[0x31] = ONE
	keymap[0x32] = TWO
	keymap[0x33] = THREE
	keymap[0x34] = FOUR
	keymap[0x35] = FIVE
	keymap[0x36] = SIX
	keymap[0x37] = SEVEN
	keymap[0x38] = EIGHT
	keymap[0x39] = NINE
	keymap[0x60] = GRAVE_ACCENT
	keymap[0x2D] = MINUS_SIGN
	keymap[0x3D] = EQUAL_SIGN
	keymap[0x5B] = OPEN_BRACKET
	keymap[0x5D] = CLOSED_BRACKET
	keymap[0x5C] = BLACKSLASH
	keymap[0x3B] = SEMICOLON
	keymap[0x27] = SINGLE_QUOTE
	keymap[0x2C] = COMMA
	keymap[0x2E] = PERIOD
	keymap[0x2F] = SLASH

}
