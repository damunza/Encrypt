//Package app contains functions and methods to create and deal with the app database
package app

import (
	_"Go_Projects/Encrypt/dbs"// prevent exposure of secret keys 
	// _ "Go_Projects/Encrypt/dbs" 
	// "fmt"
)

// Something is just a method to pretest installation of the app package
func Something(i string) string{
	return "app connected" + i
}

// Errormessage(name)
