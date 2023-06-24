package services

import (
	"context"
	"github.com/devfurkankizmaz/quiz-app/models"
)

type userService struct {
	userRepository models.UserRepository
}

func NewUserService(userRepository models.UserRepository) models.UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(ctx context.Context, arg models.CreateUserParams) (models.User, error) {
	user, err := s.userRepository.CreateUser(ctx, arg)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
