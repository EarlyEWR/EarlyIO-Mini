package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

// After proto generation, import the generated package:
// import codeconvert "github.com/EarlyEWR/EarlyIO-Mini/protos/codeconvert"
// type server struct { codeconvert.UnimplementedCodeConvertServer }
// func (s *server) ToUpper(ctx context.Context, req *codeconvert.ConvertRequest) (*codeconvert.ConvertResponse, error) { ... }

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	grpcServer := grpc.NewServer()
	// codeconvert.RegisterCodeConvertServer(grpcServer, &server{}) // enable after generation
	log.Println("CodeConvert gRPC server listening on :50051 (awaiting generated proto stubs)")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
