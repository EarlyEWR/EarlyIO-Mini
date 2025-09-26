# CodeConvert gRPC Example

This is a minimal gRPC system for converting Go code to TypeScript code.

## Files
- `codeconvert.proto`: gRPC service definition
- `server.go`: Go server implementation

## How to use
1. Generate Go code from the proto file:
   ```sh
   protoc --go_out=. --go-grpc_out=. codeconvert.proto
   ```
2. Run the server:
   ```sh
   go run server.go
   ```
3. Implement a client to call the `GoToTypeScript` RPC.

*Note: The conversion logic is a placeholder. Replace it with real code conversion logic as needed.*
