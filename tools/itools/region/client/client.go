package region

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"i-go/core/etcd"
	"i-go/tools/itools/region"
	"i-go/tools/itools/region/proto"
)

var regionService proto.RegionService

func Init() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etcd.Endpoints()
	})
	srv := grpc.NewService(
		micro.Registry(reg),
	)
	srv.Init()

	regionService = proto.NewRegionService(region.ServiceName, srv.Client())
}

func Ip2Region(ip string) string {
	in := &proto.RegionReq{
		Ip: ip,
	}

	resp, err := regionService.Ip2Region(context.Background(), in)
	if err != nil || resp.Region == "" {
		return "保留地址"
	}
	return resp.Region
}
