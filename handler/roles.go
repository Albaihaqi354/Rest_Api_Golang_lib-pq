package handler

import (
	"net/http"
	"strconv"
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

func (h *roleHandler) ViewRolesById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID" + idString,
		})
		return
	}

	role, err := h.roleService.ViewRolesById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if role == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Role not found",
		})
		return
	}

	rolesResponse := convertToRolesResponse(*role)

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
