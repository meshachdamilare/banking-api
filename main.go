package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/database"
	"github.com/meshachdamilare/banking-api/error"
	"github.com/meshachdamilare/banking-api/route"
	"log"
	"net/http"
	"time"
)

func main() {
	err := database.MigratePostgres()
	if err != nil {
		panic(err)
	}
	err = database.MigrateImmuDB()
	if err != nil {
		log.Println(err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, error.EndpointNotFound)
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusForbidden, error.MethodNotAllowed)
	})

	router.GET("/api/v1", root)
	route.BankRoute(router)

	log.Println("API running on http://localhost:8002")
	log.Fatal(router.Run(":8080"))

}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "Banking API is working.",
		"timestamp": time.Now(),
	})
}
