package ui

import (
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var app *gtk.Application

// Init creates a GTK app, initializes the ui, and calls gtk.Init
func Init() {
	gtk.Init(nil)

	const appID = "me.manen.Clicc"
	a, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		panic(err)
	}

	app = a
}

// Start starts the UI of the app
// Actually just calls logic app.Connect "activate" like a boss
// (GTK has been initialized in init)
func Start() {
	app.Connect("activate", logic)
	app.Run(os.Args)
}

// logic contains the actual UI code.
// It is called when the GTK app activates
func logic() {
	win := newWindow()
	win.ShowAll()

	gtk.Main()
}
