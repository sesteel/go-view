package tokenizer

import ()

// Character combines and stores information used to render
// tokenized text on the screen.
type Character struct {

	// Position this character is in when in a list of lines.
	Index int

	// Token is the token type this character belongs to.
	Token *Token

	// Rune which represents this character.
	Rune rune
}
