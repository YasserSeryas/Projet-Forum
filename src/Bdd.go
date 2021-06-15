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
		rows.Scan(&User, &Dislike, &Like, &Content, &Comment)
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
		rows.Scan(&User, &Date, &IdPost, &IsLike, &IsDislike)
		fmt.Println("User :" + User)
		fmt.Println("Date :" + Date)
		fmt.Println("IdPost :" + strconv.Itoa(IdPost))
		fmt.Println("IsLike :" + strconv.Itoa(IsLike))
		fmt.Println("IsDislike : " + IsDislike)
	}
}

func Post() {
	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT Email, Name, HashPwd, SessionUUID FROM Account")
	var Email string
	var Name string
	var HashPwd string
	var SessionUUID string

	for rows.Next() {
		rows.Scan(&Email, &Name, &HashPwd, &SessionUUID)
		fmt.Println("Email :" + Email)
		fmt.Println("Name :" + Name)
		fmt.Println("HashPwd :" + HashPwd)
		fmt.Println("SessionUUID :" + SessionUUID)
	}

}

func Categorie() {
	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT  Nom, Color FROM Categorie")
	var Nom string
	var Color int

	for rows.Next() {
		rows.Scan(&Nom, &Color)
		fmt.Println("Nom:" + Nom)
		fmt.Println("Color :" + Color)
	}

}

func Topic() {
	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT Id, Title, Content, UserPseudo, Categorie")
	var Id int
	var Title string
	var Content string
	var UserPseudo string
	var Categorie int

	for rows.Next() {
		rows.Scan(&Id, &Title, &Content, &UserPseudo, &Categorie)
		fmt.Println("Title:" + Title)
		fmt.Println("Content:" + Content)
		fmt.Println("UserPseudp" + UserPseudo)
		fmt.Println("Categorie" + Categorie)
	}
}
