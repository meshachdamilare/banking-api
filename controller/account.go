package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/database"
	"github.com/meshachdamilare/banking-api/error"
	"github.com/meshachdamilare/banking-api/models"
	"net/http"
	"strconv"
	"time"
)

func GetUserAccountData(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
	}
	acc, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.AccountFetchError)
	}
	txn, err := database.GetTransactions(acc.AccountNumber, 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	data := map[string]interface{}{
		"account":      acc,
		"transactions": txn,
	}
	c.JSON(http.StatusOK, data)

}

func Deposit(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
	}
	amount, got := c.GetQuery("amount")
	if !got {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query parameter amount required",
		})
		return
	}
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	amountInt := uint64(parsedAmount * 100)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Amount should be a float.",
		})
		return
	}
	result, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if err := database.UpdateAccountBalance(email, result.Balance+amountInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	txnID, err := database.CreateTransaction(models.Deposit, amount, result.AccountNumber)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to deposit amount.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          fmt.Sprintf("Amount %s deposited to account %s", amount, result.AccountNumber),
		"available_amount": float64(result.Balance + amountInt/100),
		"txn_id":           *txnID,
		"timestamp":        time.Now(),
	})
}

func Withdraw(c *gin.Context) {
	email, ok := c.MustGet("email").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, error.AuthorizationError)
	}
	amount, got := c.GetQuery("amount")
	if !got {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query parameter amount required",
		})
		return
	}
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	amountInt := uint64(parsedAmount * 100)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Amount should be a float.",
		})
		return
	}
	result, err := database.GetAccount(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if result.Balance < amountInt {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":          "Withdrawal amount more than account balance.",
			"available_amount": result.Balance,
		})
		return
	}
	if err := database.UpdateAccountBalance(email, result.Balance-amountInt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	txnID, err := database.CreateTransaction(models.Withdraw, amount, result.AccountNumber)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to withdraw amount.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          fmt.Sprintf("Amount %s withdrawed from account %s.", amount, result.AccountNumber),
		"available_amount": float64(result.Balance-amountInt) / 100,
		"txn_id":           *txnID,
		"timestamp":        time.Now(),
	})
}
