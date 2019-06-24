package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	go_protoc "myselft/protobuf-grpc/demo6-stream/protocal"
	"strconv"
	"time"
)

//Test 服务端单向流
func TestGetResult1(client go_protoc.HelloServerClient, ctx context.Context, done chan bool) {
	req := go_protoc.Request{Request:"TestGetResult1"}
	stream, err := client.GetResult1(ctx, &req)
	if err != nil && status.Code(err) == codes.DeadlineExceeded {
		fmt.Println("grpc call timeout")
	} else if err != nil {
		panic(err)
	}
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println("TestGetResult1 stream recv from server: %s", result.Reply)
	}
	done <- true
	return
}

//Test 客户端单向流
func TestGetResult2(client go_protoc.HelloServerClient, ctx context.Context, done chan bool)  {
	stream, err := client.GetResult2(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("TestGetResult2 now client send request")
	for i := 0; i < 5; i++ {
		req := go_protoc.Request{Request:"liangjf  " + strconv.Itoa(i)}
		if err = stream.Send(&req); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond*500)
		fmt.Println("TestGetResult2 stream send to server: %s", req.Request)
	}

	result, err := stream.CloseAndRecv();
	if err != nil {
		panic(err)
	}

	fmt.Println("TestGetResult2  ", result.Reply)
	done <- true
	return
}

//Test 客户端服务端双向流
func TestGetResult3(client go_protoc.HelloServerClient, ctx context.Context, done chan bool) {
	stream, err := client.GetResult3(ctx)
	if err != nil && status.Code(err) == codes.DeadlineExceeded {
		fmt.Println("grpc call timeout")
	} else if err != nil {
		panic(err)
	}

	allDone := make(chan bool, 1)
	go TestGetResult3Send(stream, allDone)

	go TestGetResult3Recv(stream)

	select {
	case <-allDone:
		done<-true
		break
	}
	return
}

func TestGetResult3Send(stream go_protoc.HelloServer_GetResult3Client, allDone chan bool) {
	fmt.Println("TestGetResult3Send now client send request")
	for i := 0; i < 5; i++ {
		req := go_protoc.Request{Request:"liangjf  " + strconv.Itoa(i)}
		if err := stream.Send(&req); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond*500)
		fmt.Println("TestGetResult3Send stream send to server: %s", req.Request)
	}

	if err := stream.CloseSend(); err != nil {
		panic(err)
	}

	allDone<-true
	return
}

func TestGetResult3Recv(stream go_protoc.HelloServer_GetResult3Client) {
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println("TestGetResult3Recv stream recv from server: %s", result.Reply)
	}
	return
}

func main() {
	clientConn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer clientConn.Close()

	client := go_protoc.NewHelloServerClient(clientConn)

	//为了测试stream，deadline设置大点
	ctx, cancelFunc := context.WithDeadline(context.TODO(), time.Now().Add(time.Duration(time.Second*10)))

	//done1 := make(chan bool, 1)
	//go TestGetResult1(client, ctx, done1)
	//
	//done2 := make(chan bool, 1)
	//go TestGetResult2(client, ctx, done2)

	done3 := make(chan bool, 3)
	go TestGetResult3(client, ctx, done3)

	select {
	case <-ctx.Done():
		fmt.Println("cancel")
		cancelFunc()
	//case <-done1:
	//	break
	//case <-done2:
	//	break
	case <-done3:
		break
	}
}
