package main

import (
	"context"
	"log"
	"net"

	"go-test-user/grpc_server"

	"google.golang.org/grpc"
)

type UserCenterServerImp struct {
	grpc_server.UnimplementedUserCenterServer
}

func (m *UserCenterServerImp) DoTest(context context.Context, req *grpc_server.DoTestReq) (*grpc_server.DoTestResp, error) {
	println(req.Text)
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9899")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	userCenterServer := &UserCenterServerImp{}

	grpc_server.RegisterUserCenterServer(s, userCenterServer)

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
