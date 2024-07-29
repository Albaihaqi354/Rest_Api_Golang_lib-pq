package response

type RoleResponse struct {
	Id          int    `json:"id"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}
