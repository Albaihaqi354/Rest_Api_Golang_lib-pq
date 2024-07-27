package request

import (
	"encoding/json"
	"time"
)

type UserRequest struct {
	Id         json.Number `json:"id" binding:"required,number"`
	Name       string      `json:"username" binding:"required"`
	Email      string      `json:"email" binding:"required,email"`
	Password   string      `json:"password" binding:"required"`
	Created_at time.Time   `json:"created_at"`
	Update_at  time.Time   `json:"update_at"`
}
