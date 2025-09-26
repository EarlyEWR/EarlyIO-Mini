package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/codeconvert"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCodeConverterClient(conn)

	goCode := `package main\nfunc main() {\n    println(\"Hello, world!\")\n}`
	if len(os.Args) > 1 {
		goCode = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GoToTypeScript(ctx, &pb.CodeRequest{GoCode: goCode})
	if err != nil {
		log.Fatalf("could not convert: %v", err)
	}
	log.Printf("TypeScript code:\n%s", r.TypescriptCode)
}
