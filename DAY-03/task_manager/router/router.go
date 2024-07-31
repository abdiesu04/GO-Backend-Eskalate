package router

import (
    "task_manager/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    taskController := controllers.NewTaskController()

    v1 := r.Group("api/v1")
    {
        v1.GET("/tasks", taskController.GetTasks)
        v1.GET("/tasks/:id", taskController.GetTaskByID)
        v1.PUT("/tasks/:id", taskController.UpdateTask)
        v1.DELETE("/tasks/:id", taskController.DeleteTask)
        v1.POST("/tasks", taskController.CreateTask)
    }

    return r
}