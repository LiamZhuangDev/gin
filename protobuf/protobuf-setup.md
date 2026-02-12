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

# JSON vs Protobuf
| Aspect                 | JSON                             | Protobuf                      |
| ---------------------- | -------------------------------- | ----------------------------- |
| Human readability      | ✅ Easy for humans to read       | ❌ Binary, not human-readable |
| Payload size           | ❌ Larger (contains field names) | ✅ Smaller (uses numeric tags)|
| Serialization speed    | Normal                           | ✅ Faster                     |
| CPU usage              | Higher                           | ✅ Lower                      |
| Type safety            | ❌ Weak, type mistakes are common| ✅ Strongly typed             |
| API contract           | Loose / flexible                 | ✅ Strict schema              |
| Code generation        | ❌ Usually handwritten models    | ✅ Auto-generated             |
| Backward compatibility | Convention-based                 | ✅ Built-in rules             |
| Debugging              | ✅ Easy with curl / browser      | ❌ Needs decoding tools       |
| Learning curve         | ✅ Easy                          | ❌ Need proto + generation    |
| Dev workflow           | Change & run                     | Change proto → regenerate     |
| Best for               | Public APIs / frontend           | Internal services / RPC       |
| Network cost           | ❌ Higher                        | ✅ Lower                      |
| High-load performance  | OK                               | ✅ Better for high QPS        |

| Scenario                         | Better choice |
| -------------------------------- | ------------- |
| Browser / third-party API        | JSON          |
| Service-to-service communication | Protobuf      |
| Mobile / poor network            | Protobuf      |
| Maximum performance              | Protobuf      |
| Need easy debugging              | JSON          |

# Protobuf → Go Type Mapping
1. Simple types<br>
| Protobuf type | Go type (generated) |
| ------------- | ------------------- |
| double        | float64             |
| float         | float32             |
| int32         | int32               |
| int64         | int64               |
| uint32        | uint32              |
| uint64        | uint64              |
| sint32        | int32               |
| sint64        | int64               |
| fixed32       | uint32              |
| fixed64       | uint64              |
| sfixed32      | int32               |
| sfixed64      | int64               |
| bool          | bool                |
| string        | string              |
| bytes         | []byte              |

2. Complex types<br>
| proto       | Go          |
| ----------- | ----------- |
| repeated    | slice       |
| message     | pointer     |
| enum        | int32       |
| map         | map         |
