package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	go_protoc "myselft/protobuf-grpc/demo6-stream/protocal"
	"net"
	"strconv"
	"time"
)

type HelloServerImpl struct {

}

func (h *HelloServerImpl)GetResult1(request *go_protoc.Request, resp go_protoc.HelloServer_GetResult1Server) error {
	fmt.Println(request.Request)

	for i := 0; i < 10; i++ {
		_ = resp.Send(&go_protoc.Result{Reply: "Hello " + ":" + request.Request + " " + strconv.Itoa(i)})
		time.Sleep(time.Millisecond*500)
	}
	return nil
}

func (h *HelloServerImpl)GetResult2(stream go_protoc.HelloServer_GetResult2Server) error {
	for {
		result, err := stream.Recv()
		if err != nil && err != io.EOF{
			panic(err)
		} else if err != nil && err == io.EOF{
			break
		} else if err == nil && result != nil {
			fmt.Println(result.Request)
		}
	}

	reply := &go_protoc.Result{Reply:"after client stream, now server send"}
	if err := stream.SendAndClose(reply); err != nil {
		panic(err)
	}
	fmt.Println("after client stream, now server send")

	return nil
}


func (h *HelloServerImpl)GetResult3(stream go_protoc.HelloServer_GetResult3Server) error {
	for {
		result, err := stream.Recv()
		if err == io.EOF || codes.Canceled == status.Code(err){
			break
		} else if err != nil {
			panic(err)
		} else if result != nil {
			fmt.Println("GetResult3  recv  ", result.Request)

			reply := go_protoc.Result{Reply:"GetResult3 " + result.Request}
			fmt.Println(reply.Reply)
			if err := stream.Send(&reply); err != nil {
				panic(err)
			}
		}
	}
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	go_protoc.RegisterHelloServerServer(grpcServer, new(HelloServerImpl))

	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
