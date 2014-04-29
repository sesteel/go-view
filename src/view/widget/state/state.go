package state

import (

)

type State uint8

const (
	NORMAL State = 1 << iota
	HOVER 
	FOCUS
	DISABLED
	ERROR
	ACTIVATED
)