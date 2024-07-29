package handler

import (
	"net/http"
	"user-management/entity"
	"user-management/response"
	"user-management/service"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *roleHandler {
	return &roleHandler{roleService: roleService}
}

func (h *roleHandler) ViewRoles(c *gin.Context) {
	roles, err := h.roleService.ViewRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var rolesResponse []response.RoleResponse

	for _, response := range roles {
		roleResponse := convertToRolesResponse(response)

		rolesResponse = append(rolesResponse, roleResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": rolesResponse,
	})
}

func convertToRolesResponse(roles entity.Role) response.RoleResponse {
	return response.RoleResponse{
		Id:          roles.Id,
		RoleName:    roles.RoleName,
		Description: roles.Description,
	}
}
