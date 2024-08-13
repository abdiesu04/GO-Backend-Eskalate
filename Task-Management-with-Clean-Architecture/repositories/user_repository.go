package repositories

import (
	"context"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	PromoteAdmin(ctx context.Context, username string) error
	GetPasswrodByUsername(ctx context.Context, username string) (string, error)
	IsDBEmpty(ctx context.Context) (bool, error)
	IsUserExist(ctx context.Context, username string) (bool, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) PromoteAdmin(ctx context.Context, username string) error {
	filter := bson.D{{Key: "username", Value: username}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: "admin"}}}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *userRepository) GetPasswrodByUsername(ctx context.Context, username string) (string, error) {
	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func (r *userRepository) IsUserExist(ctx context.Context, username string) (bool, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		return false, err
	}
	defer cursor.Close(ctx)
	return cursor.Next(ctx), nil
}

func (r *userRepository) IsDBEmpty(ctx context.Context) (bool, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	defer cursor.Close(ctx)
	return !cursor.Next(ctx), nil
}
