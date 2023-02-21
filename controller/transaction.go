package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/database"
	"github.com/meshachdamilare/banking-api/error"
	"github.com/meshachdamilare/banking-api/models"
	"net/http"
	"strconv"
)

func GetTransactions(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
		return
	}

	acc, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.AccountFetchError)
		return
	}
	var TxnLimit int
	limit, got := c.GetQuery("limit")
	if !got {
		TxnLimit = 10
	} else {
		TxnLimit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
	}
	txn, err := database.GetTransactions(acc.AccountNumber, TxnLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"transactions": txn,
	}
	c.JSON(http.StatusOK, data)
}

func GetTransactionByID(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
		return
	}
	acc, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.AccountFetchError)
		return
	}
	txnID := c.Param("txnID")
	txn, err := database.GetTransactionByID(acc.AccountNumber, txnID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, txn)
}

func GetWithdrawal(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
		return
	}
	acc, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.AccountFetchError)
		return
	}
	var TxnLimit int
	limit, got := c.GetQuery("limit")
	if !got {
		TxnLimit = 10
	} else {
		TxnLimit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
	}
	txn, err := database.GetTransactionsByType(models.Withdraw, acc.AccountNumber, TxnLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"transactions": txn,
	}
	c.JSON(http.StatusOK, data)
}

func GetDeposits(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
		return
	}
	acc, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.AccountFetchError)
		return
	}
	var TxnLimit int
	limit, got := c.GetQuery("limit")
	if !got {
		TxnLimit = 10
	} else {
		TxnLimit, err = strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
	}
	txn, err := database.GetTransactionsByType(models.Deposit, acc.AccountNumber, TxnLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	data := map[string]interface{}{
		"transactions": txn,
	}
	c.JSON(http.StatusOK, data)
}
