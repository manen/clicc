package ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/manen/clicc/info"
)

func newAboutButton() (name string, actName string) {
	ad := newAboutDialog()

	act := glib.SimpleActionNew("open-about", nil)
	act.Connect("activate", func() {
		ad.Show()
	})
	app.AddAction(act)

	return "About " + info.Name, "app.open-about"
}

func newAboutDialog() *gtk.AboutDialog {
	ad, err := gtk.AboutDialogNew()
	if err != nil {
		fmt.Println(err)
	}
	app.AddWindow(ad)

	ad.SetApplication(app)
	ad.SetAuthors(info.Authors)
	ad.SetDocumenters(info.Documenters)
	ad.SetArtists(info.Artists)
	ad.SetCopyright("Copyright " + strconv.Itoa(time.Now().Year()) + " " + info.Authors[0])
	ad.SetLicense(info.License)
	ad.SetLicenseType(gtk.LICENSE_GPL_3_0)
	ad.SetVersion(info.Version)
	ad.SetProgramName(info.Name)
	ad.SetWebsite(info.Website)

	// TODO set icon or whatever im tired

	ad.Connect("delete-event", func() bool { // <something about orphaned widgets or whatever>, but if we do this it works, so..
		ad.Hide()
		return true
	})

	return ad
}
