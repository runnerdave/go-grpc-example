# gRPC example in Go

## installation
install protocol buffers: https://github.com/protocolbuffers/protobuf#protocol-compiler-installation
````
$ brew install protobuf
````
This will install all main languages support, except Go, so follow these instructions for Go support https://github.com/golang/protobuf#installation
````
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u -v github.com/golang/protobuf/proto
````
Then make sure your $GOBIN is in your path:
````
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"

export PATH="$GOBIN:$PATH"
````
## generate the protocol buffer code
(this step is optional since the generated code is shipped along in ths repository)
```
$./generate.sh
```

## running the gRPC Server
### Server
```
$ go run calculator_server/server.go
```
### Client
```
$ go run calculator_client/client.go
```
