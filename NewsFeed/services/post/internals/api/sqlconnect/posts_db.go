package sqlconnect

import (
	"context"
	"post/internals/api/models"
	pb "proto/post/gen"
	"strconv"
	"utils"
)

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
