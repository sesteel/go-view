// +build linux,!goci
package view

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
	"reflect"
	"runtime"
	"time"
	"ui/view/color"
	"unsafe"
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
}

func (self *Window) Parent() View {
	return nil
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
	self.DrawSelf()
}

func (self *Window) Draw(surface *Surface) {
	if self.layout == nil {
		surface.SetSourceRGBA(color.HexRGBA(0xFFFFFFFF))
		surface.Paint()
	} else {
		self.layout.Draw(surface)
	}
}

func (self *Window) DrawSelf() {

	var before time.Time
	if DEBUG_DRAW_ALL {
		before = time.Now()
	}

	s := NewSurface(FORMAT_ARGB32, int(self.width), int(self.height))
	self.Draw(s)
	s.Destroy()
	self.surface.SetSourceSurface(s, 0, 0)
	self.surface.Paint()
	
	if DEBUG_DRAW_ALL {
		fmt.Println("Draw Window:", time.Since(before))
	}
}

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
	if result == 1 {
		fmt.Println("Cannot create 32bit depth display.", result)
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

	win := C.XCreateWindow(dpy, C.XDefaultRootWindow(dpy), 0, 0, 300, 200, 0, vinfo.depth, C.InputOutput, vinfo.visual, C.CWColormap|C.CWBorderPixel, &attr)

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

	window.SetName(name)
	runtime.SetFinalizer(window, func(w *Window) {
		C.XCloseDisplay(window.display)
	})

	window.DrawSelf()

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
					window.width = float64(width)
					window.height = float64(height)
					C.cairo_xlib_surface_set_size(s, event.width, event.height)
				}

			case C.Expose:
				event := (*C.XExposeEvent)(unsafe.Pointer(&ev[0]))
				if event.count > 0 {
					continue
				}
				window.DrawSelf()

			case C.MotionNotify:
				//view.DispatchMouseExit()

			case C.ButtonPress:

			default:
				C.XFlush(dpy)
			}
		}
	}
	go eventLoop()
	//	ticker := time.NewTicker(200 * time.Millisecond)
	//    go func() {
	//	    for {
	//	       select {
	//	        case <- ticker.C:
	//	        	window.DrawSelf()
	//	        }
	//    	}
	// 	}()

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
