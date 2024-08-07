package controllers

import (
    "context"
    "net/http"
    "strconv"
    "task_manager/data"
    "task_manager/models"
    "task_manager/middleware"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "log"
)

// TaskController handles task-related requests
type TaskController struct {
    service *data.TaskService
}

// NewTaskController creates a new TaskController
func NewTaskController(service *data.TaskService) *TaskController {
    return &TaskController{service: service}
}

// GetTasks retrieves all tasks
func (ctrl *TaskController) GetTasks(c *gin.Context) {
    tasks, err := ctrl.service.GetTasks(context.Background())
    if err != nil {
        log.Printf("Error getting tasks: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by ID
func (ctrl *TaskController) GetTaskByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    task, err := ctrl.service.GetTaskByID(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

// CreateTask creates a new task
func (ctrl *TaskController) CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }
    createdTask, err := ctrl.service.CreateTask(context.Background(), task)
    if err != nil {
        log.Printf("Error creating task: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask updates an existing task
func (ctrl *TaskController) UpdateTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }
    task, err := ctrl.service.UpdateTask(context.Background(), id, updatedTask)
    if err != nil {
        log.Printf("Error updating task: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID
func (ctrl *TaskController) DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    err = ctrl.service.DeleteTask(context.Background(), id)
    if err != nil {
        log.Printf("Error deleting task: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}

// Register handles user registration
func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    // Set the role to "user" by default
    if user.Role == "" {
        user.Role = "user"
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = string(hashedPassword)

    userService := c.MustGet("userService").(*data.UserService)
    if err := userService.CreateUser(&user); err != nil {
        log.Printf("Error creating user: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login handles user login
func Login(c *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    userService := c.MustGet("userService").(*data.UserService)
    user, err := userService.GetUserByUsername(loginData.Username)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if !userService.ValidatePassword(user, loginData.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := middleware.GenerateJWT(user.Username, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// PromoteAdmin promotes a user to an admin role
func (ctrl *TaskController) PromoteAdmin(c *gin.Context) {
    username := c.Param("username")

    userService := c.MustGet("userService").(*data.UserService)
    
    user, err := userService.GetUserByUsername(username)
    if err != nil {
        log.Printf("Error retrieving user: %v", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    log.Printf("User found: %+v", user) 

    err = userService.UpdateUserRole(context.Background(), username, "admin")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}
