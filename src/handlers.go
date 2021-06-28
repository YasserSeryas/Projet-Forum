package src

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html", "templates/navbar.html")
	if err != nil {
		w.WriteHeader(400)
	}

	errtemplate := tHome.Execute(w, AllData)
	if errtemplate != nil {
		log.Fatalln("In home :", errtemplate)
	}
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html", "templates/navbarLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}
	if !CheckSession(w, req) {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}

	if req.Method == "POST" {
		switch req.FormValue("formName") {
		case "addPost":
			if CreatePost(req) != nil {
				http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
			}
		case "addComment":
			CreateComment(req) // A FINIR <--------------------------
		}
	}

	errtemplate := tHomeLogged.Execute(w, AllData)
	if errtemplate != nil {
		log.Fatalln("In home :", errtemplate)
	}
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html", "templates/navbarLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tDashboard.Execute(w, nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	cookie, errCookie := req.Cookie("isLogged")

	if errCookie != http.ErrNoCookie {
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
		userID := ""
		for _, account := range Accounts {
			if (req.FormValue("usernameOrEmail") == account.Name || req.FormValue("usernameOrEmail") == account.Email) &&
				CheckPassword(req.FormValue("pwd"), account.HashPwd) == nil {
				isConnected = true
				userID = account.Email
				break
			}
		}
		UpdateSessionsBDD(req, userID)

		if isConnected {
			fmt.Println("Connected !")

			if !(HasActiveSession(userID)) {
				u := uuid.Must(uuid.NewV4()).String()
				var session Session
				session.SessionUUID = u
				session.UserID = userID
				AddSession(session)

				cookie = &http.Cookie{
					Name:     "isLogged",
					Value:    u,
					HttpOnly: true,
					Path:     "/",
					MaxAge:   AGE_SESSION,
				}
				http.SetCookie(w, cookie)
			}

			http.Redirect(w, req, "http://localhost:2030/homeLogged", http.StatusSeeOther)

		} else {
			fmt.Println("Not connected")
			tLogin.Execute(w, nil)
		}
	} else {
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
		if CreateAccount(req) {
			http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
		} else {
			tRegister.Execute(w, nil)
		}
	} else {
		tRegister.Execute(w, nil)
	}
}

func Liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/liked.html", "templates/navbarLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if !CheckSession(w, req) {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}

	tLiked.Execute(w, nil)
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html", "templates/navbarLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if !CheckSession(w, req) {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}

	tPosted.Execute(w, nil)
}
