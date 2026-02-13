package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func (r TransactionRepo) Create(t *models.Transaction) error {
	return r.DB.Create(t).Error
}

func (r TransactionRepo) GetByAccount(accountID uint) ([]models.Transaction, error) {
	var list []models.Transaction
	err := r.DB.Where("account_id = ?", accountID).
		Order("created_at desc").
		Find(&list).Error
	return list, err
}
