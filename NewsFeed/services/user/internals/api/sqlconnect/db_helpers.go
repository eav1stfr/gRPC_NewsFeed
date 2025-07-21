package sqlconnect

import (
	"context"
	"github.com/jmoiron/sqlx"
	pb "proto/user/gen"
	"utils"
)

func checkUniqueEmailAndUsername(ctx context.Context, db *sqlx.DB, user *pb.RegisterRequest) error {
	var emailExists, usernameExists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1) AS email_exists,EXISTS (SELECT 1 FROM users WHERE username = $2) AS username_exists"

	err := db.QueryRowContext(ctx, query, user.Email, user.Username).Scan(&emailExists, &usernameExists)
	if err != nil {
		return utils.DatabaseQueryError
	}

	if emailExists {
		return utils.EmailAlreadyInUse
	}
	if usernameExists {
		return utils.UsernameAlreadyInUse
	}

	return nil
}
