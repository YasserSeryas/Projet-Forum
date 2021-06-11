package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT User, Content, Like, Dislike, Comment , CreationDate FROM Post")
	var User string
	var Dislike int
	var Like int
	var Content string
	var Comment string
	var CreationDate time.Time

	for rows.Next() {
		rows.Scan(&User, &Content, &Like, &Dislike, &Comment, &CreationDate)
		fmt.Println("User : " + User)
		fmt.Println("Content :" + Content)
		fmt.Println("Like : " + strconv.Itoa(Like))
		fmt.Println("Dislike : " + strconv.Itoa(Dislike))
		fmt.Println("Comment : " + Comment)
		fmt.Println("Creation Date : ", CreationDate)
	}
}
