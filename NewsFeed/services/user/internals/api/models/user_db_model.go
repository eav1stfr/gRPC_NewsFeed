package models

type UserModel struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Id        string `db:"id"`
	Email     string `db:"email"`
	Username  string `db:"username"`
	AvatarUrl string `db:"avatar_url"`
	Bio       string `db:"bio"`
	CreatedAt string `db:"created_at"`
}
