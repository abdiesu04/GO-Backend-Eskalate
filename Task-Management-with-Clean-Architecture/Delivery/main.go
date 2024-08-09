package main

import (
	"context"
	"log"
	"time"
	"task_manager/delivery/controllers"
	"task_manager/delivery/routers"
	"task_manager/repositories"
	"task_manager/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Context with timeout for MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize database and repositories
	database := client.Database("task_manager")
	taskRepo := repositories.NewTaskRepository(database)

	// Initialize use cases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUsecase)

	// Setup router and start server
	router := routers.SetupRouter(taskController)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
