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
    taskService := data.NewTaskService(taskCollection)
    userService := &data.UserService{Collection: userCollection}

    // Register task and user services as middleware
    r.Use(func(c *gin.Context) {
        c.Set("taskService", taskService)
        c.Set("userService", userService)
        c.Next()
    })

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    authorized := r.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.POST("/tasks", controllers.NewTaskController(taskService).CreateTask)
        authorized.GET("/tasks", controllers.NewTaskController(taskService).GetTasks)
        authorized.GET("/tasks/:id", controllers.NewTaskController(taskService).GetTaskByID)
        authorized.PUT("/tasks/:id", controllers.NewTaskController(taskService).UpdateTask)
        authorized.DELETE("/tasks/:id", controllers.NewTaskController(taskService).DeleteTask)
    }

    return r
}
