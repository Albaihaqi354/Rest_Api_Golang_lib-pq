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

type userRoleHandler struct {
	userRoleService service.UserRolesService
}

func NewUserRoleHandler(userRoleService service.UserRolesService) *userRoleHandler {
	return &userRoleHandler{userRoleService: userRoleService}
}

func (h *userRoleHandler) ViewUserRoles(c *gin.Context) {
	userRoles, err := h.userRoleService.ViewUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userRolesResponse []response.UserRolesResponse

	for _, response := range userRoles {
		userRoleResponse := convertToUserRolesResponse(response)

		userRolesResponse = append(userRolesResponse, userRoleResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userRolesResponse,
	})

}

func (h *userRoleHandler) ViewUserRolesById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	userRole, err := h.userRoleService.ViewUserRolesById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if userRole == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Role not found",
		})
		return
	}

	userRoleResponse := convertToUserRolesResponse(*userRole)
	c.JSON(http.StatusOK, gin.H{
		"data": userRoleResponse,
	})
}

func (h *userRoleHandler) CreateUserRoles(c *gin.Context) {
	var userRoleRequest request.UserRolesRequest

	err := c.ShouldBindJSON(&userRoleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := userRoleRequest.Id.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	userId, err := userRoleRequest.UserId.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	roleId, err := userRoleRequest.RoleId.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	userRole := entity.UserRoles{
		Id:     int(id),
		UserId: int(userId),
		RoleId: int(roleId),
	}

	newUserRole, err := h.userRoleService.CreateUserRoles(userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userRoleResponse := convertToUserRolesResponse(*newUserRole)
	c.JSON(http.StatusOK, gin.H{
		"data": userRoleResponse,
	})
}

func (h *userRoleHandler) UpdateUserRoles(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var userRoleRequest request.UserRolesRequest

	err = c.ShouldBindJSON(&userRoleRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId, err := userRoleRequest.UserId.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	roleId, err := userRoleRequest.RoleId.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	userRole := entity.UserRoles{
		Id:     id,
		UserId: int(userId),
		RoleId: int(roleId),
	}

	updatedUserRole, err := h.userRoleService.UpdateUserRoles(userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userRoleResponse := convertToUserRolesResponse(*updatedUserRole)
	c.JSON(http.StatusOK, gin.H{
		"data": userRoleResponse,
	})
}

func (h *userRoleHandler) DeleteUserRoles(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	err = h.userRoleService.DeleteUserRoles(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Role deleted successfully",
	})
}

func convertToUserRolesResponse(userRoles entity.UserRoles) response.UserRolesResponse {
	return response.UserRolesResponse{
		Id:     userRoles.Id,
		UserId: userRoles.UserId,
		RoleId: userRoles.RoleId,
	}
}
