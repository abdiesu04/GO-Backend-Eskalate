package data

import (
	"context"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		collection: collection}
}

func (s *TaskService) GetTasks(ctx context.Context) ([]models.Task, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
	var task models.Task
	if err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&task); err != nil {
		return task, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	// Generate a new integer ID (this is a simple example, you might want to use a more robust method)
	var lastTask models.Task
	err := s.collection.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.D{{Key: "id", Value: -1}})).Decode(&lastTask)
	if err != nil && err != mongo.ErrNoDocuments {
		return task, err
	}
	task.ID = lastTask.ID + 1

	_, err = s.collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, updatedTask models.Task) (models.Task, error) {
	update := bson.M{"$set": updatedTask}
	_, err := s.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return updatedTask, err
	}
	updatedTask.ID = id
	return updatedTask, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
