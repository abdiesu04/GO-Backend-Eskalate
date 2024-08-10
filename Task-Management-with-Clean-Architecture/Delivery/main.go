package main

import (
	"context"
	"log"
	"task_manager/delivery/controllers"
	"task_manager/delivery/routers"
	"task_manager/repositories"
	"task_manager/usecases"
	"time"

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
	taskDatabase := client.Database("task_manager")
	userDatabase := client.Database("task_manager")

	taskRepo := repositories.NewTaskRepository(taskDatabase)
	userRepo := repositories.NewUserRepository(userDatabase)

	// Initialize use cases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	// Setup router and start server
	router := routers.SetupRouter(taskController, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
