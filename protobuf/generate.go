package protobuf

//go:generate protoc --proto_path=. --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --go-http_out=paths=source_relative:. derssbook.proto
