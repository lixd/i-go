package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.WithValue(context.Background(), "k1", "v1")
	ctx2 := context.WithValue(ctx1, "k1", "v11")
	value := ctx2.Value("k1")
	fmt.Println(value)
}
