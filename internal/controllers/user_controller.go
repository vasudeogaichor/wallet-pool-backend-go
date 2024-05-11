package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vasudeogaichor/wallet-pool-backend/internal/db/models"
	"github.com/vasudeogaichor/wallet-pool-backend/internal/services"
	"github.com/vasudeogaichor/wallet-pool-backend/internal/validators"
)

type UserController struct {
	UserService services.UserService
}

func ConvertToGORMUser(validatorUser *validators.User) *models.User {
	return &models.User{
		Username: validatorUser.Username,
		Email:    validatorUser.Email,
		Password: validatorUser.Password,
	}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var validatorUser validators.User
	if c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		return
	}

	if err := c.ShouldBindJSON(&validatorUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidateUser(validatorUser); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gormUser := ConvertToGORMUser(&validatorUser)

	if err := ctrl.UserService.CreateUser(gormUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gormUser)
}
