package main

import (
	log "error_handing/client/log"
	pb "error_handing/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	status "google.golang.org/grpc/status"
)

const address = "127.0.0.1:8890"

func main() {
	// connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// init client
	c := pb.NewErrorHandingClient(conn)

	// use grpc client
	req := &pb.GetUserInfoReq{UserId: "123"}
	res, err := c.GetUserInfo(context.Background(), req)
	if err != nil {
		// write log
		log.Error.Printf("%+v\n", err)

		// todo
		estu := status.Code(err)
		if estu == codes.FailedPrecondition {
			// xxxx
		}
		return
	}
	// use rsp
	log.Info.Println("user name : ", res.UserName)
}
