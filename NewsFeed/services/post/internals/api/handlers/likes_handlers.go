package handlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"post/internals/api/sqlconnect"
	pb "proto/post/gen"
	"utils"
)

func (s *Server) LikePost(ctx context.Context, req *pb.LikePostRequest) (*emptypb.Empty, error) {
	err := sqlconnect.LikePostDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return nil, nil
}

func (s *Server) UnlikePost(ctx context.Context, req *pb.UnlikePostRequest) (*emptypb.Empty, error) {
	err := sqlconnect.UnlikePostDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return nil, nil
}

func (s *Server) GetPostLikes(ctx context.Context, req *pb.GetPostLikesRequest) (*pb.GetPostLikesResponse, error) {
	userIds, err := sqlconnect.GetPostLikesDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}

	return &pb.GetPostLikesResponse{UserIds: userIds}, nil
}
