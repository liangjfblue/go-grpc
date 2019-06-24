package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
	go_protoc "myselft/protobuf-grpc/demo2-deadline/protocal"
	"sync"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../conf/client/client.pem", "../conf/client/client.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		panic(err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		panic(err)
	}

	tCred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "go_protoc.HelloServer",
		RootCAs:      certPool,
	})

	clientConn, err := grpc.Dial("127.0.0.1:8099", grpc.WithTransportCredentials(tCred))
	if err != nil {
		panic(err)
	}
	defer clientConn.Close()


	client := go_protoc.NewHelloServerClient(clientConn)

	req := go_protoc.String{Value:"liangjf"}

	ctx, cancelFunc := context.WithDeadline(context.TODO(), time.Now().Add(time.Duration(time.Second*5)))

	var wg sync.WaitGroup
	wg.Add(1)
	go func(client go_protoc.HelloServerClient, ctx context.Context, req go_protoc.String) {
		defer wg.Done()
		resp, err := client.Hello(ctx, &req)
		if err != nil && status.Code(err) == codes.DeadlineExceeded {
			fmt.Println("grpc call timeout")
		} else if err != nil {
			panic(err)
		}

		fmt.Println(resp.GetValue())
	}(client, ctx, req)

	select {
	case <-ctx.Done():
		fmt.Println("cancel")
		cancelFunc()
	default:
		wg.Wait()
	}
}
