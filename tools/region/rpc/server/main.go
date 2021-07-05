package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"i-go/core/conf"
	"i-go/tools/region/core"
	pb "i-go/tools/region/rpc/proto"
)

const (
	port = ":50051"
)

// greeterServer 定义一个结构体用于实现 .proto文件中定义的方法
// 新版本 gRPC 要求必须嵌入 pb.UnimplementedGreeterServer 结构体
type regionServer struct {
	pb.UnimplementedRegionServerServer
}

// IP2Region RPC 实现
func (r *regionServer) IP2Region(ctx context.Context, in *pb.IP) (*pb.Region, error) {
	region := core.IP2Region(in.GetIp())
	log.Printf("IP2Region ip:%s region:%s\n", in.GetIp(), region)
	return &pb.Region{Region: region}, nil
}

// IP2LatLong RPC 实现
func (r *regionServer) IP2LatLong(ctx context.Context, in *pb.IP) (*pb.LatLong, error) {
	latLong, err := core.IP2LatLong(in.GetIp())
	if err != nil {
		return nil, err
	}
	item := &pb.LatLong{
		Latitude:  latLong.Latitude,
		Longitude: latLong.Longitude,
	}
	log.Printf("IP2LatLong ip:%s lagLong:%+v\n", in.GetIp(), item)
	return item, nil
}

func main() {
	err := conf.Load("conf/config_ip.yaml")
	if err != nil {
		panic(err)
	}
	core.InitRegion()
	core.InitLatLong()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 将服务描述(server)及其具体实现(greeterServer)注册到 gRPC 中去.
	// 内部使用的是一个 map 结构存储，类似 HTTP server。
	pb.RegisterRegionServerServer(s, &regionServer{})
	log.Println("Serving gRPC on 0.0.0.0" + port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
