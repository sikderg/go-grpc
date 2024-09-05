# go-grpc
golang grpc protobuf example

## Client-Server Communication

This example demonstrates a simple client-server communication using gRPC in Go. The server implements a service defined in a `.proto` file, and the client makes requests to this service.

### Server

The server listens for incoming gRPC requests and responds according to the service definition.

### Client

The client sends requests to the gRPC server and processes the responses.

### Running the Example

1. **Generate gRPC code from the .proto file:**
   ```sh
   protoc --go_out=. --go-grpc_out=. path/to/your/service.proto
   ```

go run server/main.go

go run client/main.go