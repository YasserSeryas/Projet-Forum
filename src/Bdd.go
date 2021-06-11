package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Post() {
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
		fmt.Println("User :" + User)
		fmt.Println("Content :" + Content)
		fmt.Println("Like :" + strconv.Itoa(Like))
		fmt.Println("Dislike :" + strconv.Itoa(Dislike))
		fmt.Println("Comment : " + Comment)
	}
}

func Like() {
	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT User, Date, IdPost, IsLike, IsDislike FROM Like")
	var User string
	var Date time.Time
	var IdPost int
	var IsLike bool
	var IsDislike bool

	for rows.Next() {
		rows.Scan(&User, &Content, &Like, &Dislike, &Comment)
		fmt.Println("User :" + User)
		fmt.Println("Date :" + Date)
		fmt.Println("IdPost :" + strconv.Itoa(IdPost))
		fmt.Println("IsLike :" + strconv.Itoa(IsLike))
		fmt.Println("IsDislike : " + IsDislike)
	}
}
