package helpers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var Database, _ = sql.Open("sqlite3", "./Bdd/ProjetForumBDD.db")
var Result = []Post{}

func ShowBdd() {
	rows, _ := Database.Query("SELECT [Id-Post], User, Content, Like, Dislike, Comment , CreationDate, Category FROM Post")

	var data Post

	for rows.Next() {
		rows.Scan(&data.IdPost, &data.User, &data.Content, &data.Like, &data.Dislike, &data.Comment, &data.CreationDate, &data.Category)
		//dat := (" User : " + data.User + "\n") + (" Content :" + data.Content + "\n") + (" Like : " + strconv.Itoa(data.Like) + "\n") + (" Dislike : " + strconv.Itoa(data.Dislike) + "\n") + (" Comment : " + data.Comment + "\n") + (" Creation Date : " + data.CreationDate.String() + "\n") + (" Category : " + data.Category + "\n")

		Result = append(Result, data)
	}
	rows.Close()
}

//Insert post
func Insert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stmt, _ := Database.Prepare("INSERT INTO Post( User, Content, Like, Dislike, Comment, Creationdate, Category) VALUES ( ?, ?, ?, ?, ?, ?, ? );")
	formSelect := r.PostForm.Get("choice")
	formText := r.PostForm.Get("Usertxt")

	stmt.Exec("Yasser@test.com", formText, 0, 0, "", time.Now(), formSelect)
	fmt.Println("here", formText, formSelect)
	ShowBdd()
	http.Redirect(w, r, "/homeLogged", 301)

}
func AddComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stmt, _ := Database.Prepare("UPDATE Post SET Comment='' ||Comment where [Id-Post] = 1")

	formText := r.PostForm.Get("myInput")

	stmt.Exec(formText)
	fmt.Println("here", formText)
	ShowBdd()
	http.Redirect(w, r, "/homeLogged", 301)

}
func Home(w http.ResponseWriter, req *http.Request) {

	tHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHome.Execute(w, Result)
}

func HomeLogged(w http.ResponseWriter, req *http.Request) {

	tHomeLogged, err := template.ParseFiles("templates/homeLogged.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tHomeLogged.Execute(w, Result)
}

func Dashboard(w http.ResponseWriter, req *http.Request) {
	tDashboard, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		w.WriteHeader(400)
	}

	tDashboard.Execute(w, Result)
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
