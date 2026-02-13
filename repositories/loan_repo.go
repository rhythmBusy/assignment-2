package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type LoanRepo struct {
	DB *gorm.DB
}

func (r LoanRepo) Create(l *models.Loan) error {
	return r.DB.Create(l).Error
}

func (r LoanRepo) GetByID(id uint) (models.Loan, error) {
	var l models.Loan
	err := r.DB.Preload("Payments").
		First(&l, id).Error
	return l, err
}

func (r LoanRepo) GetByUser(userID uint) ([]models.Loan, error) {
	var list []models.Loan
	err := r.DB.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r LoanRepo) Update(l *models.Loan) error {
	return r.DB.Save(l).Error
}

func (r LoanRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Loan{}, id).Error
}
