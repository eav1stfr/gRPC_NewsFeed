package sqlconnect

import (
	"context"
	"database/sql"
	pb "proto/post/gen"
	"strconv"
	"utils"
)

func LikePostDbHandler(ctx context.Context, req *pb.LikePostRequest) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ConnectingToDbError
	}
	defer db.Close()

	query := "INSERT INTO post_likes (post_id, user_id) VALUES ($1, $2)"
	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return utils.InvalidRequestPayload
	}
	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return utils.InvalidRequestPayload
	}
	_, err = db.ExecContext(ctx, query, postId, userId)
	if err != nil {
		return utils.DatabaseQueryError
	}
	return nil
}

func UnlikePostDbHandler(ctx context.Context, req *pb.UnlikePostRequest) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ConnectingToDbError
	}
	defer db.Close()

	query := "DELETE FROM post_likes WHERE post_id = $1 AND user_id = $2"

	userId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return utils.InvalidRequestPayload
	}
	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return utils.InvalidRequestPayload
	}
	res, err := db.ExecContext(ctx, query, postId, userId)
	if err != nil {
		return utils.DatabaseQueryError
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return utils.DatabaseQueryError
	}
	if rowsAffected == 0 {
		return utils.InvalidRequestPayload
	}
	return nil
}

func GetPostLikesDbHandler(ctx context.Context, req *pb.GetPostLikesRequest) ([]string, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ConnectingToDbError
	}
	defer db.Close()

	var userIds []int

	query := "SELECT user_id FROM post_likes WHERE post_id = $1"

	postId, err := strconv.Atoi(req.PostId)
	if err != nil {
		return nil, utils.InvalidRequestPayload
	}

	err = db.SelectContext(ctx, &userIds, query, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, utils.DatabaseQueryError
	}

	likes := make([]string, len(userIds))
	for i, id := range userIds {
		likes[i] = strconv.Itoa(id)
	}

	return likes, nil
}
