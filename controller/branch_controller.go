package controllers

import (
	"net/http"
	"strconv"

	"assignment2/dto"
	"assignment2/models"
	"assignment2/services"

	"github.com/gin-gonic/gin"
)

type BranchController struct {
	Service services.BranchService
}

func NewBranchController(s services.BranchService) BranchController {
	return BranchController{Service: s}
}

func (bc BranchController) Create(c *gin.Context) {
	var req dto.CreateBranchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch := models.Branch{
		Name:    req.Name,
		IFSC:    req.IFSC,
		Address: req.Address,
		BankID:  req.BankID,
	}

	if err := bc.Service.Create(&branch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

func (bc BranchController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	branch, err := bc.Service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "branch not found"})
		return
	}

	c.JSON(http.StatusOK, branch)
}

func (bc BranchController) GetByBank(c *gin.Context) {
	bankID, _ := strconv.Atoi(c.Param("bankId"))

	list, err := bc.Service.GetByBank(uint(bankID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (bc BranchController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.CreateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch, err := bc.Service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "branch not found"})
		return
	}

	branch.Name = req.Name
	branch.IFSC = req.IFSC
	branch.Address = req.Address
	branch.BankID = req.BankID

	if err := bc.Service.Update(&branch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branch)
}

func (bc BranchController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := bc.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
