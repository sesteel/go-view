package tokenizer

import (

)

type Tokenizer interface {
	Tokenize(text string) []*Token
}
