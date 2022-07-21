package main

import (
	"context"
	"log"
	"time"

	pb "i-go/tools/region/rpc/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRegionServerClient(conn)

	in := &pb.IP{Ip: "183.69.225.99"}
	for i := 0; i < 1000; i++ {
		ip2region(c, in)
		ip2LatLong(c, in)
	}
}

func ip2region(c pb.RegionServerClient, in *pb.IP) {
	r, err := c.IP2Region(context.Background(), in)
	if err != nil {
		log.Fatalf("could not IP2Region: %v", err)
	}
	log.Printf("Region: %s", r.GetRegion())
}

func ip2LatLong(c pb.RegionServerClient, in *pb.IP) {
	long, err := c.IP2LatLong(context.Background(), in)
	if err != nil {
		log.Fatalf("could not IP2LatLong: %v", err)
	}
	log.Printf("lag: %v long:%v \n", long.GetLatitude(), long.GetLongitude())
}
