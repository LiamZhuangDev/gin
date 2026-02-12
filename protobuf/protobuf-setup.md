# Install
1. install protoc compiler, reads .proto files.
`sudo apt install protobuf-compiler`

2. install protobuf Go plugin, generates go code.
`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

3. ensure $GOPATH/bin is on $PATH, protoc will find protobuf Go plugin in $PATH.
`export PATH="$(go env GOPATH)/bin:$PATH"`

4. install protobuf runtime, required by generated .pd.go code.
`go get google.golang.org/protobuf/proto`

# Create a .proto file

# Compile and generate Go file, for example the following command will generate `user.pb.go` from the given `user.proto`. 
`protoc --go_out=. --go_opt=paths=source_relative user.proto`