package repository

import (
	"database/sql"
	"user-management/entity"
)

type RoleRepository interface {
	ViewRoles() ([]entity.Role, error)
	ViewRolesById(Id int) (*entity.Role, error)
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
