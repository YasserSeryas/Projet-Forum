package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//db, err := sql.Open("sqlite3", "./")
	//Load the pages

	//Load static folder # Front end

	//Run and listen to the server
	fmt.Println("listening on: http://localhost:2030")
	err := http.ListenAndServe(":2030", nil)

	if err != nil {
		log.Fatal(err)
	}

}
