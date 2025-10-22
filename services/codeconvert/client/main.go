package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
)

// After proto generation, import generated package and perform RPC calls.

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to CodeConvert server (stubs not generated yet)")
	// Example after generation:
	// client := codeconvert.NewCodeConvertClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	// resp, _ := client.ToUpper(ctx, &codeconvert.ConvertRequest{Text: "hello"})
	// log.Printf("Upper: %s", resp.Result)
	time.Sleep(1 * time.Second)
}
