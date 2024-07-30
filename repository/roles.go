package repository

import (
	"database/sql"
	"user-management/entity"
)

type RoleRepository interface {
	ViewRoles() ([]entity.Role, error)
	ViewRolesById(Id int) (*entity.Role, error)
	CreateRoles(role entity.Role) (*entity.Role, error)
	UpdateRoles(role entity.Role) (*entity.Role, error)
	DeleteRoles(Id int) error
}

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *roleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) ViewRoles() ([]entity.Role, error) {
	var roles []entity.Role
	rows, err := r.db.Query("SELECT id, role_name, description FROM roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role entity.Role
		err := rows.Scan(&role.Id, &role.RoleName, &role.Description)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *roleRepository) ViewRolesById(Id int) (*entity.Role, error) {
	var role entity.Role
	err := r.db.QueryRow("SELECT id, role_name, description FROM roles WHERE id = $1", Id).Scan(&role.Id, &role.RoleName, &role.Description)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) CreateRoles(role entity.Role) (*entity.Role, error) {
	err := r.db.QueryRow("INSERT INTO roles (role_name, description) VALUES ($1, $2) RETURNING id", role.RoleName, role.Description).Scan(&role.Id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) UpdateRoles(role entity.Role) (*entity.Role, error) {
	_, err := r.db.Exec("UPDATE roles SET role_name = $1, description = $2 WHERE id = $3", role.RoleName, role.Description, role.Id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) DeleteRoles(Id int) error {
	_, err := r.db.Exec("DELETE FROM roles WHERE id = $1", Id)
	return err
}
