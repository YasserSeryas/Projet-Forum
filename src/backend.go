package src

import (
	"log"
	"net"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Account creation in the database function
func CreateAccount(req *http.Request) bool {
	var newAccount Account
	// Test if password and confirmed password are equal
	isGoodSecondPwd := req.FormValue("pwd") == req.FormValue("secondPwd")
	// Test if password has good length
	isGoodPwd := len(req.FormValue("pwd")) > 7
	// Test if email has good shape and length
	isGoodEmail := EmailRegex.MatchString(req.FormValue("email")) && len(req.FormValue("email")) > 3 && len(req.FormValue("email")) < 200
	// Test if email's mx exists
	mx, errMX := net.LookupMX(strings.Split(req.FormValue("email"), "@")[1])
	isGoodMX := errMX == nil || len(mx) > 0
	// Test if email doesn't exist in DB
	emailDoesntExist := true
	for _, account := range Accounts {
		if account.Email == req.FormValue("email") {
			emailDoesntExist = false
			break
		}
	}

	// Then create an account or not
	if isGoodSecondPwd && isGoodEmail && isGoodPwd && isGoodMX && emailDoesntExist {
		newAccount.Email = req.FormValue("email")
		newAccount.Name = req.FormValue("username")
		newAccount.HashPwd = HashPassword(req.FormValue("pwd"))

		AddAccount(newAccount)
		GetAccounts()
	}

	return isGoodSecondPwd && isGoodEmail && isGoodPwd && isGoodMX && emailDoesntExist
}

// A FINIR <--------------------------------------
func CreatePost(req *http.Request) {
	var newPost Post
	newPost.User = GetUserFromCookie(req)
	newPost.Category = req.FormValue("category")
	newPost.Content = req.FormValue("usertxt")
	AddPost(newPost)
}

// Password hashing function
func HashPassword(password string) string {
	hashedPwd, errHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHash != nil {
		log.Fatalln("Failed to hash password : ", errHash)
	}
	return string(hashedPwd)
}

// Password's check function
func CheckPassword(password string, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}

// Active session test function
func HasActiveSession(user string) bool {
	for _, session := range Sessions {
		if session.UserID == user {
			return true
		}
	}
	return false
}

// Session's update function
func UpdateSessionsBDD(req *http.Request, user string) {
	cookie, errCookie := req.Cookie("isLogged")
	if errCookie == http.ErrNoCookie {
		DeleteSession(user)
	} else if errCookie != nil {
		log.Fatalln("In UpdateSessionsBDD : errCookie :", errCookie)
	} else {
		if cookie.Value == "0" {
			DeleteSession(user)
		}
	}
}

// Session's check function
func CheckSession(w http.ResponseWriter, req *http.Request) bool {

	cookie, errCookie := req.Cookie("isLogged")
	if errCookie == http.ErrNoCookie {
		http.Redirect(w, req, "http://localhost:2030/login", http.StatusSeeOther)
	} else if errCookie != nil {
		log.Fatalln("In CheckSession : errCookie :", errCookie)
	}
	for _, session := range Sessions {
		if session.SessionUUID == cookie.Value {
			cookie.MaxAge = AGE_SESSION
			return true
		}
	}

	return false
}

// Get user ID from a "isLogged" cookie
func GetUserFromCookie(req *http.Request) string {
	cookie, errCookie := req.Cookie("isLogged")
	userID := ""

	if errCookie == http.ErrNoCookie {
		return userID
	}

	for _, session := range Sessions {
		if session.SessionUUID == cookie.Value {
			userID = session.UserID
		}
	}

	return userID
}
