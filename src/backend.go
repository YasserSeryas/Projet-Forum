package src

import (
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Account creation in the database function
func CreateAccount(req *http.Request) bool {
	var newAccount Account
	// Test if password and confirmed password are equal
	isGoodSecondPwd := req.FormValue("pwd") == req.FormValue("secondPwd")
	if !isGoodSecondPwd {
		ErrorCreationPost = errors.New("veuillez rentrer le même mot de passe dans les deux champs")
	}
	// Test if password has good length
	isGoodPwd := len(req.FormValue("pwd")) > 7
	if !isGoodPwd {
		ErrorCreationPost = errors.New("veuillez saisir un mot de passe de plus de 7 charactères")
	}
	// Test if email has good shape and length
	isGoodEmail := EmailRegex.MatchString(req.FormValue("email")) && len(req.FormValue("email")) > 3 && len(req.FormValue("email")) < 200
	if !isGoodEmail {
		ErrorCreationPost = errors.New("veuillez saisir un email valide")
	}
	// Test if email's mx exists
	mx, errMX := net.LookupMX(strings.Split(req.FormValue("email"), "@")[1])
	isGoodMX := errMX == nil || len(mx) > 0
	if !isGoodMX {
		ErrorCreationPost = errors.New("veuillez saisir un domaine d'adresse valide")
	}
	// Test if email doesn't exist in DB
	emailDoesntExist := true
	for _, account := range Accounts {
		if account.Email == req.FormValue("email") {
			emailDoesntExist = false
			break
		}
	}
	if !emailDoesntExist {
		ErrorCreationPost = errors.New("un compte existe déjà pour cette adresse mail")
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

// Function of post creation in database
func CreatePost(req *http.Request) error {
	var newPost Post
	var err error
	if GetUserFromCookie(req) == "" {
		err = errors.New("in /src/backend.go/CreatePost : no user find")
	} else {
		newPost.User = GetUserFromCookie(req)
	}

	newPost.Category = req.FormValue("category")
	newPost.Content = req.FormValue("usertxt")
	newPost.Title = req.FormValue("postTitle")
	AddPost(newPost)
	return err
}

// Function of comment creation in database
func CreateComment(req *http.Request) error {
	var newComment Comment
	var err error

	IdPost, errAtoi := strconv.Atoi(req.FormValue("Idpost"))
	if errAtoi != nil {
		log.Fatalln("In backend.go/CreateComment :", errAtoi)
	} else {
		newComment.IdPost = IdPost
	}

	if GetUserFromCookie(req) == "" {
		err = errors.New("in /src/backend.go/CreatePost : no user find")
		return err
	} else {
		newComment.IdUser = GetUserFromCookie(req)
	}

	newComment.UserName = GetUsername(newComment.IdUser)
	newComment.CommentContent = req.FormValue("usertxt")

	AddComment(newComment)
	return err
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
		return false
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

func CreateLike(req *http.Request) {
	GetLike()
	var IdUser string = GetUserFromCookie(req)
	IdPost, _ := strconv.Atoi(req.FormValue("IdPost"))
	for _, like := range Likes {
		if like.IdPost == IdPost && like.IdUser == IdUser {
			if like.IsLike {
				DeleteLike(like.IdLike)
				UpdateNbrLike(IdPost, false)
				return
			} else {
				UpdateLike(like.IdLike, true)
				UpdateNbrLike(IdPost, true)
				UpdateNbrDislike(IdPost, false)
				return
			}
		}
	}
	var newLike Like
	newLike.IdPost = IdPost
	newLike.IdUser = IdUser
	newLike.IsLike = true
	AddLike(newLike)
	UpdateNbrLike(IdPost, true)
}

func CreateDislike(req *http.Request) {
	GetLike()
	var IdUser string = GetUserFromCookie(req)
	IdPost, _ := strconv.Atoi(req.FormValue("IdPost"))
	for _, like := range Likes {
		if like.IdPost == IdPost && like.IdUser == IdUser {
			if like.IsLike {
				UpdateLike(like.IdLike, false)
				UpdateNbrLike(IdPost, false)
				UpdateNbrDislike(IdPost, true)
				return
			} else {
				UpdateNbrDislike(IdPost, false)
				DeleteLike(like.IdLike)
				return
			}
		}
	}
	var newDislike Like
	newDislike.IdPost = IdPost
	newDislike.IdUser = IdUser
	newDislike.IsLike = false
	AddLike(newDislike)
	UpdateNbrDislike(IdPost, true)
}

func GetLikedPosts(req *http.Request) error {
	var likedPosts []TemplateData
	var user string
	var errCookie error
	if GetUserFromCookie(req) == "" {
		errCookie = errors.New("in /src/backend.go/GetLikedPosts : no user find")
	} else {
		user = GetUserFromCookie(req)
	}

	for _, like := range Likes {
		if like.IdUser == user {
			for _, data := range AllData {
				if data.PostData.IdPost == like.IdPost {
					likedPosts = append(likedPosts, data)
				}
			}
		}
	}
	LikedPosts = likedPosts
	return errCookie
}

func GetPostedPosts(req *http.Request) error {
	var postedPosts []TemplateData
	var user string
	var errCookie error
	if GetUserFromCookie(req) == "" {
		errCookie = errors.New("in /src/backend.go/GetPostedPosts : no user find")
	} else {
		user = GetUserFromCookie(req)
	}

	for _, data := range AllData {
		if data.PostData.User == user {
			postedPosts = append(postedPosts, data)
		}
	}
	PostedPosts = postedPosts
	return errCookie
}
