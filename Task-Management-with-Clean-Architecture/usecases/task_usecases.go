package usecases

import (
	"context"
	"task_manager/domain"
	"task_manager/repositories"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	UpdateTask(ctx context.Context, id string, updatedTask *domain.Task) error
	DeleteTask(ctx context.Context, id string) error
}

type taskUsecase struct {
	TaskRepo repositories.TaskRepository
}

func NewTaskUsecase(repo repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{
		TaskRepo: repo,
	}
}

func (u *taskUsecase) CreateTask(ctx context.Context, task *domain.Task) error {
	return u.TaskRepo.CreateTask(ctx, task)
}

func (u *taskUsecase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
    return u.TaskRepo.GetAllTasks(ctx)
}

func (u *taskUsecase) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
    return u.TaskRepo.GetTaskByID(ctx, id)
}

func (u *taskUsecase) UpdateTask(ctx context.Context, id string, updatedTask *domain.Task) error {
    return u.TaskRepo.UpdateTask(ctx, id, updatedTask)
}

func (u *taskUsecase) DeleteTask(ctx context.Context, id string) error {
    return u.TaskRepo.DeleteTask(ctx, id)
}
