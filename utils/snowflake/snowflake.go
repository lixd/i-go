package main

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

// 雪花算法
func main() {

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake Id.
	id := node.Generate()

	// Print out the Id in a few different ways.
	fmt.Printf("Int64  Id: %d\n", id)
	fmt.Printf("String Id: %s\n", id)
	fmt.Printf("Base2  Id: %s\n", id.Base2())
	fmt.Printf("Base64 Id: %s\n", id.Base64())

	// Print out the Id's timestamp
	fmt.Printf("Id Time  : %d\n", id.Time())

	// Print out the Id's node number
	fmt.Printf("Id Node  : %d\n", id.Node())

	// Print out the Id's sequence number
	fmt.Printf("Id Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("Id       : %d\n", node.Generate().Int64())
}
