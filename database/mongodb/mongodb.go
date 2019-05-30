package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// func main() {
// 	fmt.Println("------------Main Begin----------")
//
// 	client, e := ConnMongo()
// 	if e != nil {
// 		fmt.Printf("mongo.Conn error=%v \n", e)
// 	}
// 	//选择数据库和集合
// 	collection := client.Database("db2").Collection("mycol2")
// 	// 测试 增删改查
// 	//Insert(collection)
// 	//Find(collection)
// 	//Delete(collection)
// 	//Update(collection)
// 	//Others(collection)
//
// 	AggregateTest(client)
// 	fmt.Println("------------Main End----------")
// }

func AggregateTest(client *mongo.Client) {
	collection := client.Database("db2").Collection("mycol3")
	data := []interface{}{}
	data = append(data, bson.M{"title": "first", "score": 80, "type": "Java"})
	data = append(data, bson.M{"title": "second", "score": 70, "type": "Java"})
	data = append(data, bson.M{"title": "third", "score": 60, "type": "Golang"})
	data = append(data, bson.M{"title": "fourth", "score": 50, "type": "Golang"})
	data = append(data, bson.M{"title": "fifth", "score": 40, "type": "C"})
	data = append(data, bson.M{"title": "sixth", "score": 30, "type": "C"})
	tags := make([]string, 3)
	tags = append(tags, "good")
	tags = append(tags, "fun")
	tags = append(tags, "computer")
	data = append(data, bson.M{"title": "seventh", "score": 20, "type": "C", "tags": tags})
	result, _ := collection.InsertMany(context.Background(), data)
	fmt.Printf("InsertMany result= %v \n", result)
	// Aggregate()
	//db.getCollection('mycol3').aggregate([{$group:{_id:"$title",total:{$sum:"$score"}}}])
	//db.getCollection('mycol3').aggregate([{$group:{_id:"$title",total:{$sum:"$score"}}},{$match:{total:{$gt:110,$lt:150}}}])
	//db.getCollection('mycol3').aggregate(
	//[{
	//     $group:
	//        {_id:"$type",
	//        totalScore:{$sum:"$score"}}},
	//    {$match:
	//        {totalScore:{
	//            $gt:10,$lt:160}}},
	//    {$sort:
	//        {totalScore:1}}
	//])
}

func Others(collection *mongo.Collection) {
	//查询集合里面有多少数据
	fmt.Println("------------CountDocuments----------")
	CountDocuments, e := collection.CountDocuments(context.Background(), nil)
	if e != nil {
		fmt.Printf("collection.CountDocuments error=%v \n", e)
	}
	fmt.Printf("collection.CountDocuments  result=%v \n", CountDocuments)
	//查询集合里面有多少数据(age>=20的数据)
	fmt.Println("------------CountDocuments Filter----------")
	CountDocuments2, e := collection.CountDocuments(context.Background(), bson.M{"age": bson.M{"$gte": 20}})
	if e != nil {
		fmt.Printf("collection.CountDocuments error=%v \n", e)
	}
	fmt.Printf("collection.CountDocuments  result=%v \n", CountDocuments2)
	fmt.Println()
}

func Update(collection *mongo.Collection) {
	//更新一条数据
	fmt.Println("------------UpdateOne----------")
	UpdateOne, e := collection.UpdateOne(context.Background(), bson.M{"name": "illusory"}, bson.M{"$set": bson.M{"name": "new Name"}})
	if e != nil {
		fmt.Printf("collection.UpdateOne error=%v \n", e)
	}
	fmt.Printf("collection.UpdateOne  result=%v \n", UpdateOne)
	//批量更新
	fmt.Println("------------UpdateMany----------")
	UpdateMany, e := collection.UpdateMany(context.Background(), bson.M{"age": bson.M{"$gte": 20}}, bson.M{"$set": bson.M{"age": 33}})
	if e != nil {
		fmt.Printf("collection.UpdateMany error=%v \n", e)
	}
	fmt.Printf("collection.UpdateMany  result=%v \n", UpdateMany)
	fmt.Println()
}

func Delete(collection *mongo.Collection) {
	// 删除单条数据
	fmt.Println("------------DeleteOne----------")
	DeleteOne, e := collection.DeleteOne(context.Background(), bson.M{"name": "illusory"})
	if e != nil {
		fmt.Printf("collection.DeleteOne error=%v \n", e)
	}
	fmt.Printf("collection.DeleteOne  result=%v \n", DeleteOne)
	//删除多条数据
	fmt.Println("------------DeleteMany----------")
	deleteResult, e := collection.DeleteMany(context.Background(), bson.M{"name": "illusory"})
	if e != nil {
		fmt.Printf("collection.DeleteMany error=%v \n", e)
	}
	fmt.Printf("collection.DeleteMany  result=%v \n", deleteResult)
	fmt.Println()
}

type User struct {
	Id primitive.ObjectID `bson:"_id" form:"id"`
	//Name string             `bson:"name"`
	Name string `bson:"name" form:"name"`
}

func Find(collection *mongo.Collection) {
	user1 := bson.M{}
	user2 := User{}
	//查询单条数据
	fmt.Println("------------FindOne----------")
	e := collection.FindOne(context.Background(), bson.M{"name": "illusory"}).Decode(&user1)
	e = collection.FindOne(context.Background(), bson.M{"name": "illusory"}).Decode(&user2)
	if e != nil {
		fmt.Printf("collection.FindOne error=%v \n", e)
	}
	fmt.Printf("collection.FindOne to bson result=%v \n", user1)
	fmt.Printf("collection.FindOne to user result=%v \n", user2)
	//批量查询
	fmt.Println("------------FindMany----------")
	cursor, e := collection.Find(context.Background(), bson.M{})
	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			fmt.Printf("collection.Find  cursor.Decode error=%v \n", e)
		}
		fmt.Printf("collection.Find  cursor.Decode result=%v \n", result)
	}
	e = cursor.Err()
	if e != nil {
		fmt.Printf("collection.Find error=%v \n", e)
	}
	// 查询age<20的 限制查询10条 按照age 降序排列
	fmt.Println("------------Find Sort----------")
	find, e := collection.Find(context.Background(), bson.M{"age": bson.M{"$lt": 20}}, options.Find().SetLimit(10), options.Find().SetSort(bson.M{"age": -1}))
	if e != nil {
		fmt.Printf("collection.Find   error=%v \n", e)
	}
	for find.Next(context.Background()) {
		var result bson.M
		e := find.Decode(&result)
		if e != nil {
			fmt.Printf("collection.Find sort  cursor.Decode error=%v \n", e)
		}
		fmt.Printf("collection.find sort result=%v \n", find)
	}
	fmt.Println()
}

func Insert(collection *mongo.Collection) {
	// 插入单条数据
	fmt.Println("------------InsertOne----------")
	InsertOne, e := collection.InsertOne(context.Background(), bson.M{"name": "illusory"}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory", "age": 20}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory2", "age": 21}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory3", "age": 22}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory4", "age": 19}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory5", "age": 18}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory6", "age": 17}, options.InsertOne())
	_, e = collection.InsertOne(context.Background(), bson.M{"name": "illusory7", "age": 16}, options.InsertOne())
	if e != nil {
		fmt.Printf("collection.InsertOne error=%v \n", e)
	}
	fmt.Printf("collection.InsertOne result=%v \n", InsertOne)
	//插入多条数据
	fmt.Println("------------InsertMany----------")
	users := []interface{}{}
	users = append(users, bson.M{"name": "illusory"})
	users = append(users, bson.M{"age": 23})
	users = append(users, bson.M{"address": "CQ"})
	manyResult, e := collection.InsertMany(context.Background(), users)
	if e != nil {
		fmt.Printf("collection.InsertMany error=%v \n", e)
	}
	fmt.Printf("collection.InsertMany result=%v \n", manyResult)
}

func ConnMongo() (*mongo.Client, error) {
	// username:password@hostname/dbname
	// 获取 client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.138:27017").SetAuth(options.Credential{
		Username:   "admin",
		Password:   "123456",
		AuthSource: "admin"}))
	if err != nil {
		fmt.Printf("mongo.NewClient error=%v", err)
	}
	// 设置 30s 超时
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//连接
	errC := client.Connect(ctx)
	if errC != nil {
		fmt.Print(errC)
	}
	// 判断服务是否可用
	errP := client.Ping(ctx, readpref.Primary())
	if errP != nil {
		fmt.Print(errP)
	}
	return client, err
}

type MongoCollection interface {
	GetCollectionName() string
}

func GetCollection(c MongoCollection) *mongo.Collection {
	client, e := ConnMongo()
	if e != nil {
		return nil
	}
	database := client.Database("vaptcha")
	return database.Collection(c.GetCollectionName())
}
