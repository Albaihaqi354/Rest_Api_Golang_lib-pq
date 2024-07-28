package repository

import (
	"database/sql"
	"user-management/entity"
)

type UserRepository interface {
	ViewUsers() ([]entity.User, error)
	ViewUserById(Id int) (*entity.User, error)
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	DeleteUser(Id int) error
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

func (r *userRepository) CreateUser(user entity.User) (*entity.User, error) {
	err := r.db.QueryRow("INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Name, user.Email, user.Password, user.Created_at, user.Update_at).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user entity.User) (*entity.User, error) {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5",
		user.Name, user.Email, user.Password, user.Update_at, user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
