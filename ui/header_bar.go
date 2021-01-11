package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
	"github.com/manen/clicc/info"
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
	hb.SetSubtitle(info.Desc)

	hb.PackEnd(mb)

	return hb
}
