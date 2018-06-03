package text

import (
	"github.com/sesteel/go-view"
)

var charmasks map[view.Font]map[int]*view.Surface

func init() {
	charmasks = make(map[view.Font]map[int]*view.Surface)
}
