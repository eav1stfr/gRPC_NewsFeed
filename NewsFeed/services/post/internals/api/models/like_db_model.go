package models

type PostLikes struct {
	PostId int `db:"post_id"`
	UserId int `db:"user_id"`
}
