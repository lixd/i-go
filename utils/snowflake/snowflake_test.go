package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
	"i-go/utils/ip"
)

// SnowFlake SonyFlake Mist 3种算法，生成全局自增Id相关测试.
/*
SnowFlake 245.1 ns/op
SonyFlake 607551 ns/op
Mist 520.5 ns/op
*/

func TestSnowFlake(t *testing.T) {
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

// BenchmarkSnowFlake 245.1 ns/op
func BenchmarkSnowFlake(b *testing.B) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Generate a snowflake Id.
		_ = node.Generate()
	}
}

func TestSonyFlake(t *testing.T) {
	st := sonyflake.Settings{
		StartTime:      time.Now(),
		MachineID:      GetNodeId,
		CheckMachineID: func(u uint16) bool { return u > 0 },
	}

	sf := sonyflake.NewSonyflake(st)
	for i := 0; i < 10; i++ {
		id, err := sf.NextID()
		if err != nil {
			log.Println("err:", err)
		}
		fmt.Println(id)
	}
}

// BenchmarkSonyFlake 607551 ns/op
func BenchmarkSonyFlake(b *testing.B) {
	st := sonyflake.Settings{
		StartTime:      time.Now(),
		MachineID:      GetNodeId,
		CheckMachineID: func(u uint16) bool { return u > 0 },
	}
	sf := sonyflake.NewSonyflake(st)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			_, _ = sf.NextID()
		}
	}
}

// GetNodeId retrieves the private IP address of the ecs instance
// and returns its lower 16 bits.
// It works correctly on Docker as well.
func GetNodeId() (uint16, error) {
	inIP, err := ip.IntranetIP()
	if err != nil {
		return 0, err
	}
	return uint16(inIP[2])<<8 + uint16(inIP[3]), nil
}

func TestGetNodeId(t *testing.T) {
	fmt.Println(GetNodeId())
}

func TestMist(t *testing.T) {
	// 使用方法
	mist := NewMist()
	for i := 0; i < 10; i++ {
		fmt.Println(mist.Generate())
	}
}

// BenchmarkMist 520.5 ns/op
func BenchmarkMist(b *testing.B) {
	mist := NewMist()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mist.Generate()
	}
}
