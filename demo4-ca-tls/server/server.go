package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
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
	cert, err := tls.LoadX509KeyPair("../conf/server/server.pem", "../conf/server/server.key")
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

	tCret := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	grpcServer := grpc.NewServer(grpc.Creds(tCret))
	go_protoc.RegisterHelloServerServer(grpcServer, new(HelloServerImpl))

	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
