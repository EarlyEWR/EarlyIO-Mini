package main

import (
	context "context"
	fmt "fmt"
	log "log"
	net "net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/codeconvert"
)

type server struct {
	pb.UnimplementedCodeConverterServer
}

// GoToTypeScript implements the conversion logic (placeholder for now)
func (s *server) GoToTypeScript(ctx context.Context, req *pb.CodeRequest) (*pb.CodeReply, error) {
	goCode := req.GoCode
	// Placeholder: just wrap the Go code in a TypeScript comment
	tsCode := fmt.Sprintf("// Converted from Go:\n/*\n%s\n*/\n// TODO: Implement real conversion!", goCode)
	return &pb.CodeReply{TypescriptCode: tsCode}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCodeConverterServer(s, &server{})
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
