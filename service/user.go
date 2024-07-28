package service

import (
	"user-management/entity"
	"user-management/repository"
)

type UserService interface {
	ViewUsers() ([]entity.User, error)
	ViewUserById(Id int) (*entity.User, error)
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	DeleteUser(Id int) error
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

func (s *userService) CreateUser(user entity.User) (*entity.User, error) {
	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *userService) UpdateUser(user entity.User) (*entity.User, error) {
	updatedUser, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(Id int) error {
	err := s.userRepository.DeleteUser(Id)
	if err != nil {
		return err
	}
	return nil
}
