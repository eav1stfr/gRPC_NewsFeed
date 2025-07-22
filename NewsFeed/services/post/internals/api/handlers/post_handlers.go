package handlers

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"post/internals/api/sqlconnect"
	pb "proto/post/gen"
	"utils"
)

func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	if req.Content == "" || req.UserId == "" {
		return nil, utils.MissingRequestFieldsError
	}
	id, err := sqlconnect.CreatePostDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.CreatePostResponse{PostId: id}, nil
}

func (s *Server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := sqlconnect.GetPostDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}

	return &pb.GetPostResponse{Post: post}, nil
}

func (s *Server) GetUserPosts(ctx context.Context, req *pb.GetUserPostsRequest) (*pb.GetUserPostsResponse, error) {
	posts, err := sqlconnect.GetUserPostsDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.GetUserPostsResponse{Posts: posts}, nil
}

func (s *Server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	err := sqlconnect.DeletePostDbHandler(ctx, req)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &emptypb.Empty{}, nil
}
