package tokenizer

import (
	. "view/common"
)

// Character combines and stores information used to render
// tokenized text on the screen.
type Character struct {

	// Position this character is in when in a list of lines.
	Index int

	// Token is the token type this character belongs to.
	Token *Token

	// Rune which represents this character.
	Rune rune

	// Bounds stores the location, if any, where this
	// character is rendered.  This is for the convienence
	// of other packages which may need a simple way to
	// store this information.  It is nil by default.
	Bounds *Bounds
}
