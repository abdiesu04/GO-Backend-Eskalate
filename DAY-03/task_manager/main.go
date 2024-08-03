package main

import (
    "context"
    "fmt"
    "log"
    "task_manager/router"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Define MongoDB connection options
    opts := options.Client().ApplyURI("mongodb://localhost:27017")

    var client *mongo.Client
    var err error

    // Implement a retry mechanism for connecting to MongoDB
    for i := 0; i < 5; i++ {
        client, err = mongo.Connect(context.TODO(), opts)
        if err == nil {
            // Ping the MongoDB server to ensure a successful connection
            err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
            if err == nil {
                fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
                break
            }
        }
        log.Printf("Failed to connect to MongoDB (attempt %d/5): %v", i+1, err)
        time.Sleep(2 * time.Second)
    }

    if err != nil {
        log.Fatalf("Failed to connect to MongoDB after 5 attempts: %v", err)
    }

    defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            log.Fatal(err)
        }
    }()

    // Get the task collection
    taskCollection := client.Database("taskdb").Collection("tasks")

    // Set up and run the router
    r := router.SetupRouter(taskCollection)
    r.Run(":8080")
}