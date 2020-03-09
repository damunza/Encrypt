package dbs

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"errors"
	"fmt"
)

func errorMessage(){
	// error message for the app packages
	err := errors.New("- Not created or called correctly")
	fmt.Println(err)
}

func dbLink(){
	// creating a db object to access the database
	db,err := sql.Open("mysql", "root:L0ngestN1ght9828@tcp(127.0.0.1:3306)/")
	if err != nil{
		panic(err)
	}else{
		fmt.Println("connected to db")
	}
	defer db.Close()
}
