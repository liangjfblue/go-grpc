package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("before handler: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Println("after handler: %s, %v", info.FullMethod, resp)
	return resp, err
}

func LoggingInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("LoggingInterceptor1 before handler: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Println("LoggingInterceptor1 after handler: %s, %v", info.FullMethod, resp)
	return resp, err
}

func LoggingInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("LoggingInterceptor2 before handler: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Println("LoggingInterceptor2 after handler: %s, %v", info.FullMethod, resp)
	return resp, err
}

func main() {
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor))

	//go-grpc-middleware 多拦截器
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				LoggingInterceptor1,
				LoggingInterceptor2)))
	go_protoc.RegisterHelloServerServer(grpcServer, new(HelloServerImpl))

	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
