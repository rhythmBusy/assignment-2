package controllers

import (
	"net/http"
	"strconv"

	"assignment2/dto"
	"assignment2/models"
	"assignment2/services"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	Service services.BankService
}

func NewBankController(s services.BankService) BankController {
	return BankController{Service: s}
}

func (bc BankController) Create(c *gin.Context) {
	var req dto.CreateBankRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bank := models.Bank{
		Name: req.Name,
		Code: req.Code,
	}

	if err := bc.Service.Create(&bank); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bank)
}

func (bc BankController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	bank, err := bc.Service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bank not found"})
		return
	}

	c.JSON(http.StatusOK, bank)
}

func (bc BankController) GetAll(c *gin.Context) {
	list, err := bc.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (bc BankController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.UpdateBankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bank, err := bc.Service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bank not found"})
		return
	}

	if req.Name != "" {
		bank.Name = req.Name
	}
	if req.Code != "" {
		bank.Code = req.Code
	}

	if err := bc.Service.Update(&bank); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bank)
}

func (bc BankController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := bc.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
