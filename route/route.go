package route

import (
	"github.com/gin-gonic/gin"
	"github.com/meshachdamilare/banking-api/controller"
)

func BankRoute(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", controller.AuthHandler)
		user := v1.Group("/user")
		{
			user.POST("/", controller.CreateUser)
			user.GET("/", controller.GetUserData)

		}
		account := v1.Group("/account")
		{
			account.GET("/", controller.GetUserAccountData)
			account.POST("/deposit", controller.Deposit)
			account.POST("/withdraw", controller.Withdraw)
		}
		transaction := v1.Group("/transactions")
		{
			transaction.GET("/", controller.GetTransactions)
			transaction.GET("/:txnID", controller.GetTransactionByID)
			transaction.GET("/deposits", controller.GetDeposits)
			transaction.GET("/withdraws", controller.GetWithdrawal)
		}
	}
}
