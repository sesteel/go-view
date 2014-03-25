package tokenizer

import (
	"fmt"
)

type Token struct {
  Type     int
  Value    string
  Code     string
  Line     int
  Start    int 
  End      int
  Selected bool
}

func (t *Token) String() string {
   return fmt.Sprintf("{%s \"%s\" %d %d %d}", 
   	 Names[t.Type], t.Value, t.Line, t.Start, t.End)
}

