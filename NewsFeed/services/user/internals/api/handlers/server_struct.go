package handlers

import pb "proto/user/gen"

type Server struct {
	pb.UnimplementedUserServiceServer
}
