package sqlconnect

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"post/internals/api/models"
	pb "proto/post/gen"
	"utils"
)

func AddCommentDbHandler(ctx context.Context, req *pb.AddCommentRequest) (int32, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, utils.ConnectingToDbError
	}
	defer db.Close()

	commentId, err := getMaxCommentId(ctx, db, int(req.PostId))
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO comments (post_id, comment_id, user_id, content) VALUES ($1, $2, $3, $4)"

	_, err = db.ExecContext(ctx, query, req.PostId, commentId, req.UserId, req.Content)
	if err != nil {
		return 0, utils.DatabaseQueryError
	}
	return int32(commentId), nil
}

func getMaxCommentId(ctx context.Context, db *sqlx.DB, postId int) (int, error) {
	var id sql.NullInt64
	query := "SELECT MAX(comment_id) FROM comments WHERE post_id = $1"

	err := db.GetContext(ctx, &id, query, postId)
	if err != nil {
		return 0, utils.DatabaseQueryError
	}

	if !id.Valid {
		return 1, nil
	}
	return int(id.Int64 + 1), nil
}

func GetCommentsDbHandler(ctx context.Context, req *pb.GetPostCommentsRequest) ([]*pb.Comment, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var comments []models.Comment
	query := "SELECT * FROM comments WHERE post_id = $1"
	err = db.SelectContext(ctx, &comments, query, req.PostId)
	pbComments := make([]*pb.Comment, 0, len(comments))
	for _, comm := range comments {
		pbComm := &pb.Comment{
			Id:        int32(comm.CommentId),
			PostId:    int32(comm.PostId),
			CreatorId: int32(comm.UserId),
			Content:   comm.Content,
			CreatedAt: comm.CreatedAt,
		}
		pbComments = append(pbComments, pbComm)
	}
	return pbComments, nil
}

func UpdateCommentDbHandler(ctx context.Context, req *pb.UpdateCommentRequest) error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE comments SET content = $1 WHERE comment_id = $2 AND post_id = $3"
	res, err := db.ExecContext(ctx, query, req.UpdatedContent, req.CommentId, req.PostId)
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

func DeleteCommentDbHandler(ctx context.Context, req *pb.DeleteCommentRequest) error {
	db, err := ConnectDb()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM comments WHERE comment_id = $1 AND post_id = $2"
	res, err := db.ExecContext(ctx, query, req.CommentId, req.PostId)
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
