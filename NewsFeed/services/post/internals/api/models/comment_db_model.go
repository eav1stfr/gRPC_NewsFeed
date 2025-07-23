package models

type Comment struct {
	PostId    int    `db:"post_id"`
	CommentId int    `db:"comment_id"`
	UserId    int    `db:"user_id"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}
