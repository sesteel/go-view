// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright (c) 2014 Stanley Steel

// +build linux,!goci
package view

// #cgo pkg-config: cairo x11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xresource.h>
// #include <X11/keysymdef.h>
// #include <cairo/cairo-xlib.h>
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	//	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
	"unsafe"
	"view/color"
	"view/event"
)

const (
	_DELETE_WINDOW = "WM_DELETE_WINDOW"
)

func init() {
	C.XInitThreads()
}

type Window struct {
	CompositeView
	display     *C.Display
	screen      C.int
	xwindow     C.Window
	drawCounter uint
	eventLoop   func()
	drawloop    func()
	dirty       bool
	surface     *Surface
	width       float64
	height      float64
}

func (self *Window) Parent() View {
	return nil
}

func (self *Window) Start() {
	go self.eventLoop()
	go self.drawloop()
}

func (self *Window) SetLayout(layout Layout) {
	self.layout = layout
}

func (self *Window) GetLayout() Layout {
	return self.layout
}

func (self *Window) Surface() *Surface {
	return self.surface
}

func (self *Window) SetName(name string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	self.name = name
	C.XStoreName(self.display, self.xwindow, n)
}

func (self *Window) SetSize(width, height float64) {
	self.width = width
	self.height = height
	C.cairo_xlib_surface_set_size(self.surface.surface, C.int(self.width), C.int(self.height))
	C.cairo_xlib_surface_set_drawable(self.surface.surface, C.Drawable(self.xwindow), C.int(self.width), C.int(self.height))
	C.XResizeWindow(self.display, self.xwindow, C.uint(self.width), C.uint(self.height))
	self.Redraw()
}

func (self *Window) Draw(surface *Surface) {
	//	surface.SetSourceRGBA(self.Style().Background())
	//	_, h := self.Size()
	//	p := NewLinearPattern(0, 0, 0, h)
	//	defer p.Destroy()
	//	p.AddColorStop(0, color.Gray3)
	//	p.AddColorStop(1, color.Gray5)
	//	surface.SetSource(p)
	surface.SetSourceRGBA(color.Gray3)
	surface.Paint()

	// tiled alpha background
	//	tile := NewSurfaceFromPNG("res/textures/concrete.png")
	//	defer tile.Destroy()
	//	pattern1 := C.cairo_pattern_create_for_surface(tile.surface)
	//	C.cairo_set_source(surface.context, pattern1)
	//  	C.cairo_pattern_set_extend(C.cairo_get_source(surface.context), C.cairo_extend_t(EXTEND_REPEAT))
	//	surface.Paint()

	if self.layout != nil {
		self.layout.Draw(surface)
	}
}

func (self *Window) Redraw() {
	self.dirty = true
}

func (self *Window) Animate(s *Surface) {
	self.layout.Animate(s)
}

// func NewBorderlessWindow(name string, x, y, w, h uint) *Window {
// 	window := NewWindow(name, x, y, w, h)

// 	// var mwmhints C.Atom
// 	// var hints C.MwmHints
// 	//// #include <Xll/X.h>
// 	//// #include <X11/Xm/MwmUtil.h>
// 	// internal atom that we are looking for
// 	mwmhints := C.XInternAtom(window.display, C._XA_MWM_HINTS, C.False)
// 	// hints = (MwmHints *)malloc(sizeof(MwmHints));
// 	hints.decorations = 0
// 	hints.flags |= MWM_HINTS_DECORATIONS
// 	C.XChangeProperty(window.display, window.xwindow, mwmhints, mwmhints, 32, PropModeReplace, hints, 4)
// 	C.XFlush(window.display)
// 	return window
// }

func NewWindow(name string, x, y, w, h uint) *Window {
	var width C.uint = C.uint(w)
	var height C.uint = C.uint(h)
	var ev C.XEvent

	// First connect to the display server, as specified in
	// the DISPLAY environment variable.
	dpy := C.XOpenDisplay(nil)

	if dpy == nil {
		fmt.Println("unable to connect to display")
		return nil
	}

	if DEBUG_XVISUAL_INFO {
		var visual_template C.XVisualInfo
		var nxvisuals C.int
		visual_list := C.XGetVisualInfo(dpy, C.VisualScreenMask, &visual_template, &nxvisuals)
		var visualList []C.XVisualInfo
		sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&visualList)))
		sliceHeader.Cap = int(nxvisuals)
		sliceHeader.Len = int(nxvisuals)
		sliceHeader.Data = uintptr(unsafe.Pointer(&visual_list))

		for i := 0; i < len(visualList); i++ {
			if uint(visualList[i].depth) > 8 {
				fmt.Printf("  %d: visual:%d class:%d TrueColor:%t depth:%d\n", i, int(visualList[i].visualid), visualList[i].class, bool(visualList[i].class == C.TrueColor), uint(visualList[i].depth))
			}
		}
	}

	var vinfo C.XVisualInfo
	result := C.XMatchVisualInfo(dpy, C.XDefaultScreen(dpy), 24, C.TrueColor, &vinfo)
	if result == 0 {
		fmt.Println("Cannot create display at desired depth of 24.")
	}

	var attr C.XSetWindowAttributes
	attr.colormap = C.XCreateColormap(dpy, C.XDefaultRootWindow(dpy), vinfo.visual, C.AllocNone)
	attr.border_pixel = 0xFFFF00FF
	attr.background_pixel = 0xFFFF00FF
	attr.event_mask = C.ExposureMask |
		C.ButtonPressMask |
		C.ButtonReleaseMask |
		C.KeyPressMask |
		C.KeyReleaseMask |
		C.PointerMotionMask |
		C.EnterWindowMask |
		C.LeaveWindowMask |
		C.StructureNotifyMask

	/* these are macros that pull useful data out of the display object */
	/* we use these bits of info enough to want them in their own variables */
	screen_num := C.XDefaultScreen(dpy)

	win := C.XCreateWindow(dpy, C.XDefaultRootWindow(dpy), 0, 0, width, height, 0, vinfo.depth, C.InputOutput, vinfo.visual, C.CWColormap|C.CWBorderPixel, &attr)

	C.XSync(dpy, C.True)

	/* tell the display server what kind of events we would like to see */
	C.XSelectInput(dpy, win, C.ExposureMask|
		C.ButtonPressMask|
		C.ButtonReleaseMask|
		C.KeyPressMask|
		C.KeyReleaseMask|
		C.PointerMotionMask|
		C.EnterWindowMask|
		C.LeaveWindowMask|
		C.StructureNotifyMask)

	/* okay, put the window on the screen, please */
	C.XMapWindow(dpy, win)

	delwin := C.CString(_DELETE_WINDOW)
	defer C.free(unsafe.Pointer(delwin))
	WM_DELETE_WINDOW := C.XInternAtom(dpy, delwin, C.False)
	C.XSetWMProtocols(dpy, win, &WM_DELETE_WINDOW, 1)

	s := C.cairo_xlib_surface_create(dpy, C.Drawable(win), C.XDefaultVisual(dpy, 0), C.int(width), C.int(height))
	ctx := C.cairo_create(s)
	surface := &Surface{surface: s, context: ctx}

	window := new(Window)
	window.display = dpy
	window.screen = screen_num
	window.xwindow = win
	window.surface = surface
	window.name = name
	window.width = float64(width)
	window.height = float64(height)
	window.layout = nil
	window.style = NewStyle()

	window.SetName(name)
	runtime.SetFinalizer(window, func(w *Window) {
		C.XCloseDisplay(window.display)
	})

	window.Redraw()

	eventLoop := func() {

		left := false
		middle := false
		right := false

		/* as each event that we asked about occurs, we respond.  In this
		 * case we note if the window's shape changed, and exit if a button
		 * is pressed inside the window */
		for {
			C.XNextEvent(dpy, &ev)
			eventType := ev[0]
			switch eventType {
			case C.ConfigureNotify:
				evt := (*C.XConfigureEvent)(unsafe.Pointer(&ev[0]))
				if width != C.uint(evt.width) || height != C.uint(evt.height) {
					width = C.uint(evt.width)
					height = C.uint(evt.height)
					window.width = float64(width)
					window.height = float64(height)
					C.cairo_xlib_surface_set_size(s, evt.width, evt.height)
				}

			case C.Expose:
				evt := (*C.XExposeEvent)(unsafe.Pointer(&ev[0]))
				if evt.count > 0 {
					continue
				}
				window.Redraw()

			case C.MotionNotify:
				evt := (*C.XMotionEvent)(unsafe.Pointer(&ev[0]))
				window.MousePosition(event.Mouse{event.MOUSE_BUTTON_NONE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

			case C.EnterNotify:
				evt := (*C.XCrossingEvent)(unsafe.Pointer(&ev[0]))
				window.MouseEnter(event.Mouse{event.MOUSE_BUTTON_NONE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

			case C.LeaveNotify:
				evt := (*C.XCrossingEvent)(unsafe.Pointer(&ev[0]))
				window.MouseExit(event.Mouse{event.MOUSE_BUTTON_NONE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

			case C.ButtonPress:
				evt := (*C.XButtonEvent)(unsafe.Pointer(&ev[0]))

				switch evt.button {
				case 1:
					left = true
					window.MouseButtonPress(event.Mouse{event.MOUSE_BUTTON_LEFT, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

				case 2:
					middle = true
					window.MouseButtonPress(event.Mouse{event.MOUSE_BUTTON_MIDDLE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

				case 3:
					right = true
					window.MouseButtonPress(event.Mouse{event.MOUSE_BUTTON_RIGHT, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

				case 4:
					window.MouseWheelUp(event.Mouse{event.MOUSE_BUTTON_NONE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})

				case 5:
					window.MouseWheelDown(event.Mouse{event.MOUSE_BUTTON_NONE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})
				}

			case C.ButtonRelease:
				evt := (*C.XButtonEvent)(unsafe.Pointer(&ev[0]))
				switch evt.button {
				case 1:
					left = false
					window.MouseButtonRelease(event.Mouse{event.MOUSE_BUTTON_LEFT, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})
				case 2:
					middle = false
					window.MouseButtonRelease(event.Mouse{event.MOUSE_BUTTON_MIDDLE, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})
				case 3:
					right = false
					window.MouseButtonRelease(event.Mouse{event.MOUSE_BUTTON_RIGHT, event.MouseState{left, middle, right, float64(evt.x), float64(evt.y)}})
				}

			case C.KeyPress:
				evt := (*C.XKeyEvent)(unsafe.Pointer(&ev[0]))
				var keysyms_per_keycode_return C.int
				keysym := C.XGetKeyboardMapping(dpy, (C.KeyCode)(evt.keycode), 1, &keysyms_per_keycode_return)
				defer C.XFree(unsafe.Pointer(&keysym))
				symbol := uint(*keysym)
				event.DispatchKeyPress(keymap[symbol])
				//				fmt.Printf("[ %x ] %x\n", *keysym, evt.keycode)

			case C.KeyRelease:
				evt := (*C.XKeyEvent)(unsafe.Pointer(&ev[0]))
				var keysyms_per_keycode_return C.int
				keysym := C.XGetKeyboardMapping(dpy, (C.KeyCode)(evt.keycode), 1, &keysyms_per_keycode_return)
				defer C.XFree(unsafe.Pointer(&keysym))
				symbol := uint(*keysym)
				event.DispatchKeyRelease(keymap[symbol])
				//				fmt.Printf("[ %x ] %x\n", *keysym, evt.keycode)

			case C.ClientMessage:
				C.XCloseDisplay(dpy)
				os.Exit(0)

			default:
				C.XFlush(dpy)
			}
		}
	}
	window.eventLoop = eventLoop

	drawloop := func() {
		var before time.Time
		count := 0
		if DEBUG_DRAW_ALL {
			before = time.Now()
		}

		s1 := NewSurface(FORMAT_ARGB32, int(window.width), int(window.height))
		for {
			since := time.Now()
			if window.dirty {
				window.dirty = false
				s1.Destroy()
				s1 = NewSurface(FORMAT_ARGB32, int(window.width), int(window.height))
				window.Draw(s1)
				fmt.Println("time to render:", time.Since(since))
			}

			s2 := NewSurface(FORMAT_ARGB32, int(window.width), int(window.height))
			s2.SetSourceSurface(s1, 0, 0)
			s2.Paint()
			window.Animate(s2)
			s2.Flush()
			window.surface.SetSourceSurface(s2, 0, 0)
			window.surface.Paint()
			window.surface.Flush()

			s2.Destroy()

			if DEBUG_DRAW_ALL {
				if time.Since(before).Seconds() >= 1 {
					if count < 60 {
						fmt.Println("FPS:", count)
					}
					count = 0
					before = time.Now()
				} else {
					count++
				}
			}

			C.XFlush(window.display)
			time.Sleep(time.Millisecond * 16)
		}
	}
	window.drawloop = drawloop

	return window
}

// NoEventMask				No events wanted
// KeyPressMask				Keyboard down events wanted
// KeyReleaseMask			Keyboard up events wanted
// ButtonPressMask			Pointer button down events wanted
// ButtonReleaseMask		Pointer button up events wanted
// EnterWindowMask			Pointer window entry events wanted
// LeaveWindowMask			Pointer window leave events wanted
// PointerMotionMask		Pointer motion events wanted
// PointerMotionHintMask	Pointer motion hints wanted
// Button1MotionMask		Pointer motion while button 1 down
// Button2MotionMask		Pointer motion while button 2 down
// Button3MotionMask		Pointer motion while button 3 down
// Button4MotionMask		Pointer motion while button 4 down
// Button5MotionMask		Pointer motion while button 5 down
// ButtonMotionMask			Pointer motion while any button down
// KeymapStateMask			Keyboard state wanted at window entry and focus in
// ExposureMask				Any exposure wanted
// VisibilityChangeMask		Any change in visibility wanted
// StructureNotifyMask		Any change in window structure wanted
// ResizeRedirectMask		Redirect resize of this window
// SubstructureNotifyMask	Substructure notification wanted
// SubstructureRedirectMask Redirect structure requests on children
// FocusChangeMask			Any change in input focus wanted
// PropertyChangeMask		Any change in property wanted
// ColormapChangeMask		Any change in colormap wanted
// OwnerGrabButtonMask		Automatic grabs should activate with owner_events set to True

// KeyPress,          = 2
// KeyRelease		  = 3
// ButtonPress,       = 4
// ButtonRelease,     = 5
// MotionNotify       = 6
// EnterNotify,       = 7
// LeaveNotify        = 8
// FocusIn,           = 9
// FocusOut           = 10
// KeymapNotify       = 11
// Expose,            = 12
// GraphicsExpose,    = 13
// NoExpose           = 14
// VisibilityNotify   = 15
// CreateNotify,      = 16
// DestroyNotify,     = 17
// UnmapNotify,       = 18
// MapNotify,         = 19
// MapRequest,        = 20
// ReparentNotify,    = 21
// ConfigureNotify,   = 22
// ConfigureRequest,  = 23
// GravityNotify,     = 24
// ResizeRequest      = 25
// CirculateNotify,   = 26
// CirculateRequest,  = 27
// PropertyNotify,    = 28
// SelectionClear,    = 29
// SelectionRequest,  = 30
// SelectionNotify,   = 31
// ColormapNotify     = 32
// ClientMessage,     = 33
// MappingNotify,     = 34
