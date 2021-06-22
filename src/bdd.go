package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, err := sql.Open("sqlite3", "../Bdd/ProjetForumBDD.db")
	if err != nil {
		fmt.Print("Error")
	}

	rows, _ := database.Query("SELECT User, Content, Like, Dislike, Comment , CreationDate, Category FROM Post  ")
	var data Post
	result := []Post{}
	for rows.Next() {
		rows.Scan(&data.User, &data.Content, &data.Like, &data.Dislike, &data.Comment, &data.CreationDate, &data.Category)
		//dat := (" User : " + data.User + "\n") + (" Content :" + data.Content + "\n") + (" Like : " + strconv.Itoa(data.Like) + "\n") + (" Dislike : " + strconv.Itoa(data.Dislike) + "\n") + (" Comment : " + data.Comment + "\n") + (" Creation Date : " + data.CreationDate.String() + "\n") + (" Category : " + data.Category + "\n")

		result = append(result, data)
	}
	fmt.Println(result)
}
