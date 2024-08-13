package usecases

import (
	"context"
	"task_manager/domain"
	"task_manager/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskUsecaseTestSuite struct {
	suite.Suite
	mockTaskRepo *mocks.TaskRepository
	taskUsecase  TaskUsecase
}

func (suite *TaskUsecaseTestSuite) SetupTest() {
	// Initialize the mock repository and use case
	suite.mockTaskRepo = new(mocks.TaskRepository)
	suite.taskUsecase = NewTaskUsecase(suite.mockTaskRepo)
}

// Test cases for each use case
func (suite *TaskUsecaseTestSuite) TestCreateTask() {
	ctx := context.Background()
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
	}

	suite.mockTaskRepo.On("CreateTask", ctx, task).Return(nil)

	err := suite.taskUsecase.CreateTask(ctx, task)
	assert.Nil(suite.T(), err)

	suite.mockTaskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseTestSuite) TestGetAllTasks() {
	ctx := context.Background()
	tasks := []domain.Task{
		{ID: "1", Title: "complete task 8", Description: "read about testing first", Completed: false},
		{ID: "2", Title: "get ready fot startup project", Description: "revise the clean architecture and jwt ", Completed: true},
	}

	suite.mockTaskRepo.On("GetAllTasks", ctx).Return(tasks, nil)

	result, err := suite.taskUsecase.GetAllTasks(ctx)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), tasks, result)

	suite.mockTaskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseTestSuite) TestGetTaskByID() {
	ctx := context.Background()
	task := &domain.Task{ID: "1", Title: "Task 1", Description: "Task 1 description", Completed: false}

	suite.mockTaskRepo.On("GetTaskByID", ctx, "1").Return(task, nil)

	result, err := suite.taskUsecase.GetTaskByID(ctx, "1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), task, result)

	suite.mockTaskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseTestSuite) TestUpdateTask() {
	ctx := context.Background()
	updatedTask := &domain.Task{ID: "1", Title: "Updated Task", Description: "Updated description", Completed: true}

	suite.mockTaskRepo.On("UpdateTask", ctx, "1", updatedTask).Return(nil)

	err := suite.taskUsecase.UpdateTask(ctx, "1", updatedTask)
	assert.Nil(suite.T(), err)

	suite.mockTaskRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseTestSuite) TestDeleteTask() {
	ctx := context.Background()

	suite.mockTaskRepo.On("DeleteTask", ctx, "1").Return(nil)

	err := suite.taskUsecase.DeleteTask(ctx, "1")
	assert.Nil(suite.T(), err)

	suite.mockTaskRepo.AssertExpectations(suite.T())
}

// Run the test suite
func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}
