package event

import (
	. "view/event/key"
)

var GlobalKeyHandler KeyboardHandler
var keyboard Keyboard

type Keyboard struct {
	Char         rune
	LeftCtrl     bool
	RightCtrl    bool
	LeftAlt      bool
	RightAlt     bool
	LeftShift    bool
	RightShift   bool
	LeftCommand  bool
	RightCommand bool
}

func (self *Keyboard) Ctrl() bool    { return self.LeftCtrl || self.RightCtrl }
func (self *Keyboard) Alt() bool     { return self.LeftAlt || self.RightAlt }
func (self *Keyboard) Shift() bool   { return self.LeftShift || self.RightShift }
func (self *Keyboard) Command() bool { return self.LeftCommand || self.RightCommand }

func (self *Keyboard) toggleModifiers(key Key, on bool) {
	switch key {
	case LEFT_CTRL:
		self.LeftCtrl = on
	case LEFT_ALT:
		self.LeftAlt = on
	case LEFT_SHIFT:
		self.LeftShift = on
	case LEFT_CMD:
		self.LeftCommand = on
	case RIGHT_CTRL:
		self.RightCtrl = on
	case RIGHT_ALT:
		self.RightAlt = on
	case RIGHT_SHIFT:
		self.RightShift = on
	case RIGHT_CMD:
		self.RightCommand = on
	}
}

func (self Keyboard) String() string {
	return "a"
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
	keyboard.toggleModifiers(key, true)

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
	keyboard.toggleModifiers(key, false)

	if GlobalKeyHandler != nil {
		GlobalKeyHandler.KeyRelease(keyboard)
	}

	if kh, ok := focussedElement.(KeyboardHandler); ok {
		kh.KeyRelease(keyboard)
	} else if kh, ok := focussedElement.(EventDispatcher); ok {
		kh.KeyRelease(keyboard)
	}
}

