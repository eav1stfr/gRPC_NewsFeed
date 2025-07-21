package sqlconnect

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	pb "proto/user/gen"
	"user/internals/api/models"
	"utils"
)

func RegisterUserDbHandler(ctx context.Context, user *pb.RegisterRequest) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	err = checkUniqueEmailAndUsername(ctx, db, user)
	if err != nil {
		return 0, err
	}

	passwordHash, err := utils.Hash(user.Password)
	if err != nil {
		return 0, err
	}

	var id int
	query := "INSERT INTO users (first_name, last_name, email, password_hash, username) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Email, passwordHash, user.Username).Scan(&id)
	if err != nil {
		return 0, utils.DatabaseQueryError
	}
	return id, nil
}

func GetUserProfileDbHandler(ctx context.Context, req *pb.GetProfileRequest) (*pb.User, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.UserModel
	query := "SELECT id, first_name, last_name, email, username, avatar_url, bio, created_at FROM users WHERE id = $1"
	err = db.GetContext(ctx, &user, query, req.Id)
	if err != nil {
		return nil, utils.DatabaseQueryError
	}
	res := &pb.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		AvatarUrl: user.AvatarUrl,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt}
	return res, nil
}

//CREATE TABLE follows (
//	follower_id INT NOT NULL,
//	followed_id INT NOT NULL,
//	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//
//	PRIMARY KEY (follower_id, followed_id),
//
//	FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
//	FOREIGN KEY (followed_id) REFERENCES users(id) ON DELETE CASCADE,
//
//	CHECK (follower_id <> followed_id)
//);

//message FollowRequest {
//	string follower_id = 1;
//	string followee_id = 2;
//}

func FollowDbHandler(ctx context.Context, req *pb.FollowRequest) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ConnectingToDbError
	}
	defer db.Close()

	if req.FollowerId == req.FolloweeId {
		return utils.EqualIdsError
	}

	query := "INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2)"
	_, err = db.Exec(query, req.FollowerId, req.FolloweeId)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return utils.AlreadyFollowingError
		}
		return utils.DatabaseQueryError
	}
	return nil
}

func UnfollowDbHandler(ctx context.Context, req *pb.UnfollowRequest) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ConnectingToDbError
	}
	defer db.Close()

	if req.FolloweeId == req.FollowerId {
		return utils.NotFollowingError
	}

	query := "DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2"
	res, err := db.ExecContext(ctx, query, req.FollowerId, req.FolloweeId)
	if err != nil {
		return utils.DatabaseQueryError
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.DatabaseQueryError
	}
	if rowsAffected == 0 {
		return utils.NotFollowingError
	}
	return nil
}

//message FollowListRequest {
//	string user_id = 1;
//}
//
//message FollowListResponse {
//	repeated User users = 1;
//}

func GetFollowingListDbHandler(ctx context.Context, req *pb.FollowListRequest, followers bool) ([]*pb.User, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, utils.ConnectingToDbError
	}
	defer db.Close()

	var query string

	if followers {
		query = "SELECT u.id, u.first_name, u.last_name, u.last_name, u.email, u.username, u.avatar_url, u.bio, u.created_at " +
			"FROM follows f " +
			"JOIN users u ON u.id = f.follower_id " +
			"WHERE f.followed_id = $1"
	} else {
		query = "SELECT u.id, u.first_name, u.last_name, u.last_name, u.email, u.username, u.avatar_url, u.bio, u.created_at " +
			"FROM follows f " +
			"JOIN users u ON u.id = f.followed_id " +
			"WHERE f.follower_id = $1"
	}

	rows, err := db.QueryContext(ctx, query, req.UserId)
	if err != nil {
		return nil, utils.DatabaseQueryError
	}

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.AvatarUrl, &user.Bio, &user.CreatedAt)
		if err != nil {
			return nil, utils.DatabaseQueryError
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, utils.DatabaseQueryError
	}
	return users, nil
}

//message UpdateProfileRequest {
//	string user_id = 1;
//	optional string username = 2;
//	optional string avatar_url = 3;
//	optional string bio = 4;
//}

func UpdateProfileDbHandler(ctx context.Context, req *pb.UpdateProfileRequest) error {
	db, err := ConnectDB()
	if err != nil {
		return utils.ConnectingToDbError
	}
	defer db.Close()
	request := &pb.GetProfileRequest{Id: req.UserId}
	user, err := GetUserProfileDbHandler(ctx, request)
	if err != nil {
		return err
	}
	var updates = make(map[string]interface{})
	if req.Username != nil {
		if *req.Username == user.Username {
			return utils.NoChangeAppliedError
		}
		updates["username"] = *req.Username
	}
	if req.Bio != nil {
		updates["bio"] = *req.Bio
	}
	if req.AvatarUrl != nil {
		updates["avatar_url"] = *req.AvatarUrl
	}
	if len(updates) == 0 {
		return utils.NoChangeAppliedError
	}

	args := make([]interface{}, 0, len(updates))
	query := "UPDATE users SET "
	i := 1
	for key, val := range updates {
		if i > 1 {
			query += ", "
		}
		query += fmt.Sprintf("%s = $%d", key, i)
		args = append(args, val)
		i++
	}
	query += fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, req.UserId)

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return utils.DatabaseQueryError
	}
	return nil
}
