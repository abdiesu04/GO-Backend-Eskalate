package router

import (
    "task_manager/controllers"
    "task_manager/data"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(taskCollection *mongo.Collection) *gin.Engine {
    r := gin.Default()

    taskService := data.NewTaskService(taskCollection)
    taskController := controllers.NewTaskController(taskService)
  
    
    r.GET("/tasks", taskController.GetTasks)
    r.GET("/tasks/:id", taskController.GetTaskByID)
    r.POST("/tasks", taskController.CreateTask)
    r.PUT("/tasks/:id", taskController.UpdateTask)
    r.DELETE("/tasks/:id", taskController.DeleteTask)
    

    return r
}
