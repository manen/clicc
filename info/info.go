// Package info contains a bunch of constants (and vars
// that should be considered constants thanks to dumb arrays)
// used in the About page
package info

import (
	"github.com/gotk3/gotk3/gtk"
)

const (
	// Name is the name of the app
	Name string = "Clicc"

	// Desc is the description of the app
	// Preferrably short, since it has to look
	// pretty in a HeaderBar#subtitle
	Desc string = "Autoclicker"

	// Version is the current version of the app
	// Preferrably SemVer
	Version string = "0.0.1"

	// Website is the current homepage of the project
	Website string = "https://github.com/manen/clicc"

	// License is the license the app is licensed under
	License string = "GPL 3.0"

	// LicenseType is the type of license we have
	LicenseType gtk.License = gtk.LICENSE_GPL_3_0
)

// I hate arrays in Go
var (
	// Authors is the list of contributors to the project
	// The first in the list is considered the copyright holder
	//
	// Feel free to add yourself if you think you contributed a lot
	// and you deserve to be here
	Authors []string = []string{
		"manen",
	}

	// Documenters is the list of people who
	// documented all the code
	Documenters []string = []string{
		"manen",
	}

	// Artists is the list of people who
	// drew all the artwork, images, etc...
	//
	// Since Clicc has no art in it, it is empty
	Artists []string = []string{}
)
