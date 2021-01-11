package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// newHeaderBar creates a new gtk.HeaderBar
// Self-explanatory, huh?
func newHeaderBar() *gtk.HeaderBar {
	hb, err := gtk.HeaderBarNew()
	if err != nil {
		fmt.Println(err)
	}

	mb := newMenuButton()

	hb.SetShowCloseButton(true)
	hb.SetSubtitle("Autoclicker")

	hb.PackEnd(mb)

	return hb
}
