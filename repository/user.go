package repository

import (
	"database/sql"
	"user-management/entity"
)

type UserRepository interface {
	ViewUsers() ([]entity.User, error)
	ViewUserById(Id int) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) ViewUsers() ([]entity.User, error) {
	var users []entity.User
	rows, err := r.db.Query("SELECT id, name, email, password, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Update_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) ViewUserById(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Update_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
