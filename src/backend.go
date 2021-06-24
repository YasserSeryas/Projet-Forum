package src

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(req *http.Request) {
	var newAccount Account
	if req.FormValue("pwd") == req.FormValue("secondPwd") {
		newAccount.Email = req.FormValue("email")
		newAccount.Name = req.FormValue("username")
		newAccount.HashPwd = HashPassword(req.FormValue("pwd"))

		AddAccount(newAccount)
	}
}

func HashPassword(password string) string {
	hashedPwd, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		fmt.Print("Failed to hash password : ")
		log.Fatal(errHash)
	}
	return string(hashedPwd)
}

func CheckPassword(password string, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}

func HasActiveSession(user string) bool {
	for _, account := range Accounts {
		if account.Email == user {
			return true
		}
	}
	return false
}

func UpdateSessionsBDD(req *http.Request, user string) {
	cookie, errCookie := req.Cookie("isLogged")
	if errCookie == http.ErrNoCookie {
		DeleteSession(user)
	} else if errCookie != nil {
		log.Fatalln("In HomeLogged : errCookie :", errCookie)
	} else {
		if cookie.Value == "0" {
			DeleteSession(user)
		}
	}
}

func CheckSession(w http.ResponseWriter, req *http.Request) bool {

	cookie, errCookie := req.Cookie("isLogged")
	if errCookie == http.ErrNoCookie {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	} else if errCookie != nil {
		log.Fatalln("In HomeLogged : errCookie :", errCookie)
	}

	for _, session := range Sessions {
		if session.SessionUUID == cookie.Value {
			cookie.MaxAge = AGE_SESSION
			return true
		}
	}

	return false
}
