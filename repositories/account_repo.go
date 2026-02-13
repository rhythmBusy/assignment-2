package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type AccountRepo struct {
	DB *gorm.DB
}

func (r AccountRepo) Create(a *models.Account) error {
	return r.DB.Create(a).Error
}

func (r AccountRepo) GetByID(id uint) (models.Account, error) {
	var a models.Account
	err := r.DB.Preload("Transactions").
		First(&a, id).Error
	return a, err
}

func (r AccountRepo) GetByUser(userID uint) ([]models.Account, error) {
	var list []models.Account
	err := r.DB.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r AccountRepo) Update(a *models.Account) error {
	return r.DB.Save(a).Error
}

func (r AccountRepo) Delete(id uint) error {
	return r.DB.Delete(&models.Account{}, id).Error
}
