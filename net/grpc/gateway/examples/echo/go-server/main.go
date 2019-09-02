// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	pb "github.com/grpc/grpc-web/net/grpc/gateway/examples/echo"
	//"golang.org/x/net/http2"
	//"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	address = ":9090"
)

type server struct{}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	/*return &pb.EchoResponse{
		Message:      req.Message,
		MessageCount: 1,
	}, nil*/
	return nil, status.Errorf(codes.Aborted, "")
}

func (s *server) EchoAbort(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return nil, status.Errorf(codes.Aborted, "")
}

func (s *server) NoOp(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) ServerStreamingEcho(req *pb.ServerStreamingEchoRequest, _ pb.EchoService_ServerStreamingEchoServer) error {
	return nil
}

func (s *server) ServerStreamingEchoAbort(req *pb.ServerStreamingEchoRequest, _ pb.EchoService_ServerStreamingEchoAbortServer) error {
	return nil
}

func (s *server) ClientStreamingEcho(_ pb.EchoService_ClientStreamingEchoServer) error {
	return nil
}

func (s *server) FullDuplexEcho(_ pb.EchoService_FullDuplexEchoServer) error {
	return nil
}

func (s *server) HalfDuplexEcho(_ pb.EchoService_HalfDuplexEchoServer) error {
	return nil
}

func grpcHandler(grpcServer *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			println("oops!")
		}
	})
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterEchoServiceServer(grpcServer, &server{})
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	/*
		httpServer := &http.Server{
			Addr:           address,
			Handler:        h2c.NewHandler(grpcHandler(grpcServer), &http2.Server{}),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(httpServer.ListenAndServe())
	*/
}
