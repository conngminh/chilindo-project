package rpc_client

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"practice/pkg/pb/admin"
)

const adminClientPort = ":50051"

type IRPCClient interface {
	SetUpAdminClient() admin.AdminServiceClient
}

type RPCClient struct{}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}

func (R *RPCClient) SetUpAdminClient() admin.AdminServiceClient {

	addr := ":50051"

	conn, dialErr := grpc.Dial(addr, grpc.WithInsecure())
	if dialErr != nil {
		log.Fatalf("failed to connect: %v", dialErr)
	}

	adminClient := admin.NewAdminServiceClient(conn)
	fmt.Println("Listen to admin service on port: ", adminClientPort)
	return adminClient
}
