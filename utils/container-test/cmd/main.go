package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func main() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer redisC.Terminate(ctx)
	endpoint, err := redisC.Endpoint(ctx, "")
	if err != nil {
		log.Println(err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	client.Set("hello", "world", time.Second*10)
	val := client.Get("hello").Val()
	fmt.Println(val)
}
