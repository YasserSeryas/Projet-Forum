package helpers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const AGE_SESSION = 60 * 20

var Db, errBDD = sql.Open("sqlite3", "./BDD/ProjetForum.db")

func Home(w http.ResponseWriter, req *http.Request) {
	tHome, err := template.ParseFiles("templates/index.html")
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

	if checkSession(w, req) {
		tHomeLogged.Execute(w, nil)
	} else {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, req *http.Request) {
	cookie, errCookie := req.Cookie("isLogged")

	if errCookie != http.ErrNoCookie {
		cookie.MaxAge = -1
	}

	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if req.Method == "POST" {
		isConnected := false
		userID := ""
		result, errSelect := Db.Query("SELECT name, email, hashPwd FROM Account") // SELECT das la BDD
		if errSelect != nil {                                                     // Test d'erreur
			fmt.Print("In Login : errSelect : ")
			log.Fatal(errSelect)
		}
		for result.Next() { // On boucle le résultat du SELECT
			var name string
			var email string
			var hashPwd string
			result.Scan(&name, &email, &hashPwd)
			userID = email
			if (req.FormValue("usernameOrEmail") == name || req.FormValue("usernameOrEmail") == email) && checkPassword(req.FormValue("pwd"), hashPwd) == nil {
				isConnected = true
				break
			}
		}
		result.Close()

		if isConnected {
			fmt.Println("Connected !")
			fmt.Println(!(hasActiveSession(userID)))

			updateSessionsBDD(req, userID)

			if !(hasActiveSession(userID)) {
				u := uuid.Must(uuid.NewV4())
				statement, _ := Db.Prepare("INSERT INTO Session (sessionsUUID, userID) VALUES(?, ?)")
				_, errCreate := statement.Exec(u.String(), userID)
				if errCreate != nil {
					log.Fatalln("In Login : errCreate : ", errCreate)
				}
				statement.Close()

				cookie = &http.Cookie{
					Name:     "isLogged",
					Value:    u.String(),
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
	if errBDD != nil {
		fmt.Print("In Register : errBDD :")
		log.Fatal(errBDD)
	}

	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if req.Method == "POST" {
		statement, _ := Db.Prepare("INSERT INTO Account (name, email, hashPwd) VALUES(?, ?, ?)")
		if req.FormValue("pwd") == req.FormValue("secondPwd") {
			_, errCreate := statement.Exec(req.FormValue("username"), req.FormValue("email"), hashPassword(req.FormValue("pwd")))
			if errCreate != nil {
				fmt.Print("In Register : errCreate :")
				log.Fatal(errCreate)
			}
		}
		statement.Close()
		showBDD()
	}

	tRegister.Execute(w, nil)
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if checkSession(w, req) {
		tDashboard.Execute(w, nil)
	} else {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}
}

func Liked(w http.ResponseWriter, req *http.Request) {
	tLiked, err := template.ParseFiles("templates/liked.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if checkSession(w, req) {
		tLiked.Execute(w, nil)
	} else {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}
}

func Posted(w http.ResponseWriter, req *http.Request) {
	tPosted, err := template.ParseFiles("templates/posted.html")
	if err != nil {
		w.WriteHeader(400)
	}

	if checkSession(w, req) {
		tPosted.Execute(w, nil)
	} else {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	}
}

func showBDD() {
	result, errSelect := Db.Query("SELECT name, email, hashPwd FROM Account")
	if errSelect != nil {
		fmt.Print("In showBDD : errSelect : ")
		log.Fatal(errSelect)
	}
	for result.Next() {
		var name string
		var email string
		var hashPwd string
		result.Scan(&name, &email, &hashPwd)
		fmt.Println(name, email, hashPwd)
	}
	result.Close()
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

func hasActiveSession(user string) bool {
	userInTable := ""
	result, errSelect := Db.Query("SELECT userID FROM Session") // SELECT dans la BDD
	if errSelect != nil {                                       // Test d'erreur
		log.Fatalln("In dashboard : errSelect : ", errSelect)
	}
	for result.Next() { // On boucle le résultat du SELECT
		result.Scan(&userInTable)
		if userInTable == user {
			result.Close()
			return true
		}
	}
	result.Close()
	return false
}

func updateSessionsBDD(req *http.Request, user string) {
	cookie, errCookie := req.Cookie("isLogged")
	if errCookie == http.ErrNoCookie {
		statement, errDelete := Db.Prepare("DELETE FROM Session WHERE userID = ?;")
		statement.Exec(user)
		if errDelete != nil {
			log.Fatalln("In main : errDelete : ", errDelete)
		}
	} else if errCookie != nil {
		log.Fatalln("In HomeLogged : errCookie :", errCookie)
	} else {
		if cookie.Value == "0" {
			statement, errDelete := Db.Prepare("DELETE FROM Session WHERE userID = ?;")
			statement.Exec(user)
			if errDelete != nil {
				log.Fatalln("In main : errDelete : ", errDelete)
			}
		}
	}
}

func checkSession(w http.ResponseWriter, req *http.Request) bool {
	cookie, errCookie := req.Cookie("isLogged")
	if errCookie == http.ErrNoCookie {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	} else if errCookie != nil {
		log.Fatalln("In HomeLogged : errCookie :", errCookie)
	}

	var currentUser string
	var userID string
	var sessionsUUID string
	result, errSelect := Db.Query("SELECT sessionsUUID, userID FROM Session")
	if errSelect != nil {
		log.Fatalln("In dashboard : errSelect : ", errSelect)
	}
	for result.Next() {
		result.Scan(&sessionsUUID, &userID)
		if cookie.Value == sessionsUUID {
			currentUser = userID
			break
		}
	}
	result.Close()

	if currentUser != "" {
		cookie.MaxAge = AGE_SESSION
	}

	return currentUser != ""
}
