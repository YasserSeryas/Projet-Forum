package helpers

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Home(w http.ResponseWriter, req *http.Request) {

	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHome.Execute(w, Result2)

}

func HomeLogged(w http.ResponseWriter, req *http.Request) {

	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	er := tHomeLogged.Execute(w, Result2)
	fmt.Println(er)
}
func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tDashboard.Execute(w, nil)

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
