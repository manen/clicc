package ui

import (
	"fmt"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

var i = 0

func newBox() *gtk.Box {
	b, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	if err != nil {
		fmt.Println(err)
	}

	addClass(b, "main-box")

	l, err := gtk.LabelNew(strconv.Itoa(i))
	if err != nil {
		fmt.Println(err)
	}

	bu, err := gtk.ButtonNew()
	if err != nil {
		fmt.Println(err)
	}

	bu.SetLabel("++")
	bu.Connect("clicked", func() {
		i++
		l.SetLabel(strconv.Itoa(i))
	})

	b.Add(l)
	b.Add(bu)

	return b
}
