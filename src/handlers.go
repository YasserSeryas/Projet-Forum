package helpers

import (
	"html/template"
	"net/http"
	"io/ioutil"
)

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tHome.Execute(w, nil)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tHomeLogged.Execute(w, nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tLogin.Execute(w, nil)
}

func Register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tRegister.Execute(w, nil)
}

func Liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/liked.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tLiked.Execute(w, nil)
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tPosted.Execute(w, nil)
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tDashboard.Execute(w, nil)
}

func Profile(w http.ResponseWriter, req *http.Request) {
	tProfile, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		data := []byte(err.Error())
		ioutil.WriteFile("test.txt", data, 0644)
	}

	tProfile.Execute(w, nil)
}
