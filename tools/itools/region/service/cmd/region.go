package main

import (
	"flag"
	"time"

	"github.com/micro/go-plugins/registry/etcdv3"
	"i-go/core/conf"
	"i-go/core/etcd"
	"i-go/tools/itools/region"
	pb "i-go/tools/itools/region/proto"
	"i-go/tools/itools/region/service/logic"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "conf/config_region.yaml", "the config file path")
	flag.Parse()

	if err := conf.Load(file); err != nil {
		panic(err)
	}

	logic.Init()
	Registry()
}

func Registry() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcd.Endpoints()
	})
	srv := grpc.NewService(
		micro.Name(region.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(reg),
	)
	srv.Init()

	_ = pb.RegisterRegionHandler(srv.Server(), new(logic.Region))

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
