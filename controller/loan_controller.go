package controllers

import (
	"net/http"
	"strconv"

	"assignment2/dto"
	"assignment2/models"
	"assignment2/services"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	Service services.LoanService
}

func NewLoanController(s services.LoanService) LoanController {
	return LoanController{Service: s}
}

func (lc LoanController) Create(c *gin.Context) {
	var req dto.CreateLoanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan := models.Loan{
		Base: models.Base{
			Status: models.StatusActive,
		},
		UserID:    req.UserID,
		BranchID:  req.BranchID,
		Principal: req.Principal,
	}

	if err := lc.Service.CreateLoan(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func (lc LoanController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	loan, err := lc.Service.GetLoan(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func (lc LoanController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := lc.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "loan deleted"})
}

func (lc LoanController) Repay(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.RepayLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := lc.Service.RepayLoan(uint(id), req.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment recorded"})
}
