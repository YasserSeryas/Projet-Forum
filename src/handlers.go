package helpers

import (
	"fmt"
	"html/template"
	"net/http"
)

type accountStruct struct {
	username string
	email    string
	pwd      string
}

var account accountStruct

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHome.Execute(w, nil)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	isGoodSession := false
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	for _, cookie := range req.Cookies() {
		if cookie.Name == "isLogged" && cookie.Value == "1" {
			isGoodSession = true
			break
		}
	}

	if isGoodSession {
		tHomeLogged.Execute(w, nil)
	} else {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}

}

func Login(w http.ResponseWriter, req *http.Request) {
	cookie, errCookie := req.Cookie("isLogged")

	// No cookie
	if errCookie == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "isLogged",
			Value: "0",
		}
	}

	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if req.Method == "POST" {
		if (req.FormValue("usernameOrEmail") == account.username || req.FormValue("usernameOrEmail") == account.email) && req.FormValue("pwd") == account.pwd {
			fmt.Println("Connected !")
			cookie = &http.Cookie{
				Name:  "isLogged",
				Value: "1",
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, req, "http://localhost:2030/homeLogged", http.StatusSeeOther)
		} else {
			fmt.Println("Not connected")
			tLogin.Execute(w, nil)
		}
	} else {
		cookie = &http.Cookie{
			Name:  "isLogged",
			Value: "0",
		}
		http.SetCookie(w, cookie)
		tLogin.Execute(w, nil)
	}

}

func Register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if req.Method == "POST" {
		account.username = req.FormValue("username")
		account.email = req.FormValue("email")
		if req.FormValue("pwd") == req.FormValue("secondPwd") {
			account.pwd = req.FormValue("pwd")
		}
		fmt.Println("Username :", account.username, "\nEmail :", account.email, "\nPassword :", account.pwd)
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
