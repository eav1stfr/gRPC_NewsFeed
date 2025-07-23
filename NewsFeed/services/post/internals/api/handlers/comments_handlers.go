package handlers

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"post/internals/api/sqlconnect"
	pb "proto/post/gen"
	"utils"
)

func (s *Server) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	if req.PostId == 0 || req.UserId == 0 || req.Content == "" {
		return nil, utils.ToGRPCStatus(utils.InvalidRequestPayload)
	}
	commentId, err := sqlconnect.AddCommentDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.AddCommentResponse{CommentId: commentId}, nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	if req.PostId == 0 {
		return nil, utils.ToGRPCStatus(utils.InvalidRequestPayload)
	}
	comments, err := sqlconnect.GetCommentsDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.GetPostCommentsResponse{Comments: comments}, nil
}

func (s *Server) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*emptypb.Empty, error) {
	if req.CommentId == 0 || req.PostId == 0 || req.UpdatedContent == "" {
		return nil, utils.ToGRPCStatus(utils.InvalidRequestPayload)
	}
	err := sqlconnect.UpdateCommentDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	if req.PostId == 0 || req.CommentId == 0 {
		return nil, utils.ToGRPCStatus(utils.InvalidRequestPayload)
	}
	err := sqlconnect.DeleteCommentDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &emptypb.Empty{}, nil
}
