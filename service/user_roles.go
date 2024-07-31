package service

import (
	"user-management/entity"
	"user-management/repository"
)

type UserRolesService interface {
	ViewUserRoles() ([]entity.UserRoles, error)
	ViewUserRolesById(Id int) (*entity.UserRoles, error)
	CreateUserRoles(role entity.UserRoles) (*entity.UserRoles, error)
	UpdateUserRoles(role entity.UserRoles) (*entity.UserRoles, error)
	DeleteUserRoles(Id int) error
}

type userRolesService struct {
	userRolesRepository repository.UserRoleRepository
}

func NewUserRoleService(userRolesRepository repository.UserRoleRepository) *userRolesService {
	return &userRolesService{userRolesRepository}
}

func (s *userRolesService) ViewUserRoles() ([]entity.UserRoles, error) {
	userRoles, err := s.userRolesRepository.ViewUserRoles()
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (s *userRolesService) ViewUserRolesById(Id int) (*entity.UserRoles, error) {
	userRole, err := s.userRolesRepository.ViewUserRolesById(Id)
	if err != nil {
		return nil, err
	}
	return userRole, nil
}

func (s *userRolesService) CreateUserRoles(userRole entity.UserRoles) (*entity.UserRoles, error) {
	NewUserRole, err := s.userRolesRepository.CreateUserRoles(userRole)
	if err != nil {
		return nil, err
	}
	return NewUserRole, nil
}

func (s *userRolesService) UpdateUserRoles(userRole entity.UserRoles) (*entity.UserRoles, error) {
	updatedUserRole, err := s.userRolesRepository.UpdateUserRoles(userRole)
	if err != nil {
		return nil, err
	}
	return updatedUserRole, nil
}

func (s *userRolesService) DeleteUserRoles(Id int) error {
	err := s.userRolesRepository.DeleteUserRoles(Id)
	if err != nil {
		return err
	}
	return nil
}
