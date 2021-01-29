package logic

import (
	"context"

	pb "i-go/tools/itools/region/proto"
)

type Region struct {
}

func (reg *Region) Ip2Region(_ context.Context, req *pb.RegionReq, resp *pb.RegionResp) error {
	resp.Region = reg.ip2Region(req.Ip)
	return nil
}
