package services

import (
	"assignment2/models"
	"assignment2/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	Repo repositories.UserRepo
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{
		Repo: repositories.UserRepo{DB: db},
	}
}

func (s UserService) Create(u *models.User) error {
	return s.Repo.Create(u)
}

func (s UserService) Get(id uint) (models.User, error) {
	return s.Repo.GetByID(id)
}

func (s UserService) GetAll() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s UserService) Update(u *models.User) error {
	return s.Repo.Update(u)
}

func (s UserService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
