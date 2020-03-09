//Package app contains functions and methods to create and deal with the app database
package app

import _ "Go_Projects/Encrypt/app/dbs" // prevent exposure of secret keys 
import "fmt"

// Something is just a method to pretest installation of the app package
func Something(i string) string{
	db := dbs.dbLink()
	fmt.Println(db)
	return "app connected" + i
}