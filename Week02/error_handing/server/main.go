package main

import (
	pb "error_handing/pb"
	dao "error_handing/server/dao"
	myerr "error_handing/server/err"
	log "error_handing/server/log"
	"errors"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"net"
)

const address = "127.0.0.1:8890"

type grpcService struct{}

func (s grpcService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRsp, error) {
	resp := new(pb.GetUserInfoRsp)

	user, err := dao.QueryUserByID(req.UserId)
	if err != nil {
		// write log
		log.Error.Printf("%+v\n", err)
		// assertion error type, and only return root error
		if errors.Is(err, myerr.ErrNoRowsSql) {
			err = status.Error(codes.FailedPrecondition, err.Error())
		}
		return nil, err
	}

	// todo: use info
	resp.UserId = user.ID
	resp.UserName = user.Name
	resp.Age = int32(user.Age)
	return resp, nil
}

var GrpcService = grpcService{}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Error.Printf("Failed to listen: %v", err)
		return
	}

	// new grpc Server
	s := grpc.NewServer()

	// Register Service
	pb.RegisterErrorHandingServer(s, GrpcService)

	log.Info.Println("Listen on " + address)
	err = s.Serve(listen)
	if err != nil {
		log.Error.Fatalf("grpc Serve start err : %v", err)
		return
	}
}
