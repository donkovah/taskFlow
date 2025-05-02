package userController

import (
	"be/src/domain/models"
	"be/src/domain/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}
func (tc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := tc.service.GetUser(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
	}
	c.JSON(http.StatusOK, user)
}

func (tc UserController) GetUsers(c *gin.Context) {
	users, err := tc.service.GetUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch user"})
	}
	c.JSON(http.StatusOK, users)
}

func (tc UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := tc.service.CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	}

	c.JSON(http.StatusOK, createdUser)
}

func (tc UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var userBody *models.User

	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.service.GetUser(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
	}
	user.Email = userBody.Email
	user.Username = userBody.Username

	updatedUser, err := tc.service.UpdateUser(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (ts UserController) DeleteUser(c *gin.Context) {
	id := c.Param(("id"))
	err := ts.service.DeleteUser(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
	}
	c.JSON(http.StatusNoContent, nil)
}
