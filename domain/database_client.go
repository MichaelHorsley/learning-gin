package domain

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	// Set client options
	fullUri := "mongodb+srv://<username>:<password>@<server>/?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(fullUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func SaveNewTodo(todo Todo) (id interface{}, err error) {
	collection := mongoClient.Database("daily_planning").Collection("todos")

	insertResult, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatal(err)
	}

	x := insertResult.InsertedID

	fmt.Println("Inserted a single document: ", x)

	return x, err
}
