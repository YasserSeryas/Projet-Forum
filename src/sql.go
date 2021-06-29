package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Database, _ = sql.Open("sqlite3", "./Bdd/ProjetForumBDD.db")
var Likes = []Like{}
var Result1 = []Comment{}
var Result2 = []TemplateData{}

/*func ShowComment() {
	rows, _ := Database.Query("SELECT [Id-Comment], [Id-Post], [Id-User], [Comment-Content] FROM Comment;")
	var data Comment
	for rows.Next() {
		rows.Scan(&data.IdComment, &data.IdPost, &data.IdUser, &data.CommentContent)

		Result1 = append(Result1, data)

	}
	rows.Close()
}*/
func GetComment() {
	for i, val := range Result2 {
		Id := val.PostData.IdPost
		fmt.Println("Id=", Id)
		stmt, _ := Database.Prepare("SELECT [Id-Comment], [Id-Post], [Id-User], [Comment-Content], UserName FROM Comment WHERE [Id-Post] = ?  ")
		rows, _ := stmt.Query(Id)

		var data []Comment

		for rows.Next() {
			var Comments Comment
			rows.Scan(&Comments.IdComment, &Comments.IdPost, &Comments.IdUser, &Comments.CommentContent, &Comments.UserName)
			data = append(data, Comments)
			val.Comments = data

			fmt.Println(data)
		}
		Result2[i].Comments = data

	}

}
func ShowPost() {
	rows, _ := Database.Query("SELECT [Id-Post], User, Content, Like, Dislike, CreationDate, Category FROM Post")

	var data Post
	var TempData TemplateData
	for rows.Next() {
		rows.Scan(&data.IdPost, &data.User, &data.Content, &data.Like, &data.Dislike, &data.CreationDate, &data.Category)
		TempData.PostData = data
		Result2 = append(Result2, TempData)

	}
	rows.Close()
}

//Insert a post
func Insert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stmt, _ := Database.Prepare("INSERT INTO Post( User, Content, Like, Dislike, Creationdate, Category) VALUES ( ?, ?, ?, ?, ?, ? );")
	formSelect := r.PostForm.Get("choice")
	formText := r.PostForm.Get("Usertxt")

	stmt.Exec("Yasser@test.com", formText, 0, 0, time.Now(), formSelect)
	fmt.Println("here", formText, formSelect)
	stmt.Close()
	ShowPost()

	http.Redirect(w, r, "/homeLogged", http.StatusMovedPermanently)

}

//Add a Comment
func AddComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var UserName string
	stmtSelect, _ := Database.Prepare("SELECT Name FROM Account WHERE Email = ?  ")
	row := stmtSelect.QueryRow("ValeurBrute")
	err := row.Scan(&UserName)
	if err != nil {
		log.Fatal(err)
	}

	stmt, _ := Database.Prepare("INSERT INTO Comment( [Id-Post], [Id-User], [Comment-Content], UserName) VALUES ( ?, ?, ?, ? );")

	formText := r.PostForm.Get("Usertxt")
	GetId := r.PostForm.Get("Idpost")
	stmt.Exec(GetId, "User", formText, UserName)
	fmt.Println("Get:", formText)
	fmt.Println("GetID:", GetId)
	stmt.Close()
	ShowPost()

	http.Redirect(w, r, "/homeLogged", http.StatusMovedPermanently)

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stmt, err := Database.Prepare("SELECT  User, Content, Like, Dislike, Creationdate, Category FROM Post WHERE Category = ?;")
	Select := r.PostForm.Get("Category")
	rows, _ := stmt.Query(Select)
	var data Post
	if err != nil {
		fmt.Println("error Query")
	}
	fmt.Println(Select)
	for rows.Next() {
		rows.Scan(&data.User, &data.Content, &data.Like, &data.Dislike, &data.CreationDate, &data.Category)

	}
	rows.Close()
	fmt.Println(data)

	http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
}

// func AddLike(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	// Add like
// 	likeRows, err := Database.Prepare("UPDATE Post SET Like = Like + 1 Where [Id-Post] = ? ;")
// 	GetId := r.PostForm.Get("IdPostLike")
// 	likeRows.Exec(GetId)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 	}
// 	fmt.Println(GetId)
// 	likeRows.Close()

// 	http.Redirect(w, r, "/homeLogged", http.StatusMovedPermanently)

// count := 0
// for likeRows.Next() {
// 	count++
// }
// if count >= 1 {
// 	// Remove like if exists
// 	_, err = Database.Exec(`
// 		DELETE from likes WHERE User=? AND IdPost=?
// 	`, User, IdPost)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 	}
// 	//
// } else {
// 	// Ad like if not exists
// 	_, err = Database.Exec(`
// 		INSERT OR IGNORE INTO likes (User, IdPost) VALUES (?, ?)
// 	`, User, IdPost)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 	}
// }

// // Remove dislike
// _, err = Database.Exec(`
// 	DELETE from dislikes WHERE User=? AND IdPost=?
// `, User, IdPost)
// if err != nil {
// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// }
// }// func AddDislike(newDislike Like) {
// 	stmt, _ := Database.Prepare("INSERT INTO Like ( IdUser, IsLike, IdPost) VALUES ( ?, ?, ?); ")
// 	stmt.Exec(newDislike.IdUser, newDislike.IsLike, newDislike.IdPost)
// 	stmt.Close()
// 	GetDislike()

// }
func AddLike(newLike Like) {
	stmt, _ := Database.Prepare("INSERT INTO Like ( IdUser, IsLike, IdPost) VALUES ( ?, ?, ?); ")
	stmt.Exec(newLike.IdUser, newLike.IsLike, newLike.IdPost)
	stmt.Close()
	GetLike()
}

func DeleteLike(IdLike int) {
	stmt, _ := Database.Prepare("DELETE FROM Like WHERE IdLike = ?;")
	stmt.Exec(IdLike)
	stmt.Close()
	GetLike()
}
func UpdateLike(IsLike bool) {
	stmt, _ := Database.Prepare("UPDATE Like SET IsLike =?;")
	stmt.Exec(IsLike)
	stmt.Close()
	GetLike()
}
func GetLike() {
	var likes []Like
	rows, _ := Database.Query("SELECT IdLike, IdUser, IsLike, IdPost FROM Like ;")
	for rows.Next() {
		var like Like
		rows.Scan(&like.IdLike, &like.IdUser, &like.IsLike, &like.IdPost)
		likes = append(likes, like)
	}
	Likes = likes
	rows.Close()

}
