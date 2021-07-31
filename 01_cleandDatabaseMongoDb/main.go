package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	type Checkpoint struct {
		Id      string    `bson:"_id,omitempty"`
		Data    time.Time `bson:"data"`
		Tag     string    `bson:"tag"`
		TagCode string    `bson:"tagCode,omitempty"`
	}

	/*
	   Connect to my cluster
	*/
	var mongoUrl string = "mongodb://username:password@localhost:27017/databasename?authSource=admin"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	/*
	   Get my collection instance
	*/
	collection := client.Database("unica").Collection("checkpoints")
	fmt.Println(collection)

	// cur, currErr := collection.Find(ctx, bson.D{})
	cur, currErr := collection.Find(nil, bson.M{"$and": []bson.M{{"tagCode": nil}, {"tag": nil}, {"data": nil}}})
	if currErr != nil {
		panic(currErr)
	}
	fmt.Println(cur)

	var checkpoints []Checkpoint
	if err = cur.All(ctx, &checkpoints); err != nil {
		panic(err)
	}
	for i, s := range checkpoints {
		fmt.Println(i, s)
	}

	fmt.Println(len(checkpoints))

	cursor, cursorErr := collection.DeleteMany(nil, bson.M{"$and": []bson.M{{"tagCode": nil}, {"tag": nil}, {"data": nil}}})

	if cursorErr != nil {
		panic(currErr)
	}
	fmt.Println(cursor)

}
