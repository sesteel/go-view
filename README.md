go-view
=======

Go-View is a project to build a rudimentary GUI library for Go upon a Cairo backend.  Current work centers around building foundational pieces upon which more complexity can be derived.  The core concepts revolve around the View interface (which is sadly a God interface at the moment), a Component interface, a Composite interface, and a Layout interface.  Types which implement these interfaces are used to construct a view tree in a similar fashion to most other GUI toolkits and libraries.	

```go
type Drawer interface {
	// Traverses the view heirarchy drawing dirty
	// views.
	Draw(*Surface)

	// Marks the dirty path up the view heirarchy.
	Redraw()
}

type View interface {
	SetParent(View)
	Parent() View
	Surface() *Surface
	SetName(string)
	Name() string
	Position() (float64, float64)
	SetSize(float64, float64)
	Size() (float64, float64)
	SetStyle(Style)
	Style() Style
	Drawer
	event.FocusNotifier
	event.FocusHandler
	event.MouseNotifier
	event.MouseHandler
}

type Composite interface {
	View
	SetLayout(Layout)
	Layout() Layout
}

type Component interface {
	View
	Focus() bool
	SetFocus(bool)
}

type Layout interface {
	Drawer
	event.MouseHandler
}
```

The following screenshot is from running the textbox application in the examples subdirectory.  
<img src=http://i.imgur.com/bLXLbXj.png>
