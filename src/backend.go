package src

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

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
