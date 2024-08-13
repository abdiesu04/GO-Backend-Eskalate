package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager/domain"
	"task_manager/usecases/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskControllerTestSuite struct {
	suite.Suite
	mockTaskUsecase *mocks.TaskUsecase
	router          *gin.Engine
	controller      *TaskController
}

func (suite *TaskControllerTestSuite) SetupTest() {
	suite.mockTaskUsecase = new(mocks.TaskUsecase)
	suite.controller = NewTaskController(suite.mockTaskUsecase)

	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.router.POST("/tasks", suite.controller.CreateTask)
	suite.router.GET("/tasks", suite.controller.GetAllTasks)
	suite.router.GET("/tasks/:id", suite.controller.GetTaskByID)
	suite.router.PUT("/tasks/:id", suite.controller.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.controller.DeleteTask)
}

func (suite *TaskControllerTestSuite) TestCreateTask_Success() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
	}

	suite.mockTaskUsecase.On("CreateTask", mock.Anything, task).Return(nil)

	taskJSON, _ := json.Marshal(task)
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
	suite.NoError(err)

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Equal(http.StatusCreated, recorder.Code)
	suite.mockTaskUsecase.AssertExpectations(suite.T())
}

// test for invalid json
func (suite *TaskControllerTestSuite) TestCreateTask_BadRequest() {
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte("invalid json")))
	suite.NoError(err)

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Equal(http.StatusBadRequest, recorder.Code)
}

// test when usecases returns an error
func (suite *TaskControllerTestSuite) TestCreateTask_InternalServerError() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
	}

	suite.mockTaskUsecase.On("CreateTask", mock.Anything, task).Return(errors.New("database error"))

	taskJSON, _ := json.Marshal(task)
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJSON))
	suite.NoError(err)

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Equal(http.StatusInternalServerError, recorder.Code)
}

func (suite *TaskControllerTestSuite) TestGetAllTasks_Success() {
	tasks := []domain.Task{
		{ID: "1", Title: "Task 1", Description: "Desc 1", Completed: false},
		{ID: "2", Title: "Task 2", Description: "Desc 2", Completed: true},
	}

	suite.mockTaskUsecase.On("GetAllTasks", mock.Anything).Return(tasks, nil)

	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	suite.NoError(err)

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Equal(http.StatusOK, recorder.Code)
	suite.mockTaskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetAllTasks_InternalServerError() {
	suite.mockTaskUsecase.On("GetAllTasks", mock.Anything).Return(nil, errors.New("database error"))

	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	suite.NoError(err)

	recorder := httptest.NewRecorder()
	suite.router.ServeHTTP(recorder, req)

	suite.Equal(http.StatusInternalServerError, recorder.Code)
	suite.JSONEq(`{"error":"database error"}`, recorder.Body.String())
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
