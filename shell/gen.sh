mkdir -p go/video
protoc -I=proto/video --go_out=go/video --go-grpc_out=go/video proto/video/*.proto
