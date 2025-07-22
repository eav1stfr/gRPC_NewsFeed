package handlers

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "proto/post/gen"
)

func (s *Server) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	return nil, nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	return nil, nil
}

func (s *Server) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) DeleteComment(ctx context.Context, pb *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	return nil, nil
}
