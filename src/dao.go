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
func AddPost(r *http.Request) {
	r.ParseForm()
	stmt, _ := Database.Prepare("INSERT INTO Post( User, Content, Like, Dislike, Creationdate, Category) VALUES ( ?, ?, ?, ?, ?, ? );")
	formSelect := r.PostForm.Get("choice")
	formText := r.PostForm.Get("Usertxt")

	stmt.Exec("Yasser@test.com", formText, 0, 0, time.Now(), formSelect)
	fmt.Println("here", formText, formSelect)
	stmt.Close()
	GetPosts()
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
	fmt.Println("Get:", formText)
	fmt.Println("GetID:", GetId)
	stmt.Close()
	GetComments()
}

// Add an account to DB
func AddAccount(newAccount Account) {
	statement, _ := Database.Prepare("INSERT INTO Account (name, email, hashPwd) VALUES(?, ?, ?)")
	statement.Exec(newAccount.Name, newAccount.Email, newAccount.HashPwd)
	statement.Close()
}

// ---READ---

// Formerly ShowPost()
// Get all posts from DB
func GetPosts() {
	rows, _ := Database.Query("SELECT [Id-Post], User, Content, Like, Dislike, CreationDate, Category FROM Post")

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
		AllData[i].Comments = data
	}
}

// Get all accounts from DB
func GetAccounts() {
	rows, _ := Database.Query("SELECT * FROM Account")
	defer rows.Close()

	for rows.Next() {
		var account Account
		rows.Scan(&account.Email, &account.Name, &account.HashPwd, &account.SessionUUID)
		Accounts = append(Accounts, account)
	}
}

// Get all sessions from DB
func GetSessions() {
	rows, _ := Database.Query("SELECT * FROM Session")
	defer rows.Close()
	for rows.Next() {
		var session Session
		rows.Scan(&session.SessionUUID, &session.UserID)
		Sessions = append(Sessions, session)
	}
}

// ---UPDATE---

// ---DELETE---

// Delete session in DB corresponding to userID specified
func DeleteSession(userID string) {
	stmt, _ := Database.Prepare("DELETE FROM Session WHERE userID = ?;")
	stmt.Exec(userID)
	stmt.Close()
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
