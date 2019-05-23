package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("main start")
	//连接
	client := ConnRedis("127.0.0.1:6379", "", 1)
	//存
	err := client.Set("name", "illusory", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("client=%v \n", client)
	//取
	get, err := client.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", get)
	}
	ConnMongo()
}
func ConnRedis(addr string, pwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db})
	return client
}
func ConnMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.138:27017"))
	if err != nil {
		fmt.Printf("mongo.NewClient error=%v", err)
		return
	}
	collection := client.Database("baz").Collection("qux")
	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		fmt.Printf("collection.InsertOne error=%v", err)
	}
	fmt.Println(res)

}

