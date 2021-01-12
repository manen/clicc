package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

var i = 0

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

func newEnableButton() *gtk.Button {
	b, err := gtk.ButtonNew()
	if err != nil {
		fmt.Println(err)
	}

	b.SetLabel("Enable")
	addClass(b, "butt")

	return b
}

func newDescLabel() *gtk.Label {
	l, err := gtk.LabelNew("Clicc is disabled")
	if err != nil {
		fmt.Println(err)
	}

	return l
}
