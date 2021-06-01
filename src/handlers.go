package helpers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/home.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHome.Execute(w, nil)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHomeLogged.Execute(w, nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tLogin.Execute(w, nil)
}

func Register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tRegister.Execute(w, nil)
}

func Liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/liked.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tLiked.Execute(w, nil)
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tPosted.Execute(w, nil)
}
