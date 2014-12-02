
Current work has moved to:

  https://github.com/vizstra/ui

where opengl has replaced Cairo for drawing calls.  The tradeoffs (complexity) to get efficient text rendering using Cairo are not something I want to make at the moment.  At least not without trying a different approach using OpenGL first.

go-view
=======

Go-View is an experimental project to build a rudimentary GUI library for Go upon a Cairo backend.  It is being built out in the open for anybody to use, but it is a personal project and should be viewed as such.  I am exploring tangential issues surrounding UI and UX design.  Below, I have placed a few screenshots to give a sense of the asthetic.

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
