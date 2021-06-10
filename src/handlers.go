package helpers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var Db, errBDD = sql.Open("sqlite3", "/home/nschneid/Projet-Forum/BDD/ProjetForum.db")

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
		isConnected := false
		result, errSelect := Db.Query("SELECT name, email, hashPwd FROM Account") // SELECT das la BDD
		if errSelect != nil {                                                     // Test d'erreur
			log.Fatal(errSelect)
		}
		for result.Next() { // On boucle le résultat du SELECT
			var name string
			var email string
			var hashPwd string
			result.Scan(&name, &email, &hashPwd)
			if (req.FormValue("usernameOrEmail") == name || req.FormValue("usernameOrEmail") == email) && checkPassword(req.FormValue("pwd"), hashPwd) == nil {
				isConnected = true
				break
			}
		}
		if isConnected {
			fmt.Println("Connected !")
			cookie = &http.Cookie{
				Name:  "isLogged",
				Value: "1",
			}
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
		tLogin.Execute(w, nil)
	}
	http.SetCookie(w, cookie)
}

func Register(w http.ResponseWriter, req *http.Request) {
	if errBDD != nil {
		fmt.Print("Into register : ")
		log.Fatal(errBDD)
	}

	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if req.Method == "POST" {
		statement, errCreate := Db.Prepare("INSERT INTO Account (name, email, hashPwd) VALUES(?, ?, ?)")
		if errCreate != nil {
			fmt.Println("err Db.prepare")
			log.Fatal(errCreate)
		}
		if req.FormValue("pwd") == req.FormValue("secondPwd") {
			statement.Exec(req.FormValue("username"), req.FormValue("email"), hashPassword(req.FormValue("pwd")))
		}
		showBDD()
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

func showBDD() {
	result, errSelect := Db.Query("SELECT name, email, hashPwd FROM Account")
	if errSelect != nil {
		log.Fatal(errSelect)
	}
	for result.Next() {
		var name string
		var email string
		var hashPwd string
		result.Scan(&name, &email, &hashPwd)
		fmt.Println(name, email, hashPwd)
	}
}

func hashPassword(password string) string {
	hashedPwd, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		fmt.Print("Failed to hash password : ")
		log.Fatal(errHash)
	}
	return string(hashedPwd)
}

func checkPassword(password string, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}
