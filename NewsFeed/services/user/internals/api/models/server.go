package models

import pb "proto/user/gen"

type Server struct {
	pb.UnimplementedUserServiceServer
}
