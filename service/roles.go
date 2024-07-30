package service

import (
	"user-management/entity"
	"user-management/repository"
)

type RoleService interface {
	ViewRoles() ([]entity.Role, error)
	ViewRolesById(Id int) (*entity.Role, error)
	CreateRoles(role entity.Role) (*entity.Role, error)
	UpdateRoles(role entity.Role) (*entity.Role, error)
	DeleteRoles(Id int) error
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

func (s *roleService) CreateRoles(roles entity.Role) (*entity.Role, error) {
	newRole, err := s.roleRepository.CreateRoles(roles)
	if err != nil {
		return nil, err
	}
	return newRole, nil
}

func (s *roleService) UpdateRoles(roles entity.Role) (*entity.Role, error) {
	updateRole, err := s.roleRepository.UpdateRoles(roles)
	if err != nil {
		return nil, err
	}
	return updateRole, nil
}

func (s *roleService) DeleteRoles(Id int) error {
	err := s.roleRepository.DeleteRoles(Id)
	if err != nil {
		return err
	}
	return nil
}
