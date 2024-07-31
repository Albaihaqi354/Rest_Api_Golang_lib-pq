package repository

import (
	"database/sql"
	"user-management/entity"
)

type UserRoleRepository interface {
	ViewUserRoles() ([]entity.UserRoles, error)
	ViewUserRolesById(Id int) (*entity.UserRoles, error)
	CreateUserRoles(userRoles entity.UserRoles) (*entity.UserRoles, error)
	UpdateUserRoles(userRoles entity.UserRoles) (*entity.UserRoles, error)
	DeleteUserRoles(Id int) error
}

type userRoleRepository struct {
	db *sql.DB
}

func NewUserRoleRepository(db *sql.DB) *userRoleRepository {
	return &userRoleRepository{db}
}

func (r *userRoleRepository) ViewUserRoles() ([]entity.UserRoles, error) {
	rows, err := r.db.Query("SELECT * FROM user_roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []entity.UserRoles
	for rows.Next() {
		var userRole entity.UserRoles
		err := rows.Scan(&userRole.Id, &userRole.UserId, &userRole.RoleId)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, userRole)
	}

	return userRoles, nil
}

func (r *userRoleRepository) ViewUserRolesById(Id int) (*entity.UserRoles, error) {
	var userRole entity.UserRoles
	err := r.db.QueryRow("SELECT * FROM user_roles WHERE id = $1", Id).Scan(&userRole.Id, &userRole.UserId, &userRole.RoleId)
	if err != nil {
		return nil, err
	}

	return &userRole, nil
}

func (r *userRoleRepository) CreateUserRoles(userRoles entity.UserRoles) (*entity.UserRoles, error) {
	err := r.db.QueryRow("INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2) RETURNING id", userRoles.UserId, userRoles.RoleId).Scan(&userRoles.Id)
	if err != nil {
		return nil, err
	}

	return &userRoles, nil
}

func (r *userRoleRepository) UpdateUserRoles(userRoles entity.UserRoles) (*entity.UserRoles, error) {
	_, err := r.db.Exec("UPDATE user_roles SET user_id = $1, role_id = $2 WHERE id = $3", userRoles.UserId, userRoles.RoleId, userRoles.Id)
	if err != nil {
		return nil, err
	}

	return &userRoles, nil
}

func (r *userRoleRepository) DeleteUserRoles(Id int) error {
	_, err := r.db.Exec("DELETE FROM user_roles WHERE id = $1", Id)
	if err != nil {
		return err
	}

	return nil
}
