package text

import (
	"github.com/sesteel/go-view"
	"github.com/sesteel/go-view/color"
	"github.com/sesteel/go-view/geometry"
	"github.com/sesteel/go-view/tokenizer"
)

// Character combines and stores information used to render
// tokenized text on the screen.
type Character struct {

	// Provides the styling for the token
	styler styler

	// Position this character is in when in a list of lines.
	Index int

	// Token is the token type this character belongs to.
	Token *tokenizer.Token

	// Rune which represents this character.
	Rune rune

	// The space occupied by this character.
	Bounds geometry.Bounds
}

func (self *Character) Draw(s *view.Surface) {
	s.SetSourceRGBA(color.Black)
	s.DrawRune(self.Rune, 100, 100)
}
