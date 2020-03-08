package dbs

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func errorMessage(element string)string{
	// error message for the app packages
	return element + "- Not created or called correctly"
}

func dbLink(){
	// creating a db object to access the database
	db,err := sql.Open("mysql", "root:L0ngestN1ght9828@tcp(127.0.0.1:3306)/")
	if err != nil{
		err = errorMessage(db)
		panic(err)
	}else{
		fmt.Println("connected to db")
	}
}
defer db.Close()