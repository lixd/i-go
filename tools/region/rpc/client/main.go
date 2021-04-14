package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "i-go/tools/region/rpc/proto"
	"i-go/utils/ip"
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
	for {
		c := pb.NewRegionServerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		externalIP, err := ip.ExternalIP()
		if err != nil {
			intranetIP, _ := ip.IntranetIP()
			externalIP = intranetIP
		}
		r, err := c.IP2Region(ctx, &pb.IP{
			Ip: externalIP,
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Region: %s", r.GetRegion())
		cancel()
		time.Sleep(time.Second * 3)
	}
}
