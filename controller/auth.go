package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/database"
	"github.com/meshachdamilare/banking-api/models"
	"github.com/meshachdamilare/banking-api/util"
	"net/http"
	"time"
)

func AuthHandler(c *gin.Context) {
	var user models.UserAuth
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := database.AuthUser(&user); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("Authentication failed: %s", err),
		})
		return
	}
	tokenString, err := util.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	cookie := http.Cookie{
		Name:    "JWT",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 30),
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, gin.H{
		"message": "User authorized.",
		"token":   tokenString,
	})

}
