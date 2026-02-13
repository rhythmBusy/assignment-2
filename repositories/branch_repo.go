package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type BranchRepo struct {
	DB *gorm.DB
}

func (r BranchRepo) Create(b *models.Branch) error {
	return r.DB.Create(b).Error
}

func (r BranchRepo) GetByID(id uint) (models.Branch, error) {
	var b models.Branch
	err := r.DB.First(&b, id).Error
	return b, err
}

func (r BranchRepo) GetByBank(bankID uint) ([]models.Branch, error) {
	var list []models.Branch
	err := r.DB.Where("bank_id = ?", bankID).Find(&list).Error
	return list, err
}

func (r BranchRepo) Update(b *models.Branch) error {
	return r.DB.Save(b).Error
}

func (r BranchRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Branch{}, id).Error
}
