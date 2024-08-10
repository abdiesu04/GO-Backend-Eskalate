package repositories

import (
	"context"
	"strconv"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	UpdateTask(ctx context.Context, id string, updatedTask *domain.Task) error
	DeleteTask(ctx context.Context, id string) error
}

type taskRepository struct {
	collection *mongo.Collection
	counterCol *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepository{
		collection: db.Collection("tasks"),
		counterCol: db.Collection("counters"),
	}
}

// getNextID generates a new incremented string ID
func (r *taskRepository) getNextID(ctx context.Context) (string, error) {
	filter := bson.D{{Key: "_id", Value: "task_id"}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "seq", Value: 1}}}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result struct {
		Seq int `bson:"seq"`
	}

	err := r.counterCol.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(result.Seq), nil
}

func (r *taskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
	// Generate the next ID
	id, err := r.getNextID(ctx)
	if err != nil {
		return err
	}
	task.ID = id

	_, err = r.collection.InsertOne(ctx, task)
	return err
}

func (r *taskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, id string, updatedTask *domain.Task) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: updatedTask.Title},
			{Key: "description", Value: updatedTask.Description},
			{Key: "completed", Value: updatedTask.Completed},
		}},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *taskRepository) DeleteTask(ctx context.Context, id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
