package controllers

import (
	"net/http"
	"strconv"

	"assignment2/repositories"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	Repo repositories.TransactionRepo
}

func NewTransactionController(r repositories.TransactionRepo) TransactionController {
	return TransactionController{Repo: r}
}

// GET /accounts/:id/transactions
func (tc TransactionController) GetByAccount(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	list, err := tc.Repo.GetByAccount(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}
