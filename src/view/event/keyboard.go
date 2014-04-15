package event

import (
	. "view/event/key"
)

var GlobalKeyHandler KeyboardHandler

func DispatchKeyPress(code Key) {
	keyboard.key = code
	if GlobalKeyHandler != nil {
		GlobalKeyHandler.KeyPress(keyboard)
	}
	
	kh, ok := focussedElement.(KeyboardHandler)
	if ok {
		kh.KeyPress(keyboard)
	}
}

var keyboard defaultKeyboard
type Keyboard interface {
	Key() Key
	Ctrl() bool
	Alt() bool
	Shift() bool
	Caps() bool
	Command() bool
}

type defaultKeyboard struct {
	key     Key
	ctrl    bool
	alt     bool
	shift   bool
	caps    bool
	command bool
}

func (self defaultKeyboard) Key() Key {return self.key}
func (self defaultKeyboard) Ctrl() bool {return self.ctrl}
func (self defaultKeyboard) Alt() bool {return self.alt}
func (self defaultKeyboard) Shift() bool {return self.shift}
func (self defaultKeyboard) Caps() bool {return self.caps}
func (self defaultKeyboard) Command() bool {return self.command}

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

func (self *KeyboardEventDispatcher) KeyPress(key Key) {
	keyboard.key = key
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

func (self *KeyboardEventDispatcher) KeyRelease(key Key) {
	keyboard.key = key
	for i := 0; i < len(self.keyReleaseHandlers); i++ {
		self.keyReleaseHandlers[i](keyboard)
	}
}
