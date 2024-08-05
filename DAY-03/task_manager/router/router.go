package router

import (
    "task_manager/controllers"
    "task_manager/data"
    "task_manager/middleware"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
)


func SetupRouter(taskCollection *mongo.Collection, userCollection *mongo.Collection) *gin.Engine {
    r := gin.Default()

    // Create task and user services.
    taskService := data.NewTaskService(taskCollection)
    userService := &data.UserService{Collection: userCollection}

    // Register task and user services as middleware.
    r.Use(func(c *gin.Context) {
        c.Set("taskService", taskService)
        c.Set("userService", userService)
        c.Next()
    })

    // Register API endpoints for user registration and authentication.
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Group authorized endpoints.
    authorized := r.Group("/")

    // Require authentication for authorized endpoints.
    authorized.Use(middleware.AuthMiddleware())

    // Register API endpoints for task management.
    {
        authorized.POST("/tasks", controllers.NewTaskController(taskService).CreateTask)
        authorized.GET("/tasks", controllers.NewTaskController(taskService).GetTasks)
        authorized.GET("/tasks/:id", controllers.NewTaskController(taskService).GetTaskByID)
        authorized.PUT("/tasks/:id", controllers.NewTaskController(taskService).UpdateTask)
        authorized.DELETE("/tasks/:id", controllers.NewTaskController(taskService).DeleteTask)
    }

    // Return the configured router.
    return r
}
