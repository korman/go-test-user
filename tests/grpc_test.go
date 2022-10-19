package tests

import (
	"context"
	"testing"

	"go-test-user/grpc_server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestBaseGrpc(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9899", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Error(err)
	}

	userClient := grpc_server.NewUserCenterClient(conn)
	defer conn.Close()

	loginResp, err := userClient.DoTest(context.Background(), &grpc_server.DoTestReq{
		Text: "hello world",
	})

	if err != nil {
		t.Error(err)
	}

	println(loginResp.Result)
}
