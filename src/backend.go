package helpers

import (
	"net/http"
	"strconv"
)

func CreateLike(req *http.Request) {
	var IdUser string = "test"
	IDPost, _ := strconv.Atoi(req.FormValue("idPost"))
	for _, like := range Likes {
		if like.IdPost == IDPost && like.IdUser == IdUser {
			if like.IsLike {
				DeleteLike(like.IdLike)
				return
			} else {
				UpdateLike(true)
				return
			}
		}
	}
	var newLike Like
	newLike.IdPost = IDPost
	newLike.IdUser = IdUser
	newLike.IsLike = true
	AddLike(newLike)
}
func CreateDislike(req *http.Request) {
	var IdUser string = "test"
	IDPost, _ := strconv.Atoi(req.FormValue("idPost"))
	for _, like := range Likes {
		if like.IdPost == IDPost && like.IdUser == IdUser {
			if like.IsLike {
				UpdateLike(false)
				return
			} else {
				DeleteLike(like.IdLike)
				return
			}
		}
	}
	var newDislike Like
	newDislike.IdPost, _ = strconv.Atoi(req.FormValue("idPost"))
	newDislike.IdUser = "test"
	newDislike.IsLike = false
	AddLike(newDislike)
}
