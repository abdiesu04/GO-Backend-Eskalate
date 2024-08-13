package usecases

import (
	"context"
	"errors"
	"task_manager/domain"
	"task_manager/infrastructure"
	"task_manager/repositories"
)

type UserUsecase interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, user *domain.User) (string, error)
	PromoteAdmin(ctx context.Context, username string) error
}

type userUsecase struct {
	UserRepository repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: repo,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *domain.User) error {
	exists, err := u.UserRepository.IsUserExist(ctx, user.Username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("username already exists")
	}

	isEmpty, err := u.UserRepository.IsDBEmpty(ctx)
	if err != nil {
		return err
	}

	if isEmpty {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// Hash the password
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set the hashed password
	user.Password = hashedPassword

	// Register the user
	return u.UserRepository.CreateUser(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, user *domain.User) (string, error) {
	storedPassword, err := u.UserRepository.GetPasswrodByUsername(ctx, user.Username)
	if err != nil {
		return "", err
	}

	if !infrastructure.ComparePassword(storedPassword, user.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := infrastructure.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUsecase) PromoteAdmin(ctx context.Context, username string) error {
	return u.UserRepository.PromoteAdmin(ctx, username)
}
