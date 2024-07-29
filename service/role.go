package service

import (
	"user-management/entity"
	"user-management/repository"
)

type RoleService interface {
	ViewRoles() ([]entity.Role, error)
	ViewRolesById(Id int) (*entity.Role, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) *roleService {
	return &roleService{roleRepository}
}

func (s *roleService) ViewRoles() ([]entity.Role, error) {
	roles, err := s.roleRepository.ViewRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *roleService) ViewRolesById(Id int) (*entity.Role, error) {
	roles, err := s.roleRepository.ViewRolesById(Id)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
