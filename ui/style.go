package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// css is the CssProvider that contains the
// main styles (./ui/style.css)
var css *gtk.CssProvider

func initStyle() {
	cssP, err := gtk.CssProviderNew()
	if err != nil {
		fmt.Println(err)
	}
	err = cssP.LoadFromPath("./ui/style.css")
	if err != nil {
		fmt.Println(err)
	}

	sc, err := gdk.ScreenGetDefault()
	if err != nil {
		fmt.Println(err)
	}

	gtk.AddProviderForScreen(sc, cssP, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	css = cssP
}
