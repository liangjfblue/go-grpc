package main

import (
	"myselft/protobuf-grpc/demo1-rpc/protocal"
	"net"
	"net/rpc"
)

type HelloServer struct {

}

func (h *HelloServer)Hello(req *go_protoc.String, resp *go_protoc.String) error {
	resp.Value = "hello " + req.GetValue()
	return nil
}

func main() {
	err := rpc.RegisterName("HelloServer", new(HelloServer))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	rpc.ServeConn(conn)
}
