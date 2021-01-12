package ui

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

var enableButton *gtk.Button
var descLabel *gtk.Label

// newMainBox creates the main box containing
// the actual content
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
	enableButton = bu
	descLabel = l

	b.Add(bu)
	b.Add(l)

	return b
}

// newEnableButton creates the "main" button used to
// turn clicc on/off
func newEnableButton() *gtk.Button {
	b, err := gtk.ButtonNew()
	if err != nil {
		fmt.Println(err)
	}

	b.SetLabel("Enable")
	addClass(b, "butt")

	return b
}

// newDescLabel creates the label below the "main"
// button, that says whether Clicc is enabled/disabled,
// countdown, etc
func newDescLabel() *gtk.Label {
	l, err := gtk.LabelNew("Clicc is disabled")
	if err != nil {
		fmt.Println(err)
	}

	return l
}

// Adjust UI to when Clicc is disabled
func disableUI() {
	enableButton.SetLabel("Enable")
	ChangeDescLabel("Clicc is disabled")
}

// Adjust UI to when Clicc is enabled
func enableUI() {
	enableButton.SetLabel("Disable")
	ChangeDescLabel("Clicc is enabled")
}

// ChangeEnabledStatus changes the status whether
// Clicc is enabled
func ChangeEnabledStatus(enabled bool) {
	if enableButton == nil {
		fmt.Println("Enable button has not been initialized, ChangeEnabledStatus discarded")
		return
	}

	if enabled {
		enableUI()
	} else {
		disableUI()
	}
}

// ChangeDescLabel changes the text of descLabel
func ChangeDescLabel(newDescLabel string) {
	if descLabel == nil {
		fmt.Println("Desc label has not been initialized, ChangeDescLabel discarded")
		return
	}

	descLabel.SetLabel(newDescLabel)
}
