package response

type UserRolesResponse struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}
