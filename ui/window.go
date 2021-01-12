package ui

import (
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

	b := newMainBox()

	hb := newHeaderBar()

	win.SetTitlebar(hb)
	win.Add(b)

	win.SetTitle(info.Name)
	win.SetDefaultSize(300, 125)

	win.Connect("destroy", func() {
		win.Destroy()
		gtk.MainQuit()
		closer.Close()
	})

	return win
}
