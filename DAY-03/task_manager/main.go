package main

import (
    "context"
    "log"
    "task_manager/router"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "github.com/joho/godotenv"
    "os"
    "time"
)


// main connects to the database, sets up the router, and starts the server.
func main() {
    // Set up the MongoDB client
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    mongoURL := os.Getenv("MONGO_URL")
    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
    if err != nil {
        log.Fatal(err)
    }

    // Create a context with a timeout of 10 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel() // Cancel the context when the function exits
    
    // Connect to the MongoDB server
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Ping the primary to check the connection
    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatal(err)
    }

    // Get collections from the database for task  and users 
    taskCollection := client.Database("task_manager_db").Collection("tasks")
    userCollection := client.Database("task_manager_db").Collection("users")

    // Set up the router
    r := router.SetupRouter(taskCollection, userCollection)

    r.Run(":8080")
}
