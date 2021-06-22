package helpers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

var Database, _ = sql.Open("sqlite3", "./Bdd/ProjetForumBDD.db")
var Result = []Post{}

func ShowBdd() {
	rows, _ := Database.Query("SELECT [Id-Post], User, Content, Like, Dislike, Comment , CreationDate, Category FROM Post")

	var data Post

	for rows.Next() {
		rows.Scan(&data.IdPost, &data.User, &data.Content, &data.Like, &data.Dislike, &data.Comment, &data.CreationDate, &data.Category)

		Result = append(Result, data)
	}
	rows.Close()
}

//Insert a post
func Insert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stmt, _ := Database.Prepare("INSERT INTO Post( User, Content, Like, Dislike, Comment, Creationdate, Category) VALUES ( ?, ?, ?, ?, ?, ?, ? );")
	formSelect := r.PostForm.Get("choice")
	formText := r.PostForm.Get("Usertxt")

	stmt.Exec("Yasser@test.com", formText, 0, 0, "", time.Now(), formSelect)
	fmt.Println("here", formText, formSelect)
	ShowBdd()
	http.Redirect(w, r, "/homeLogged", http.StatusMovedPermanently)

}

//Add a Comment
func AddComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stmt, _ := Database.Prepare("UPDATE Post SET Comment=?||Comment where [Id-Post] = 1")

	formText := r.PostForm.Get("Usertxt") + "  "

	stmt.Exec(formText)
	fmt.Println("Get:", formText)

	ShowBdd()
	http.Redirect(w, r, "/homeLogged", http.StatusMovedPermanently)

}
