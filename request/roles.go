package request

import "encoding/json"

type Rolesrequest struct {
	Id          json.Number `json:"id" binding:"required,number"`
	RoleName    string      `json:"role_name" binding:"required"`
	Description string      `json:"description" binding:"required"`
}
