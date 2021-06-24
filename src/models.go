package helpers

import "time"

type Account struct {
	Email       string
	Name        string
	HasgPwd     string
	SessionUUID string
}

type Category struct {
	Name  string
	Color int
}

type Post struct {
	IdPost       int
	User         string
	Content      string
	Like         int
	Dislike      int
	CreationDate time.Time
	Category     string
}
type Comment struct {
	IdComment      int
	IdPost         int
	IdUser         string
	CommentContent string
	UserName       string
}
type Like struct {
	IdPost    int
	User      string
	IsLike    bool
	IsDislike bool
	Date      time.Time
}
type TemplateData struct {
	PostData Post
	Comments []Comment
}
