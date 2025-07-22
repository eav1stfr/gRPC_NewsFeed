package sqlconnect

import (
	"post/internals/api/models"
	pb "proto/post/gen"
	"strconv"
	"time"
)

func convertDbPostToGrpc(post models.Post) *pb.Post {
	return &pb.Post{
		PostId:       strconv.Itoa(int(post.Id)),
		UserId:       strconv.Itoa(int(post.UserId)),
		Content:      post.Content,
		ImageUrl:     post.ImageUrl,
		CreatedAt:    post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    post.UpdatedAt.Format(time.RFC3339),
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
	}
}

func convertDbPostListToGrpcList(posts []models.Post) []*pb.Post {
	res := make([]*pb.Post, 0, len(posts))
	for _, post := range posts {
		grpcPost := convertDbPostToGrpc(post)
		res = append(res, grpcPost)
	}
	return res
}
