package tokenizer

import ()

func ToLines(tkns []*Token) [][]*Token {
	lines := make([][]*Token, 0, 0)
	line := make([]*Token, 0)
	lines = append(lines, line)
	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]
		line = append(line, tkn)
		if tkn.Type == NEWLINE {
			line = make([]*Token, 0)
			lines = append(lines, line)
		}
	}
	return lines
}
