package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	h "./src"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var Db, errBDD = sql.Open("sqlite3", "BDD/ProjetForum.db")
	if errBDD != nil {
		fmt.Println("here")
		log.Fatal(errBDD)
	}
	resultAccount, errSelect := Db.Query("SELECT name, email, hashPwd FROM Account")
	if errSelect != nil {
		log.Fatal(errSelect)
	}
	for resultAccount.Next() {
		var name string
		var email string
		var hashPwd string
		resultAccount.Scan(&name, &email, &hashPwd)
		fmt.Println(name, email, hashPwd)
	}

	resultPost, errSelect := Db.Query("SELECT IDPost, user, content, likes, dislikes FROM Post")
	if errSelect != nil {
		log.Fatal(errSelect)
	}
	for resultPost.Next() {
		var IDPost int
		var user string
		var content string
		var likes int
		var dislikes int
		resultPost.Scan(&IDPost, &user, &content, &likes, &dislikes)
		fmt.Println(IDPost, user, content, likes, dislikes)
	}
	// statement, errCreate := Db.Prepare("INSERT INTO Account (name, email, hashPwd) VALUES(?, ?, ?)")
	// if errCreate != nil {
	// 	fmt.Println("err Db.Prepare")
	// 	log.Fatal(errCreate)
	// }
	// statement.Exec("Nathan", "nathan.schneider4505@gmail.com", "salutlesgens")

	//Load the pages
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/home", h.Home)
	http.Handle("/home/", http.NotFoundHandler())
	http.HandleFunc("/homeLogged", h.HomeLogged)
	http.Handle("/homeLogged/", http.NotFoundHandler())
	http.HandleFunc("/login", h.Login)
	http.Handle("login/", http.NotFoundHandler())
	http.HandleFunc("/register", h.Register)
	http.Handle("/register/", http.NotFoundHandler())
	http.HandleFunc("/liked", h.Liked)
	http.Handle("/liked/", http.NotFoundHandler())
	http.HandleFunc("/posted", h.Posted)
	http.Handle("/posted/", http.NotFoundHandler())

	//Load static folder # Front end
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//Run and listen to the server
	fmt.Println("listening on: http://localhost:2030")
	fmt.Println("Home page : http://localhost:2030/home")
	errPort := http.ListenAndServe(":2030", nil)

	if errPort != nil {
		log.Fatal(errPort)
	}

}
