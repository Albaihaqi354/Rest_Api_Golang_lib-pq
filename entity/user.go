package entity

import "time"

type User struct {
	Id         int
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Update_at  time.Time
}
