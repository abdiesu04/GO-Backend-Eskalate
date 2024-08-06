package data

import (
    "context"
    "errors"
    "task_manager/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    Collection *mongo.Collection // MongoDB collection for users
}

// CreateUser inserts a new user into the MongoDB collection
func (s *UserService) CreateUser(user *models.User) error {
    var existingUser models.User
    err := s.Collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
    if err == nil {
        return errors.New("username already exists")
    }
    if err != mongo.ErrNoDocuments {
        return err
    }

    // Check if there are any users in the collection
    userCount, err := s.Collection.CountDocuments(context.TODO(), bson.M{})
    if err != nil {
        return err
    }

    // Assign role based on the user count
    if userCount == 0 {
        user.Role = "admin"
    } else {
        user.Role = "user"
    }

    _, err = s.Collection.InsertOne(context.TODO(), user)
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

// UpdateUserRole updates the role of a user
func (s *UserService) UpdateUserRole(ctx context.Context, username string, role string) error {
    filter := bson.M{"username": username}
    update := bson.M{"$set": bson.M{"role": role}}

    result, err := s.Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    
    if result.MatchedCount == 0 {
        return err
    }

    return nil
}
