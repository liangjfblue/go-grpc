package main

import (
	"fmt"
	go_protoc "myselft/protobuf-grpc/demo1-rpc/protocal"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	resp := go_protoc.String{}
	req := go_protoc.String{
		Value:"liangjf",
	}

	if err = client.Call("HelloServer.Hello", &req, &resp); err != nil {
		panic(err)
	}

	fmt.Println(resp.Value)
}
