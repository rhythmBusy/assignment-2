package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type BankRepo struct {
	DB *gorm.DB
}

func (r BankRepo) Create(bank *models.Bank) error {
	return r.DB.Create(bank).Error
}

func (r BankRepo) GetByID(id uint) (models.Bank, error) {
	var bank models.Bank
	err := r.DB.Preload("Branches").First(&bank, id).Error
	return bank, err
}

func (r BankRepo) GetAll() ([]models.Bank, error) {
	var banks []models.Bank
	err := r.DB.Find(&banks).Error
	return banks, err
}

func (r BankRepo) Update(bank *models.Bank) error {
	return r.DB.Save(bank).Error
}

func (r BankRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Bank{}, id).Error
}
