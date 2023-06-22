# Go-gRPC

```
protoc -I greet/proto --go_out=. --go_opt=module=github.com/suhelzende/go-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/suhelzende/go-grpc  .\greet\proto\greet.proto
```