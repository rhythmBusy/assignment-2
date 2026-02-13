package main

import (
	"assignment2/config"
	controllers "assignment2/controller"
	"assignment2/routes"
	"assignment2/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	r := gin.Default()

	// services
	userSvc := services.NewUserService(db)
	bankSvc := services.NewBankService(db)
	branchSvc := services.NewBranchService(db)
	accountSvc := services.NewAccountService(db)
	loanSvc := services.NewLoanService(db)

	// controllers
	ctrls := routes.Controllers{
		User:    controllers.NewUserController(userSvc),
		Bank:    controllers.NewBankController(bankSvc),
		Branch:  controllers.NewBranchController(branchSvc),
		Account: controllers.NewAccountController(accountSvc),
		Loan:    controllers.NewLoanController(loanSvc),
	}

	routes.RegisterRoutes(r, ctrls)

	r.Run(":8080")
}
