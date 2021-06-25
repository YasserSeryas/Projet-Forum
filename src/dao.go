package src

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Database, _ = sql.Open("sqlite3", "./Bdd/ProjetForumBDD.db")

// ---CREATE---

// Add a post to DB
func AddPost(newPost Post) {
	stmt, _ := Database.Prepare("INSERT INTO Post( User, Content, NbrLike, NbrDislike, Creationdate, Category) VALUES ( ?, ?, ?, ?, ?, ? );")

	stmt.Exec(newPost.User, newPost.Content, 0, 0, time.Now(), newPost.Category)
	stmt.Close()
	GetPosts() // Update struct go variable
}

// Add a comment to DB | A FINIR | <-----
func AddComment(r *http.Request) {
	r.ParseForm()
	var UserName string
	stmtSelect, _ := Database.Prepare("SELECT Name FROM Account WHERE Email = ?  ")
	defer stmtSelect.Close()
	row := stmtSelect.QueryRow("ValeurBrute")
	err := row.Scan(&UserName)
	if err != nil {
		log.Fatalln("In AddComment :", err)
	}

	stmt, _ := Database.Prepare("INSERT INTO Comment( [Id-Post], [Id-User], [Comment-Content], UserName) VALUES ( ?, ?, ?, ? );")

	formText := r.PostForm.Get("Usertxt")
	GetId := r.PostForm.Get("Idpost")
	stmt.Exec(GetId, "ValeurBrute", formText, UserName)
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

// ---READ---

// Formerly ShowPost()
// Get all posts from DB
func GetPosts() {
	rows, _ := Database.Query("SELECT [Id-Post], User, Content, NbrLike, NbrDislike, CreationDate, Category FROM Post")

	var data Post
	var TempData TemplateData
	for rows.Next() {
		rows.Scan(&data.IdPost, &data.User, &data.Content, &data.Like, &data.Dislike, &data.CreationDate, &data.Category)
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

		var data []Comment

		for rows.Next() {
			var Comments Comment
			rows.Scan(&Comments.IdComment, &Comments.IdPost, &Comments.IdUser, &Comments.CommentContent, &Comments.UserName)
			data = append(data, Comments)
			val.Comments = data

			// fmt.Println(data)
		}
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

// ---UPDATE---

// ---DELETE---

// Delete session in DB corresponding to userID specified
func DeleteSession(userID string) {
	stmt, _ := Database.Prepare("DELETE FROM Session WHERE userID = ?;")
	stmt.Exec(userID)
	stmt.Close()
	GetSessions() // Update struct go variable
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
