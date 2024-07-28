package request

import (
	"encoding/json"
)

type UserRequest struct {
	Id         json.Number `json:"id" binding:"required,number"`
	Name       string      `json:"name" binding:"required"`
	Email      string      `json:"email" binding:"required,email"`
	Password   string      `json:"password" binding:"required"`
	Created_at string      `json:"created_at"`
	Update_at  string      `json:"update_at"`
}
