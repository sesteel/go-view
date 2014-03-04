// +build linux,!goci
package core

import (

)

type Window interface {
	// 
	GetSurface() *Surface
	
	// 
	SetName(string)
	
	// 
	GetName() string
	
	// 
	SetSize(uint, uint)
	
	// 
	GetSize() (uint, uint)
}
