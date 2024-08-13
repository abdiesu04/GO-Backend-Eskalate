package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"task_manager/Delivery/controllers"
	"task_manager/domain"
	"task_manager/usecases/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	userController *controllers.UserController
	mockUsecase    *mocks.UserUsecase
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.UserUsecase)
	suite.userController = controllers.NewUserController(suite.mockUsecase)
}

func (suite *UserControllerTestSuite) TestRegister_Success() {
	user := &domain.User{
		Username: "abdiesu04",
		Password: "pass123123",
	}
	suite.mockUsecase.On("Register", mock.Anything, user).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/register", strings.NewReader(`{"username":"abdiesu04","password":"pass123123"}`))
	ctx.Request.Header.Set("Content-Type", "application/json")

	suite.userController.Register(ctx)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "User registered successfully")
	suite.mockUsecase.AssertCalled(suite.T(), "Register", mock.Anything, user)
}

func (suite *UserControllerTestSuite) TestRegister_BadRequest() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/register", strings.NewReader(`{"username":"testuser"`)) // Invalid JSON
	ctx.Request.Header.Set("Content-Type", "application/json")

	suite.userController.Register(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid request data")
	assert.Contains(suite.T(), w.Body.String(), "details")
}

func (suite *UserControllerTestSuite) TestRegister_InternalServerError() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}
	suite.mockUsecase.On("Register", mock.Anything, user).Return(errors.New("database error"))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/register", strings.NewReader(`{"username":"testuser","password":"testpassword"}`))
	ctx.Request.Header.Set("Content-Type", "application/json")

	suite.userController.Register(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Failed to register user")
	assert.Contains(suite.T(), w.Body.String(), "details")
}

func (suite *UserControllerTestSuite) TestLogin_Success() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}
	token := "mocked_token"
	suite.mockUsecase.On("Login", mock.Anything, user).Return(token, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser","password":"testpassword"}`))
	ctx.Request.Header.Set("Content-Type", "application/json")

	suite.userController.Login(ctx)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), token)
}

func (suite *UserControllerTestSuite) TestLogin_BadRequest() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser"`)) // Invalid JSON
	ctx.Request.Header.Set("Content-Type", "application/json")

	suite.userController.Login(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Invalid login data")
	assert.Contains(suite.T(), w.Body.String(), "details")
}

func (suite *UserControllerTestSuite) TestLogin_Unauthorized() {
	user := &domain.User{
		Username: "testuser",
		Password: "testpassword",
	}
	suite.mockUsecase.On("Login", mock.Anything, user).Return("", errors.New("invalid credentials"))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser","password":"testpassword"}`))
	ctx.Request.Header.Set("Content-Type", "application/json")

	suite.userController.Login(ctx)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Unauthorized access")
	assert.Contains(suite.T(), w.Body.String(), "details")
}

func (suite *UserControllerTestSuite) TestPromoteAdmin_Success() {
	username := "testuser"
	suite.mockUsecase.On("PromoteAdmin", mock.Anything, username).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{{Key: "username", Value: username}}

	suite.userController.PromoteAdmin(ctx)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "User promoted to admin successfully")
}

func (suite *UserControllerTestSuite) TestPromoteAdmin_InternalServerError() {
	username := "testuser"
	suite.mockUsecase.On("PromoteAdmin", mock.Anything, username).Return(errors.New("database error"))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{{Key: "username", Value: username}}

	suite.userController.PromoteAdmin(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Failed to promote user to admin")
	assert.Contains(suite.T(), w.Body.String(), "details")
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
