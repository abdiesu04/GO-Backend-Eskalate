package controllers

import (
	"net/http"
	"task_manager/domain"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: userUsecase,
	}
}

// Register handles user registration
func (c *UserController) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.UserUsecase.Register(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Login handles user login
func (c *UserController) Login(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.UserUsecase.Login(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// PromoteToAdmin handles promoting a user to admin
func (c *UserController) PromoteAdmin(ctx *gin.Context) {
	username := ctx.Param("username")
	err := c.UserUsecase.PromoteAdmin(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}
