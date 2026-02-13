package routes

import (
	controllers "assignment2/controller" // alias explicitly

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	Account controllers.AccountController
	User    controllers.UserController
	Bank    controllers.BankController
	Branch  controllers.BranchController
	Loan    controllers.LoanController
}

func RegisterRoutes(r *gin.Engine, c Controllers) {

	api := r.Group("/api")

	// USERS
	api.POST("/users", c.User.Create)
	api.GET("/users/:id", c.User.Get)
	api.DELETE("/users/:id", c.User.Delete)

	// BANKS
	api.POST("/banks", c.Bank.Create)
	api.GET("/banks/:id", c.Bank.Get)

	// BRANCH
	api.POST("/branches", c.Branch.Create)
	api.GET("/branches/:id", c.Branch.Get)

	// ACCOUNTS
	api.POST("/accounts", c.Account.Create)
	api.GET("/accounts/:id", c.Account.Get)
	api.DELETE("/accounts/:id", c.Account.Delete)
	api.POST("/accounts/:id/deposit", c.Account.Deposit)
	api.POST("/accounts/:id/withdraw", c.Account.Withdraw)

	// LOANS
	api.POST("/loans", c.Loan.Create)
	api.GET("/loans/:id", c.Loan.Get)
	api.DELETE("/loans/:id", c.Loan.Delete)
	api.POST("/loans/:id/repay", c.Loan.Repay)
}
