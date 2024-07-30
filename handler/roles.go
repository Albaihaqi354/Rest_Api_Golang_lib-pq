package handler

import (
	"net/http"
	"strconv"
	"user-management/entity"
	"user-management/request"
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

func (h *roleHandler) CreateRole(c *gin.Context) {
	var roleRequest request.Rolesrequest

	err := c.ShouldBindJSON(&roleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := roleRequest.Id.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	role := entity.Role{
		Id:          int(id),
		RoleName:    roleRequest.RoleName,
		Description: roleRequest.Description,
	}

	newRole, err := h.roleService.CreateRoles(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResponse := convertToRolesResponse(*newRole)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})

}

func (h *roleHandler) UpdateRole(c *gin.Context) {
	var roleRequest request.Rolesrequest

	err := c.ShouldBindJSON(&roleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := roleRequest.Id.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	role := entity.Role{
		Id:          int(id),
		RoleName:    roleRequest.RoleName,
		Description: roleRequest.Description,
	}

	updatedRole, err := h.roleService.UpdateRoles(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResponse := convertToRolesResponse(*updatedRole)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *roleHandler) DeleteRole(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID" + idString,
		})
		return
	}

	err = h.roleService.DeleteRoles(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role deleted successfully",
	})
}

func convertToRolesResponse(roles entity.Role) response.RoleResponse {
	return response.RoleResponse{
		Id:          roles.Id,
		RoleName:    roles.RoleName,
		Description: roles.Description,
	}
}
