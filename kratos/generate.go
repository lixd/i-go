package kratos

// //go:generate protoc --proto_path=./features/proto \
//  --go_out=./features/proto --go_opt=paths=source_relative \
//    --go-grpc_out=./features/proto --go-grpc_opt=paths=source_relative \
//   --grpc-gateway_out=./features/proto --grpc-gateway_opt=paths=source_relative \
//   ./features/proto/gateway/gateway.proto

//go:generate protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --openapi_out=. helloworld/hello_world.proto

// protobuf 生成 openAPI https://mp.weixin.qq.com/s/WQ3HP02ePA9t1SpKd-kP7A
