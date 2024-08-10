package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	router := gin.New()

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	apiGroup := router.Group("/")
	apiGroup.Use(infrastructure.AuthMiddleware())

	// Routes that only require authentication
	apiGroup.GET("/tasks", taskController.GetAllTasks)
	apiGroup.GET("/tasks/:id", taskController.GetTaskByID)

	// Routes that require both authentication and admin role
	adminGroup := apiGroup.Group("/")
	adminGroup.Use(infrastructure.RoleMiddleware("admin"))
	adminGroup.POST("/tasks", taskController.CreateTask)
	adminGroup.PUT("/tasks/:id", taskController.UpdateTask)
	adminGroup.DELETE("/tasks/:id", taskController.DeleteTask)
	adminGroup.POST("/promote/:username", userController.PromoteAdmin)

	return router
}
