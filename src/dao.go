package src

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var Database, _ = sql.Open("sqlite3", "./Bdd/ProjetForumBDD.db")

// CRUD : Create Read Update Delete

// ---CREATE---

// Add a post to DB
func AddPost(newPost Post) {
	stmt, _ := Database.Prepare("INSERT INTO Post( User, Title, Content, NbrLike, NbrDislike, Creationdate, Category) VALUES ( ?, ?, ?, ?, ?, ?, ? );")
	stmt.Exec(newPost.User, newPost.Title, newPost.Content, 0, 0, time.Now(), newPost.Category)
	stmt.Close()
	GetPosts() // Update struct go variable
}

// Add a comment to DB | A FINIR | <-----
func AddComment(newComment Comment) {
	stmt, _ := Database.Prepare("INSERT INTO Comment( [Id-Post], [Id-User], [Comment-Content], UserName) VALUES ( ?, ?, ?, ? );")
	stmt.Exec(newComment.IdPost, newComment.IdUser, newComment.CommentContent, newComment.UserName)
	stmt.Close()
	GetComments() // Update struct go variable
}

// Add an account to DB
func AddAccount(newAccount Account) {
	statement, _ := Database.Prepare("INSERT INTO Account (name, email, hashPwd) VALUES(?, ?, ?)")
	statement.Exec(newAccount.Name, newAccount.Email, newAccount.HashPwd)
	statement.Close()
	GetAccounts() // Update struct go variable
}

// Add a session to DB
func AddSession(newSession Session) {
	statement, _ := Database.Prepare("INSERT INTO Session (SessionUUID, UserID) VALUES (?, ?)")
	statement.Exec(newSession.SessionUUID, newSession.UserID)
	statement.Close()
	GetSessions() // Update struct go variable
}

func AddLike(newLike Like) {
	fmt.Println(newLike)
	stmt, errBDD := Database.Prepare("INSERT INTO Like ( User, IsLike, IdPost) VALUES ( ?, ?, ?); ")
	if errBDD != nil {
		fmt.Println(errBDD)
	}
	_, errExec := stmt.Exec(newLike.IdUser, newLike.IsLike, newLike.IdPost)
	if errExec != nil {
		fmt.Println(errExec)
	}
	stmt.Close()
	GetLike()
}

// ---READ---

// Formerly ShowPost()
// Get all posts from DB
func GetPosts() {
	rows, _ := Database.Query("SELECT [Id-Post], User, Title, Content, NbrLike, NbrDislike, CreationDate, Category FROM Post")

	var data Post
	var TempData TemplateData
	for rows.Next() {
		rows.Scan(&data.IdPost, &data.User, &data.Title, &data.Content, &data.Like, &data.Dislike, &data.CreationDate, &data.Category)
		TempData.PostData = data
		AllData = append(AllData, TempData)

	}
	rows.Close()
}

// Get all comments from DB
func GetComments() {
	for i, val := range AllData {
		Id := val.PostData.IdPost
		stmt, _ := Database.Prepare("SELECT [Id-Comment], [Id-Post], [Id-User], [Comment-Content], UserName FROM Comment WHERE [Id-Post] = ?  ")
		rows, _ := stmt.Query(Id)
		stmt.Close()

		var data []Comment

		for rows.Next() {
			var Comments Comment
			rows.Scan(&Comments.IdComment, &Comments.IdPost, &Comments.IdUser, &Comments.CommentContent, &Comments.UserName)
			data = append(data, Comments)
			val.Comments = data

			// fmt.Println(data)
		}
		rows.Close()
		AllData[i].Comments = data
	}
}

// Get all accounts from DB
func GetAccounts() {
	rows, _ := Database.Query("SELECT * FROM Account")

	for rows.Next() {
		var account Account
		rows.Scan(&account.Email, &account.Name, &account.HashPwd)
		Accounts = append(Accounts, account)
	}
	rows.Close()
}

// Get all sessions from DB
func GetSessions() {
	rows, _ := Database.Query("SELECT * FROM Session")
	var sessions []Session
	for rows.Next() {
		var session Session
		rows.Scan(&session.SessionUUID, &session.UserID)
		sessions = append(Sessions, session)
	}
	Sessions = sessions
	rows.Close()
}

func GetUsername(IdUser string) string {
	var username string

	stmt, _ := Database.Prepare("SELECT Name FROM Account WHERE Email = ?")
	row := stmt.QueryRow(IdUser)
	row.Scan(&username)
	stmt.Close()

	return username
}

func GetLike() {
	var likes []Like
	rows, _ := Database.Query("SELECT IdLike, User, IsLike, IdPost FROM Like ;")
	for rows.Next() {
		var like Like
		rows.Scan(&like.IdLike, &like.IdUser, &like.IsLike, &like.IdPost)
		likes = append(likes, like)
	}
	Likes = likes
	rows.Close()
}

// ---UPDATE---

func UpdateLike(IsLike bool) {
	stmt, _ := Database.Prepare("UPDATE Like SET IsLike =?;")
	stmt.Exec(IsLike)
	stmt.Close()
	GetLike()
}

// ---DELETE---

// Delete session in DB corresponding to userID specified
func DeleteSession(userID string) {
	stmt, _ := Database.Prepare("DELETE FROM Session WHERE userID = ?;")
	stmt.Exec(userID)
	stmt.Close()
	GetSessions() // Update struct go variable
}

func DeleteLike(IdLike int) {
	stmt, _ := Database.Prepare("DELETE FROM Like WHERE IdLike = ?;")
	stmt.Exec(IdLike)
	stmt.Close()
	GetLike()
}

// ---OTHER---

// Print in terminal accounts informations
func PrintBDD() {
	result, errSelect := Database.Query("SELECT name, email, hashPwd FROM Account")
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
