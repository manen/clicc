package clicker

import "fmt"

// enabled is whether Clicc is enabled
var enabled bool = false

// Switch switches the clicker
//
// So if it was turned on, it turns off, and
// if it was turned off, it turns on
func Switch() {
	if enabled {
		disable()
	} else {
		enable()
	}
	enabled = !enabled
}

// disable disables Clicc
func disable() {
	fmt.Println("Disabled")
}

// enable enables Clicc
func enable() {
	fmt.Println("Enabled")
}
