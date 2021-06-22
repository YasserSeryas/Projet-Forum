package helpers

import "time"

type Account struct {
	Email       string
	Name        string
	HasgPwd     string
	SessionUUID string
}

type Categorie struct {
	Name  string
	Color int
}

type Post struct {
	IdPost       int
	User         string
	Content      string
	Like         int
	Dislike      int
	Comment      string
	CreationDate time.Time
	Category     string
}

type Like struct {
	IdPost    int
	User      string
	IsLike    bool
	IsDislike bool
	Date      time.Time
}
