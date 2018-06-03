// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel
package event

import (
	. "github.com/sesteel/go-view/event/key"
)

var runeMap map[Key]*keyMap

type keyMap struct {
	a, b rune
}

func init() {
	runeMap = make(map[Key]*keyMap)

	// TODO: Add 'Some' International Support
	runeMap[A] = &keyMap{'a', 'A'}
	runeMap[B] = &keyMap{'b', 'B'}
	runeMap[C] = &keyMap{'c', 'C'}
	runeMap[D] = &keyMap{'d', 'D'}
	runeMap[E] = &keyMap{'e', 'E'}
	runeMap[F] = &keyMap{'f', 'F'}
	runeMap[G] = &keyMap{'g', 'G'}
	runeMap[H] = &keyMap{'h', 'H'}
	runeMap[I] = &keyMap{'i', 'I'}
	runeMap[J] = &keyMap{'j', 'J'}
	runeMap[K] = &keyMap{'k', 'K'}
	runeMap[L] = &keyMap{'l', 'L'}
	runeMap[M] = &keyMap{'m', 'M'}
	runeMap[N] = &keyMap{'n', 'N'}
	runeMap[O] = &keyMap{'o', 'O'}
	runeMap[P] = &keyMap{'p', 'P'}
	runeMap[Q] = &keyMap{'q', 'Q'}
	runeMap[R] = &keyMap{'r', 'R'}
	runeMap[S] = &keyMap{'s', 'S'}
	runeMap[T] = &keyMap{'t', 'T'}
	runeMap[U] = &keyMap{'u', 'U'}
	runeMap[V] = &keyMap{'v', 'V'}
	runeMap[W] = &keyMap{'w', 'W'}
	runeMap[X] = &keyMap{'x', 'X'}
	runeMap[Y] = &keyMap{'y', 'Y'}
	runeMap[Z] = &keyMap{'z', 'Z'}
	runeMap[GRAVE_ACCENT] = &keyMap{'`', '~'}
	runeMap[ONE] = &keyMap{'1', '!'}
	runeMap[TWO] = &keyMap{'2', '@'}
	runeMap[THREE] = &keyMap{'3', '#'}
	runeMap[FOUR] = &keyMap{'4', '$'}
	runeMap[FIVE] = &keyMap{'5', '%'}
	runeMap[SIX] = &keyMap{'6', '^'}
	runeMap[SEVEN] = &keyMap{'7', '&'}
	runeMap[EIGHT] = &keyMap{'8', '*'}
	runeMap[NINE] = &keyMap{'9', '('}
	runeMap[ZERO] = &keyMap{'0', ')'}
	runeMap[MINUS_SIGN] = &keyMap{'-', '_'}
	runeMap[EQUAL_SIGN] = &keyMap{'=', '+'}
	runeMap[OPEN_BRACKET] = &keyMap{'[', '{'}
	runeMap[CLOSED_BRACKET] = &keyMap{']', '}'}
	runeMap[BLACKSLASH] = &keyMap{'\\', '|'}
	runeMap[SEMICOLON] = &keyMap{';', ':'}
	runeMap[SINGLE_QUOTE] = &keyMap{'\'', '"'}
	runeMap[COMMA] = &keyMap{',', '<'}
	runeMap[PERIOD] = &keyMap{'.', '>'}
	runeMap[SLASH] = &keyMap{'/', '?'}
	runeMap[NUM_EQUALS] = &keyMap{'=', '='}
	runeMap[NUM_DIVIDE] = &keyMap{'/', '/'}
	runeMap[NUM_MULTIPLY] = &keyMap{'*', '*'}
	runeMap[NUM_SUBTRACT] = &keyMap{'-', '-'}
	runeMap[NUM_ADD] = &keyMap{'+', '+'}
	runeMap[SPACE] = &keyMap{' ', ' '}
	runeMap[TAB] = &keyMap{'\t', '\t'}

	// TODO : Fix Numlock values
	runeMap[NUM_ADD] = &keyMap{'+', '+'}
	runeMap[NUM_DECIMAL] = &keyMap{'.', '.'}
	runeMap[NUM_ZERO] = &keyMap{'0', '0'}
	runeMap[NUM_ONE] = &keyMap{'1', '1'}
	runeMap[NUM_TWO] = &keyMap{'2', '2'}
	runeMap[NUM_THREE] = &keyMap{'3', '3'}
	runeMap[NUM_FOUR] = &keyMap{'4', '4'}
	runeMap[NUM_FIVE] = &keyMap{'5', '5'}
	runeMap[NUM_SIX] = &keyMap{'6', '6'}
	runeMap[NUM_SEVEN] = &keyMap{'7', '7'}
	runeMap[NUM_EIGHT] = &keyMap{'8', '8'}
	runeMap[NUM_NINE] = &keyMap{'9', '9'}
}
