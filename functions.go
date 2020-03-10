//Package app contains functions and methods to create and deal with the app database
package app

import (
	"Go_Projects/Encrypt/dbs"
	_ "Go_Projects/Encrypt/dbs" // prevent exposure of secret keys 
	// "fmt"
)

// Something is just a method to pretest installation of the app package
func Something(i string) string{
	dbs.Run()
	return "app connected" + i
}

// Errormessage(name)
