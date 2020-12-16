package service

import (
	"context"
	pb "mykit/api/auth/grpc"
)

func (s Service) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	resp := new(pb.LoginResp)
	token, err := s.repo.Login(req.UserName, req.Password)
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return resp, nil
}
