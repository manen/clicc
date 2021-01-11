package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// newMenuButton creates the button in the HeaderBar
// that when clicked, opens the menu
func newMenuButton() *gtk.MenuButton {
	mb, err := gtk.MenuButtonNew()
	if err != nil {
		fmt.Println(err)
	}

	img, err := gtk.ImageNewFromIconName("open-menu-symbolic", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		fmt.Println(err)
	}

	mb.SetImage(img)
	mb.SetMenuModel(&newMenu().MenuModel)

	return mb
}

// newMenu creates a glib.Menu
func newMenu() *glib.Menu {
	m := glib.MenuNew()

	m.Append(newAboutButton())

	return m
}
