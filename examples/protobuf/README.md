# Protobuf generation

- Need to install the Go protocol buffer plugin:

`go get -u github.com/golang/protobuf/protoc-gen-go`

- To generate the code:

`protoc -I=. --go_out=. ./user.proto`
