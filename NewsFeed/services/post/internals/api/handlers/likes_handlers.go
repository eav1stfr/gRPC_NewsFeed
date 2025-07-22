package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "proto/post/gen"
)

func (s *Server) LikePost(ctx context.Context, req *pb.LikePostRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) UnlikePost(ctx context.Context, req *pb.UnlikePostRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *Server) GetPostLikes(ctx context.Context, req *pb.GetPostLikesRequest) (*pb.GetPostLikesResponse, error) {
	return nil, nil
}
