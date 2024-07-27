package handler

import (
	"net/http"
	"strconv"
	"user-management/entity"
	"user-management/response"
	"user-management/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ViewUsers(c *gin.Context) {
	users, err := h.userService.ViewUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var usersResponse []response.UserResponse

	for _, response := range users {
		userResponse := convertToUserResponse(response)

		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usersResponse,
	})
}

func (h *userHandler) ViewUserById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID: " + idString})
		return
	}

	user, err := h.userService.ViewUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userResponse := convertToUserResponse(*user)

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func convertToUserResponse(user entity.User) response.UserResponse {
	return response.UserResponse{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		Created_at: user.Created_at,
		Update_at:  user.Update_at,
	}
}
