package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Load the pages

	//Load static folder # Front end

	//Run and listen to the server
	fmt.Println("listening on: http://localhost:2030")
	err := http.ListenAndServe(":2030", nil)

	if err != nil {
		log.Fatal(err)
	}

}
