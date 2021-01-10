package ui

import (
	"fmt"
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
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		panic(err)
	}

	win.SetTitle("Clicc")

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("Hello World!")
	if err != nil {
		fmt.Println(err)
	}

	win.Add(l)
	win.SetDefaultSize(600, 300)
	win.ShowAll()

	gtk.Main()
}
