package response

import "time"

type UserResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created at"`
	Update_at  time.Time `json:"update at"`
}
