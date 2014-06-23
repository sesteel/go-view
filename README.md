go-view
=======

Go-View is an experimental project to build a rudimentary GUI library for Go upon a Cairo backend.  Current work centers around building foundational pieces upon which more complexity can be derived.  The core concepts revolve around the View interface, a Component interface, a Composite interface, and a Layout interface.  Types which implement these interfaces are used to construct a view tree in a similar fashion to most other GUI toolkits and libraries.  There are severe limitations and the project is really meant to explore tangential issues surrounding UI and UX design.  This project is not in a usable state at the moment, but pull requests are welcome.

Screenshot from the Button Example:<br>
<img src=https://raw.githubusercontent.com/sesteel/go-view/master/res/screenshots/button_example.png>

Screenshot from the Checkbox Example:<br>
<img src=https://raw.githubusercontent.com/sesteel/go-view/master/res/screenshots/checkbox_example.png>

Screenshot from the Progress Bar Example:<br>
<img src=https://raw.githubusercontent.com/sesteel/go-view/master/res/screenshots/progress_bar_example.png>

Screenshot from the Text Box Example:<br>
<img src=https://raw.githubusercontent.com/sesteel/go-view/master/res/screenshots/text_box_example.png>

Screenshot from the Editor Example:<br>
<img src=https://raw.githubusercontent.com/sesteel/go-view/master/res/screenshots/editor_example.png>

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

Supported Unicode:
0x0000 - 0x007F : Basic Latin (128)

0x0080 - 0x00FF : Latin-1 Supplement (128)

0x0100 - 0x017F : Latin Extended-A (128)

0x0180 - 0x024F : Latin Extended-B (208)

0x0250 - 0x02AF : IPA Extensions (96)


