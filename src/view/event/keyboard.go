package event

import (
	. "view/event/key"
)

var GlobalKeyHandler KeyboardHandler
var keyboard Keyboard

type Keyboard struct {
	Value        Key
	LeftCtrl     bool
	RightCtrl    bool
	LeftAlt      bool
	RightAlt     bool
	LeftShift    bool
	RightShift   bool
	LeftCommand  bool
	RightCommand bool
	Caps         bool
	NumLock      bool
}

func (self *Keyboard) Ctrl() bool    { return self.LeftCtrl || self.RightCtrl }
func (self *Keyboard) Alt() bool     { return self.LeftAlt || self.RightAlt }
func (self *Keyboard) Shift() bool   { return self.LeftShift || self.RightShift }
func (self *Keyboard) Command() bool { return self.LeftCommand || self.RightCommand }

func (self *Keyboard) toggleModifiersOn(key Key) {
	switch key {
	case NUM_LOCK:
		self.NumLock = !self.NumLock
	case CAPS:
		self.Caps = !self.Caps
	case LEFT_CTRL:
		self.LeftCtrl = true
	case LEFT_ALT:
		self.LeftAlt = true
	case LEFT_SHIFT:
		self.LeftShift = true
	case LEFT_CMD:
		self.LeftCommand = true
	case RIGHT_CTRL:
		self.RightCtrl = true
	case RIGHT_ALT:
		self.RightAlt = true
	case RIGHT_SHIFT:
		self.RightShift = true
	case RIGHT_CMD:
		self.RightCommand = true
	default:
		self.Value = key
	}
}

func (self *Keyboard) toggleModifiersOff(key Key) {
	switch key {
	case LEFT_CTRL:
		self.LeftCtrl = false
	case LEFT_ALT:
		self.LeftAlt = false
	case LEFT_SHIFT:
		self.LeftShift = false
	case LEFT_CMD:
		self.LeftCommand = false
	case RIGHT_CTRL:
		self.RightCtrl = false
	case RIGHT_ALT:
		self.RightAlt = false
	case RIGHT_SHIFT:
		self.RightShift = false
	case RIGHT_CMD:
		self.RightCommand = false
	default:
		self.Value = key
	}
}


func (self Keyboard) Rune() rune {
	rm := runeMap[self.Value]
	if rm == nil {
		return rune(0)
	}
	
	if self.Caps {
		if self.Shift() && !self.Value.NumberKey() {
			return rm.a
		} else {
			return rm.b
		}
	} 
	
	if self.Shift() {
		return rm.b
	} else {
		return rm.a
	}
}

func (self Keyboard) String() string {
	v := self.Rune()
	if v == 0 {
		return ""
	} else {
		return string(v)
	}
}

type KeyboardNotifier interface {
	AddKeyPressHandler(func(Keyboard))
	AddKeyReleaseHandler(func(Keyboard))
}

type KeyboardHandler interface {
	KeyPress(state Keyboard)
	KeyRelease(state Keyboard)
}

type KeyboardEventDispatcher struct {
	keyPressHandlers   []func(Keyboard)
	keyReleaseHandlers []func(Keyboard)
}

func (self *KeyboardEventDispatcher) AddKeyPressHandler(f func(Keyboard)) {
	self.keyPressHandlers = append(self.keyPressHandlers, f)
}

func (self *KeyboardEventDispatcher) RemoveFocusGainedHandler(f func(Keyboard)) {
	// TODO - implement
}

func (self *KeyboardEventDispatcher) KeyPress(keyboard Keyboard) {
	for i := 0; i < len(self.keyPressHandlers); i++ {
		self.keyPressHandlers[i](keyboard)
	}
}

func (self *KeyboardEventDispatcher) AddKeyReleaseHandler(f func(Keyboard)) {
	self.keyReleaseHandlers = append(self.keyReleaseHandlers, f)
}

func (self *KeyboardEventDispatcher) RemoveKeyReleaseHandler(f func(Keyboard)) {
	// TODO - implement
}

func (self *KeyboardEventDispatcher) KeyRelease(keyboard Keyboard) {
	for i := 0; i < len(self.keyReleaseHandlers); i++ {
		self.keyReleaseHandlers[i](keyboard)
	}
}

func DispatchKeyPress(key Key) {
	keyboard.toggleModifiersOn(key)

	if GlobalKeyHandler != nil {
		GlobalKeyHandler.KeyPress(keyboard)
	}

	if kh, ok := focussedElement.(KeyboardHandler); ok {
		kh.KeyPress(keyboard)
	} else if kh, ok := focussedElement.(EventDispatcher); ok {
		kh.KeyPress(keyboard)
	}
}

func DispatchKeyRelease(key Key) {
	keyboard.toggleModifiersOff(key)

	if GlobalKeyHandler != nil {
		GlobalKeyHandler.KeyRelease(keyboard)
	}

	if kh, ok := focussedElement.(KeyboardHandler); ok {
		kh.KeyRelease(keyboard)
	} else if kh, ok := focussedElement.(EventDispatcher); ok {
		kh.KeyRelease(keyboard)
	}
}

