package main

import (
	"fmt"
	"log"
	"net/http"

	src "./src"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load the DB
	src.GetPosts()
	src.GetComments()
	src.GetAccounts()
	src.GetSessions()

	// Dispatch handles
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/home", src.Home)
	http.Handle("/home/", http.NotFoundHandler())
	http.HandleFunc("/homeLogged", src.HomeLogged)
	http.Handle("/homeLogged/", http.NotFoundHandler())
	http.HandleFunc("/dashboard", src.Dashboard)
	http.Handle("/dashboard/", http.NotFoundHandler())
	http.HandleFunc("/login", src.Login)
	http.Handle("login/", http.NotFoundHandler())
	http.HandleFunc("/register", src.Register)
	http.Handle("/register/", http.NotFoundHandler())
	http.HandleFunc("/liked", src.Liked)
	http.Handle("/liked/", http.NotFoundHandler())
	http.HandleFunc("/posted", src.Posted)
	http.Handle("/posted/", http.NotFoundHandler())

	// Load static folder # Front end
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Run and listen to the server
	fmt.Println("listening on: http://localhost:2030/home")
	errServe := http.ListenAndServe(":2030", nil)

	if errServe != nil {
		log.Fatalln("In main :", errServe)
	}
}
