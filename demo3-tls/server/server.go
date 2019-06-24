package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	go_protoc "myselft/protobuf-grpc/demo2-deadline/protocal"
	"net"
)

type HelloServerImpl struct {

}

func (h *HelloServerImpl)Hello(ctx context.Context, args *go_protoc.String) (*go_protoc.String, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.New(codes.Canceled, "Client cancelled, abandoning.").Err()
	}

	reply := &go_protoc.String{Value:"hello "+args.GetValue()}
	return reply, nil
}

func (h *HelloServerImpl)Hi(ctx context.Context, args *go_protoc.String) (*go_protoc.String, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.New(codes.Canceled, "Client cancelled, abandoning.").Err()
	}

	reply := &go_protoc.String{Value:"hi "+args.GetValue()}
	return reply, nil
}

func main() {
	tReadentials, err := credentials.NewServerTLSFromFile("../conf/server.pem", "../conf/server.key")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tReadentials))
	go_protoc.RegisterHelloServerServer(grpcServer, new(HelloServerImpl))

	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
