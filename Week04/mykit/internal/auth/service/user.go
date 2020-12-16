package service

import (
	"context"
	pb "mykit/api/auth/grpc"
)

func (s Service) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	resp := new(pb.LoginResp)
	resp.Token = "TokenToken"
	return resp, nil
}
