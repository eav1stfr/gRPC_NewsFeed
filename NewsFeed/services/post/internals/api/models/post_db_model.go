package models

import "time"

type Post struct {
	Id           int32     `db:"id"`
	UserId       int32     `db:"user_id"`
	Content      string    `db:"content"`
	ImageUrl     string    `db:"image_url"`
	LikeCount    int32     `db:"like_count"`
	CommentCount int32     `db:"comment_count"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
