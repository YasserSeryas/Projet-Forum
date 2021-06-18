package helpers

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var Database, _ = sql.Open("sqlite3", "./Bdd/ProjetForumBDD.db")

func Home(w http.ResponseWriter, req *http.Request) {

	rows, err := Database.Query("SELECT User, Content, Like, Dislike, Comment , CreationDate, Category FROM Post")

	var data Post
	result := []Post{}
	for rows.Next() {
		rows.Scan(&data.User, &data.Content, &data.Like, &data.Dislike, &data.Comment, &data.CreationDate, &data.Category)
		//dat := (" User : " + data.User + "\n") + (" Content :" + data.Content + "\n") + (" Like : " + strconv.Itoa(data.Like) + "\n") + (" Dislike : " + strconv.Itoa(data.Dislike) + "\n") + (" Comment : " + data.Comment + "\n") + (" Creation Date : " + data.CreationDate.String() + "\n") + (" Category : " + data.Category + "\n")

		result = append(result, data)
	}
	rows.Close()

	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHome.Execute(w, result)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {
	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHomeLogged.Execute(w, nil)
}

func Login(w http.ResponseWriter, req *http.Request) {
	tLogin, err := template.ParseFiles("templates/login.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tLogin.Execute(w, nil)
}

func Register(w http.ResponseWriter, req *http.Request) {
	tRegister, err := template.ParseFiles("templates/register.html")
	if err != nil {
		w.WriteHeader(400)
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
