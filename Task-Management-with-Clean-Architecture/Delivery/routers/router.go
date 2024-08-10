package routers

import (
	"task_manager/delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	// Task-related routes
	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks", taskController.GetAllTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	// User-related routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	return router
}
