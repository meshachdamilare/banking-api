package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/database"
	"github.com/meshachdamilare/banking-api/error"
	"github.com/meshachdamilare/banking-api/models"
	"net/http"
	"net/mail"
	"time"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	user.CreatedAt = time.Now()
	if _, err := mail.ParseAddress(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email address",
		})
		return
	}
	accNumber, err := database.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":        fmt.Sprintf("User '%s' created.", user.Email),
		"Account Number": accNumber,
	})

}

func GetUserData(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
		return
	}

	result, err := database.GetUser(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.UserFetchError)
		return
	}
	c.JSON(http.StatusOK, result)
}
