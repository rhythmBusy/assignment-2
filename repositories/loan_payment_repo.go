package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type LoanPaymentRepo struct {
	DB *gorm.DB
}

func (r LoanPaymentRepo) Create(p *models.LoanPayment) error {
	return r.DB.Create(p).Error
}

func (r LoanPaymentRepo) GetByLoan(loanID uint) ([]models.LoanPayment, error) {
	var list []models.LoanPayment
	err := r.DB.Where("loan_id = ?", loanID).
		Order("created_at desc").
		Find(&list).Error
	return list, err
}
