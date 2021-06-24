package src

import "time"

type (
	Account struct {
		Email       string
		Name        string
		HashPwd     string
		SessionUUID string
	}

	Session struct {
		SessionUUID string
		UserID      string
	}

	Post struct {
		IdPost       int
		User         string
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
		IdPost    int
		User      string
		IsLike    bool
		IsDislike bool
		Date      time.Time
	}

	TemplateData struct {
		PostData Post
		Comments []Comment
	}
)
