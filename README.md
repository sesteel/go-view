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
	Drawer
	event.FocusNotifier
	event.FocusHandler
	event.MouseNotifier
	event.MouseHandler
	SetParent(View)
	Parent() View
	Surface() *Surface
	Name() string
	SetStyle(Style)
	Style() Style
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
Screenshot from the Button Example:<br>
<img src=https://raw.githubusercontent.com/sesteel/go-view/master/res/screenshots/button_example.png>

Screenshot from the Text Box Example:
<img src=http://i.imgur.com/bLXLbXj.png>
