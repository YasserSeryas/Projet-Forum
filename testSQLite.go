package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./cheminDeVotreBase.db")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println(db)
}
