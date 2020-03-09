//Package app contains functions and methods to create and deal with the app database
package app

//_"Go_Projects/Encrypt/app/dbs" import to prevent exposing secretkey
import "Go_Projects/Encrypt/app/dbs"
import "fmt"

// Something is just a method to pretest installation of the app package
func Something(i string) string{
	fmt.Println(dbs.dbLink())
	return "app connected" + i
}