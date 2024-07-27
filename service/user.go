package service

import (
	"user-management/entity"
	"user-management/repository"
)

type UserService interface {
	ViewUsers() ([]entity.User, error)
	ViewUserById(Id int) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) ViewUsers() ([]entity.User, error) {
	users, err := s.userRepository.ViewUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) ViewUserById(Id int) (*entity.User, error) {
	user, err := s.userRepository.ViewUserById(Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
