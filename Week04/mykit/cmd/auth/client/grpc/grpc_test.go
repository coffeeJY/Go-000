package main

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "mykit/api/auth/grpc"
	"testing"
)

const address = "127.0.0.1:8890"

func TestRepository_Login(t *testing.T) {
	convey.Convey("TestRepository_Login", t, func(c convey.C) {
		// connection
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			grpclog.Fatalln(err)
		}
		defer conn.Close()

		// init client
		client := pb.NewAuthServerClient(conn)

		// use grpc client
		req := &pb.LoginReq{UserName: "123", Password: "123456"}
		res, err := client.Login(context.Background(), req)
		if err != nil {
			// write log
			fmt.Printf("%+v\n", err)
			// todo
			return
		}
		// use rsp
		fmt.Println("user Token : ", res.Token)
		c.So(err, convey.ShouldBeNil)
		t.Log(res)
	})
}
