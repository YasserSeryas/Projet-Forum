package src

import (
	"fmt"
	"html/template"
	"log"
	"net/http"



	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

var (
	Logger  *log.Logger
)

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //Si err n'est pas nul, une erreur interne du serveur est retourné
		Logger.Printf("Erreur interne du serveur sur la page: Home | ", err.Error()) //affiche dans le fichier test.txt le message entre guillemets
			return 
	}

	errtemplate := tHome.Execute(w, AllData)
	if errtemplate != nil {
		log.Fatalln("In home :", errtemplate)
	}
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html", "templates/navbarLogged.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: HomeLogged(1) | ", err.Error()) 
			return 
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
		case "actions":
			if req.FormValue("isLike") == "like" {
				CreateLike(req)
			} else if req.FormValue("isLike") == "dislike" {
				CreateDislike(req)
			}
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Dashboard | ", err.Error()) 
			return 
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Login | ", err.Error()) 
			return 
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Register | ", err.Error()) 
			return 
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Liked | ", err.Error()) 
			return 
	}

	if !CheckSession(w, req) {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}

	tLiked.Execute(w, nil)
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html", "templates/navbarLogged.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Posted | ", err.Error()) 
			return 
	}

	if !CheckSession(w, req) {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}

	tPosted.Execute(w, nil)
}

func Profile(w http.ResponseWriter, req *http.Request) {
	tProfile, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Logger.Printf("Erreur interne du serveur sur la page: Profile | \n", err.Error()) 
			return 
	}

	tProfile.Execute(w, nil)
} 