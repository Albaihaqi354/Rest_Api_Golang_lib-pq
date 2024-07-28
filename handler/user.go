package handler

import (
	"net/http"
	"strconv"
	"time"
	"user-management/entity"
	"user-management/request"
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

func (h *userHandler) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := userRequest.Id.Int64()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format Id Not Valid",
		})
		return
	}

	now := time.Now().UTC()
	created_at := now
	update_at := now

	if userRequest.Created_at != "" {
		created_at, err = time.Parse(time.RFC3339, userRequest.Created_at)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Format Created At Not Valid",
			})
			return
		}
	}

	if userRequest.Update_at != "" {
		update_at, err = time.Parse(time.RFC3339, userRequest.Update_at)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Format Update At Not Valid",
			})
			return
		}
	}

	user := entity.User{
		Id:         int(id),
		Name:       userRequest.Name,
		Email:      userRequest.Email,
		Password:   userRequest.Password,
		Created_at: created_at,
		Update_at:  update_at,
	}

	newUser, err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := convertToUserResponse(*newUser)

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID: " + idString})
		return
	}

	var userRequest request.UserRequest

	err = c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()
	update_at := now

	if userRequest.Update_at != "" {
		update_at, err = time.Parse(time.RFC3339, userRequest.Update_at)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Format Update At Not Valid",
			})
			return
		}
	}

	user := entity.User{
		Id:        id,
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		Update_at: update_at,
	}

	updatedUser, err := h.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := convertToUserResponse(*updatedUser)

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID: " + idString})
		return
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
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
