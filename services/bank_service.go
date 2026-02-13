package services

import (
	"assignment2/models"
	"assignment2/repositories"

	"gorm.io/gorm"
)

type BankService struct {
	Repo repositories.BankRepo
}

func NewBankService(db *gorm.DB) BankService {
	return BankService{
		Repo: repositories.BankRepo{DB: db},
	}
}

func (s BankService) Create(bank *models.Bank) error {
	return s.Repo.Create(bank)
}

func (s BankService) Get(id uint) (models.Bank, error) {
	return s.Repo.GetByID(id)
}

func (s BankService) GetAll() ([]models.Bank, error) {
	return s.Repo.GetAll()
}

func (s BankService) Update(bank *models.Bank) error {
	return s.Repo.Update(bank)
}

func (s BankService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
