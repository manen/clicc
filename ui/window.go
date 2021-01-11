package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/manen/clicc/info"
	"github.com/xlab/closer"
)

// newWindow creates a new ApplicationWindow,
// puts in all the widgets we need, and show the window
func newWindow() *gtk.ApplicationWindow {
	win, err := gtk.ApplicationWindowNew(app)
	if err != nil {
		panic(err)
	}

	l, err := gtk.LabelNew("Hello World!")
	if err != nil {
		fmt.Println(err)
	}

	hb := newHeaderBar()

	win.SetTitlebar(hb)
	win.Add(l)

	win.SetTitle(info.Name)
	win.SetDefaultSize(600, 350)

	win.Connect("destroy", func() {
		win.Destroy()
		gtk.MainQuit()
		closer.Exit(0)
	})

	return win
}
