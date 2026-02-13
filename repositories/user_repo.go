package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r UserRepo) Create(u *models.User) error {
	return r.DB.Create(u).Error
}

func (r UserRepo) GetByID(id uint) (models.User, error) {
	var u models.User
	err := r.DB.Preload("Accounts").
		Preload("Loans").
		First(&u, id).Error
	return u, err
}

func (r UserRepo) GetAll() ([]models.User, error) {
	var list []models.User
	err := r.DB.Find(&list).Error
	return list, err
}

func (r UserRepo) Update(u *models.User) error {
	return r.DB.Save(u).Error
}

func (r UserRepo) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
