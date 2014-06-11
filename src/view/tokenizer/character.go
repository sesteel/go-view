package tokenizer

import (
	. "view/common"
)

// Character combines and stores information used to render 
// tokenized text on the screen. 
type Character struct {

	// Position this character is in when in a list of lines
	Index int

	// Token is the token this character belongs to.
	Token *Token

	// Rune used to render this character.
	Rune rune

	// Bounds stores the location, if any, where this 
	// character is rendered.
	Bounds Bounds
}
