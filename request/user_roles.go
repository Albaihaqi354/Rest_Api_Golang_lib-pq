package request

import "encoding/json"

type UserRolesRequest struct {
	Id     json.Number `json:"id" binding:"required,number"`
	UserId int         `json:"userId" binding:"required,number"`
	RoleId int         `json:"roleId" binding:"required,number"`
}
