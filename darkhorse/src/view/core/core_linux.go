// +build linux,!goci
package core

// #cgo pkg-config: cairo x11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xresource.h>
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
	"runtime"
	"unsafe"
)

func init() {
	C.XInitThreads()
}

type XWindow struct {
	cairoSurface *Surface
	display      *C.Display
	screen       C.int
	xwindow      C.Window
	name         string
	width        uint
	height       uint
}

func (self *XWindow) GetSurface() *Surface {
	return self.cairoSurface
}

func (self *XWindow) SetName(name string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	self.name = name
	C.XStoreName(self.display, self.xwindow, n)
}

func (self *XWindow) GetName() string {
	return self.name
}

func (self *XWindow) SetSize(width, height uint) {
	self.width = width
	self.height = height
	C.cairo_xlib_surface_set_size(self.cairoSurface.surface, C.int(self.width), C.int(self.height))
	C.cairo_xlib_surface_set_drawable(self.cairoSurface.surface, C.Drawable(self.xwindow), C.int(self.width), C.int(self.height))
	C.XResizeWindow(self.display, self.xwindow, C.uint(self.width), C.uint(self.height))
	self.paint()
}

func (self *XWindow) GetSize() (width, height uint) {
	return self.width, self.height
}

func (self *XWindow) paint() {
	fmt.Println("PAINT")
	self.cairoSurface.SetSourceRGB(0, 0, 0)
	self.cairoSurface.Paint()
	self.cairoSurface.SetLineWidth(1.5)
	self.cairoSurface.SetSourceRGB(0, 1, 1)
	self.cairoSurface.Rectangle(10, 10, float64(self.width)/1.2, float64(self.height)/1.2)
	self.cairoSurface.Stroke()
}

func NewWindow(name string, x, y, w, h uint) Window {
	var width C.uint = C.uint(w)
	var height C.uint = C.uint(h)
	var ev C.XEvent

	/* First connect to the display server, as specified in the DISPLAY
	environment variable. */
	dpy := C.XOpenDisplay(nil)

	if dpy == nil {
		fmt.Println("unable to connect to display")
		return nil
	}

	/* these are macros that pull useful data out of the display object */
	/* we use these bits of info enough to want them in their own variables */
	screen_num := C.XDefaultScreen(dpy)
	background := C.XBlackPixel(dpy, screen_num)
	border := C.XWhitePixel(dpy, screen_num)

	win := C.XCreateSimpleWindow(dpy, C.XDefaultRootWindow(dpy), /* display, parent */
		0, 0, /* x, y: the window manager will place the window elsewhere */
		width, height, /* width, height */
		2, border, /* border width & colour, unless you have a window manager */
		background) /* background colour */

	/* tell the display server what kind of events we would like to see */
	C.XSelectInput(dpy, win, C.ExposureMask        |
							 C.ButtonPressMask     |
							 C.ButtonReleaseMask   |
							 C.KeyPressMask        |
							 C.KeyReleaseMask      |
							 C.PointerMotionMask   |
							 C.EnterWindowMask     |
							 C.LeaveWindowMask     |
							 C.StructureNotifyMask )

	/* okay, put the window on the screen, please */
	C.XMapWindow(dpy, win)

	s := C.cairo_xlib_surface_create(dpy, C.Drawable(win), C.XDefaultVisual(dpy, 0), C.int(width), C.int(height))
	ctx := C.cairo_create(s)
	surface := &Surface{surface: s, context: ctx}

	window := &XWindow{surface, dpy, screen_num, win, name, uint(width), uint(height)}
	window.SetName(name)
	runtime.SetFinalizer(window, func(w *XWindow) {
		C.XCloseDisplay(window.display)
	})

	window.paint()

	eventLoop := func() {
		/* as each event that we asked about occurs, we respond.  In this
		 * case we note if the window's shape changed, and exit if a button
		 * is pressed inside the window */
		for {
			C.XNextEvent(dpy, &ev)
			eventType := ev[0]
			switch eventType {
			case C.ConfigureNotify:
				event := (*C.XConfigureEvent)(unsafe.Pointer(&ev[0]))
				if width != C.uint(event.width) || height != C.uint(event.height) {
					width = C.uint(event.width)
					height = C.uint(event.height)
					window.width = uint(width)
					window.height = uint(height)
					C.cairo_xlib_surface_set_size(s, event.width, event.height)
					fmt.Println("Size changed to: %d by %d\n", width, height)
					window.paint()
				}

			case C.Expose:
				fmt.Println("Expose")
				window.paint()

			case C.MotionNotify:
				fmt.Println("MOUSE MOVE \n")

			case C.ButtonPress:
				C.XCloseDisplay(dpy)
				return
			}
		}
	}
	go eventLoop()

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
