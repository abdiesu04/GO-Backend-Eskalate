package data

import (
    "context"
    "task_manager/models"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
    Collection *mongo.Collection
}

func (s *UserService) CreateUser(user *models.User) error {
    _, err := s.Collection.InsertOne(context.TODO(), user)
    return err
}


// GetUserByUsername retrieves a user by username from MongoDB
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    err := s.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// ValidatePassword checks if the provided password matches the hashed password
func (s *UserService) ValidatePassword(user *models.User, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}
