gRPC is a high-performance Remote Procedure Call (RPC) framework created by Google.
It lets a program call a function that runs on another machine as if it were local.

Core ideas of gRPC:
1. Define services(API) with Protocol Buffers
2. Use HTTP/2 under the hood
3. Binary serialization
---
Why HTTP/2?
1. multiplexing (many requests on one connection)
2. streaming
3. lower latency

What is multiplexing in HTTP/2?
HTTP/2 delivers each request in an independent stream, and the server handles each stream concurrently.  
| Request    | Stream ID | Handler   |
| ---------- | --------- | --------- |
| GetUser    | 1         | A         |
| CreateUser | 3         | B         |
| DeleteUser | 5         | C         |
---
gRPC vs REST
|                  | gRPC              | REST           |
| ---------------- | ----------------- | -------------- |
| Data format      | Binary (protobuf) | Usually JSON   |
| Speed            | Very fast         | Slower         |
| Streaming        | Built-in          | Harder         |
| Browser friendly | Not directly      | Yes            |
| Schema           | Strict, defined   | Often flexible |
---
Why backend engineers love it? Especially in microservices:

✅ Strong type safety
✅ Auto-generated clients
✅ High performance
✅ Clear contracts
✅ Great for internal service-to-service calls
✅ Multi-language support
---
Install
1. install protoc compiler, reads .proto files.<br>
`sudo apt install protobuf-compiler`

2. install protobuf Go plugin, generates go code.<br>
`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`<br>
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

3. ensure `$GOPATH/bin` is on `$PATH`, protoc will find protobuf Go plugin in `$PATH`.<br>
`export PATH="$(go env GOPATH)/bin:$PATH"`

4. install gRPC and proto runtime
`go get google.golang.org/grpc`
`go get google.golang.org/protobuf/proto`
---
Generate go code
```
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       user.proto
```