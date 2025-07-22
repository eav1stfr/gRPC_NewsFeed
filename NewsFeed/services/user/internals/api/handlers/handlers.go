package handlers

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "proto/user/gen"
	"strconv"
	"user/internals/api/sqlconnect"
	"utils"
)

func (s *Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := sqlconnect.RegisterUserDbHandler(ctx, request)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.RegisterResponse{UserId: strconv.Itoa(id)}, nil
}

//message LoginRequest {
//	string email = 1;
//	string password = 2;
//}
//
//message LoginResponse {
//	string jwt_token = 1;
//}

func (s *Server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	passwordHash, err := sqlconnect.GetPasswordHashByUsername(ctx, request.Email)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	err = utils.VerifyPassword(passwordHash, request.Password)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}

	return nil, nil
}

func (s *Server) GetProfile(ctx context.Context, request *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	user, err := sqlconnect.GetUserProfileDbHandler(ctx, request)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.GetProfileResponse{User: user}, nil
}

func (s *Server) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*emptypb.Empty, error) {
	err := sqlconnect.UpdateProfileDbHandler(ctx, request)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) Follow(ctx context.Context, request *pb.FollowRequest) (*emptypb.Empty, error) {
	err := sqlconnect.FollowDbHandler(ctx, request)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) Unfollow(ctx context.Context, request *pb.UnfollowRequest) (*emptypb.Empty, error) {
	err := sqlconnect.UnfollowDbHandler(ctx, request)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) GetFollowing(ctx context.Context, request *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	users, err := sqlconnect.GetFollowingListDbHandler(ctx, request, false)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.FollowListResponse{Users: users}, nil
}

func (s *Server) GetFollowers(ctx context.Context, request *pb.FollowListRequest) (*pb.FollowListResponse, error) {
	users, err := sqlconnect.GetFollowingListDbHandler(ctx, request, true)
	if err != nil {
		return nil, utils.ToGRPCStatus(err)
	}
	return &pb.FollowListResponse{Users: users}, nil
}
