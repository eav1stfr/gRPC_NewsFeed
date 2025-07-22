package sqlconnect

import (
	"context"
	"github.com/jmoiron/sqlx"
	"post/internals/api/handlers"
	"post/internals/api/models"
	pb "proto/post/gen"
	"strconv"
	"utils"
)

//message Post {
//	string post_id = 1;
//	string user_id = 2; // author id
//	string content = 3;
//	string image_url = 4;
//	string created_at = 5;
//	string updated_at = 6;
//	int32 like_count = 7;
//	int32 comment_count = 8;
//}

//message CreatePostRequest {
//	string user_id = 1;
//	string content = 2;
//	string image_url = 3;
//}

// message CreatePostResponse {
//	string post_id = 1;
//}

func CreatePostDbHandler(ctx context.Context, req *pb.CreatePostRequest) (string, error) {
	db, err := ConnectDb()
	if err != nil {
		return "", err
	}
	defer db.Close()

	query := "INSERT INTO posts (user_id, content, image_url) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err = db.QueryRowContext(ctx, query, req.UserId, req.Content, req.ImageUrl).Scan(&id)
	if err != nil {
		return "", utils.DatabaseQueryError
	}

	return strconv.Itoa(id), nil
}

//message GetPostRequest {
//	string post_id = 1;
//}
//
//message GetPostResponse {
//	Post post = 1;
//}

//type Post struct {
//	Id           int32     `db:"id"`
//	UserId       int32     `db:"user_id"`
//	Content      string    `db:"content"`
//	ImageUrl     string    `db:"image_url"`
//	LikeCount    int32     `db:"like_count"`
//	CommentCount int32     `db:"comment_count"`
//	CreatedAt    time.Time `db:"created_at"`
//	UpdatedAt    time.Time `db:"updated_at"`
//}

func GetPostDbHandler(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var post models.Post
	query := "SELECT id, user_id, content, image_url, like_count, comment_count, created_at, updated_at FROM posts WHERE id = $1"
	err = db.GetContext(ctx, &post, query, req.PostId)
	if err != nil {
		return nil, utils.DatabaseQueryError
	}
	return convertDbPostToGrpc(post), nil
}

//message GetUserPostsRequest {
//	string user_id =1;
//}

//message GetUserPostsResponse {
//	repeated Post posts = 1;
//}

//type Post struct {
//	Id           int32     `db:"id"`
//	UserId       int32     `db:"user_id"`
//	Content      string    `db:"content"`
//	ImageUrl     string    `db:"image_url"`
//	LikeCount    int32     `db:"like_count"`
//	CommentCount int32     `db:"comment_count"`
//	CreatedAt    time.Time `db:"created_at"`
//	UpdatedAt    time.Time `db:"updated_at"`
//}

func GetUserPostsDbHandler(ctx context.Context, req *pb.GetUserPostsRequest) ([]*pb.Post, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, utils.InvalidRequestPayload
	}

	var posts []models.Post
	query := "SELECT * FROM posts WHERE user_id = $1"
	err = db.SelectContext(ctx, &posts, query, userId)
	if err != nil {
		return nil, utils.DatabaseQueryError
	}
	return convertDbPostListToGrpcList(posts), nil
}

//message DeletePostRequest {
//	string post_id = 1;
//	string user_id = 2;
//}

func DeletePostDbHandler(ctx context.Context, req *pb.DeletePostRequest) error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM posts WHERE user_id = $1 AND id = $2"
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return utils.InvalidRequestPayload
	}
	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return utils.InvalidRequestPayload
	}
	res, err := db.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return utils.DatabaseQueryError
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.DatabaseQueryError
	}
	if rowsAffected == 0 {
		return utils.DatabaseQueryError
	}
	return nil
}
