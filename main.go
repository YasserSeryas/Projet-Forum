package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	h "./src"
	_ "github.com/mattn/go-sqlite3"
)

var Database, _ = sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")

func main() {
	h.ShowBdd()
	//Load the pages
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/home", h.Home)
	http.Handle("/home/", http.NotFoundHandler())
	http.HandleFunc("/homeLogged", h.HomeLogged)
	http.Handle("/homeLogged/", http.NotFoundHandler())
	http.HandleFunc("/dashboard", h.Dashboard)
	http.Handle("/dashboard/", http.NotFoundHandler())
	http.HandleFunc("/login", h.Login)
	http.Handle("login/", http.NotFoundHandler())
	http.HandleFunc("/register", h.Register)
	http.Handle("/register/", http.NotFoundHandler())
	http.HandleFunc("/liked", h.Liked)
	http.Handle("/liked/", http.NotFoundHandler())
	http.HandleFunc("/posted", h.Posted)
	http.Handle("/posted/", http.NotFoundHandler())
	http.HandleFunc("/insert", h.Insert)
	http.Handle("/insert/", http.NotFoundHandler())
	http.HandleFunc("/AddComment", h.AddComment)
	http.Handle("/AddComment/", http.NotFoundHandler())

	//Load static folder # Front end
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//Run and listen to the server
	fmt.Println("listening on: http://localhost:2030/home")
	err := http.ListenAndServe(":2030", nil)

	if err != nil {
		log.Fatal(err)
	}

}
