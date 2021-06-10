package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT User, Content, Like, Dislike, Comment FROM Post")
	var User string
	var Dislike int
	var Like int
	var Content string
	var Comment string

	for rows.Next() {
		rows.Scan(&User, &Content, &Like, &Dislike, &Comment)
		fmt.Println("User : " + User)
		fmt.Println("Content :" + Content)
		fmt.Println("Like : " + strconv.Itoa(Like))
		fmt.Println("Dislike : " + strconv.Itoa(Dislike))
		fmt.Println("Comment : " + Comment)
	}
}
