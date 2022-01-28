gen-gw:
	protoc --proto_path=./google \
     --go_out=/proto --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
      .hello_world.proto
