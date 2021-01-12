package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

var i = 0

// newMainBox creates the main box containing
// the actual content
func newMainBox() *gtk.Box {
	b, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	if err != nil {
		fmt.Println(err)
	}
	addClass(b, "main-box")
	b.SetHAlign(gtk.ALIGN_CENTER)
	b.SetVAlign(gtk.ALIGN_CENTER)

	bu := newEnableButton()
	l := newDescLabel()

	b.Add(bu)
	b.Add(l)

	return b
}

// newEnableButton creates the "main" button used to
// turn clicc on/off
func newEnableButton() *gtk.Button {
	b, err := gtk.ButtonNew()
	if err != nil {
		fmt.Println(err)
	}

	b.SetLabel("Enable")
	addClass(b, "butt")

	return b
}

// newDescLabel creates the label below the "main"
// button, that says whether Clicc is enabled/disabled,
// countdown, etc
func newDescLabel() *gtk.Label {
	l, err := gtk.LabelNew("Clicc is disabled")
	if err != nil {
		fmt.Println(err)
	}

	return l
}
