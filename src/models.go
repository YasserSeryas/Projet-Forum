package src

import "time"

type (
	Account struct {
		Email   string
		Name    string
		HashPwd string
	}

	Session struct {
		SessionUUID string
		UserID      string
	}

	Post struct {
		IdPost       int
		User         string
		Title        string
		Content      string
		Like         int
		Dislike      int
		CreationDate time.Time
		Category     string
	}

	Comment struct {
		IdComment      int
		IdPost         int
		IdUser         string
		CommentContent string
		UserName       string
	}

	Category struct {
		Name  string
		Color int
	}

	Like struct {
		IdLike int
		IdPost int
		IdUser string
		IsLike bool
	}

	TemplateData struct {
		PostData Post
		Comments []Comment
	}
)
