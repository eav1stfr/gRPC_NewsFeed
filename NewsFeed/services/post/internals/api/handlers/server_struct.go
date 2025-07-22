package handlers

import (
	pb "proto/post/gen"
)

type Server struct {
	pb.UnimplementedPostServiceServer
}
