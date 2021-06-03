package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	//db, err := sql.Open("sqlite3", "./")
	//Load the pages
	http.Handle("/", http.NotFoundHandler())
	http.HandleFunc("/home", home)
	http.Handle("/home/", http.NotFoundHandler())
	http.HandleFunc("/homeLogged", homeLogged)
	http.Handle("/homeLogged/", http.NotFoundHandler())
	http.HandleFunc("/login", login)
	http.Handle("login/", http.NotFoundHandler())
	http.HandleFunc("/register", register)
	http.Handle("/register/", http.NotFoundHandler())
	http.HandleFunc("/liked", liked)
	http.Handle("/liked/", http.NotFoundHandler())
	http.HandleFunc("/posted", posted)
	http.Handle("/posted/", http.NotFoundHandler())

	//Load static folder # Front end
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//Run and listen to the server
	fmt.Println("listening on: http://localhost:2030")
	err := http.ListenAndServe(":2030", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/home.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHome.Execute(w, nil)
}

func homeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHomeLogged.Execute(w, nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tLogin.Execute(w, nil)
}

func register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tRegister.Execute(w, nil)
}

func liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/liked.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tLiked.Execute(w, nil)
}

func posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tPosted.Execute(w, nil)
}
