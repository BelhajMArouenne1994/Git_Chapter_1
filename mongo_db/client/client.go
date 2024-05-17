package client

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, context.Context, context.CancelFunc) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://belhajmarwen5:OhRno4rqFq14J7bE@sfmcmarouenne.q1axxvs.mongodb.net/?tls=true&retryWrites=true&w=majority&appName=sfmcMarouenne").SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, ctx, cancel
}
