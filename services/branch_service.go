package services

import (
	"assignment2/models"
	"assignment2/repositories"

	"gorm.io/gorm"
)

type BranchService struct {
	Repo repositories.BranchRepo
}

func NewBranchService(db *gorm.DB) BranchService {
	return BranchService{
		Repo: repositories.BranchRepo{DB: db},
	}
}

func (s BranchService) Create(b *models.Branch) error {
	return s.Repo.Create(b)
}

func (s BranchService) Get(id uint) (models.Branch, error) {
	return s.Repo.GetByID(id)
}

func (s BranchService) GetByBank(bankID uint) ([]models.Branch, error) {
	return s.Repo.GetByBank(bankID)
}

func (s BranchService) Update(b *models.Branch) error {
	return s.Repo.Update(b)
}

func (s BranchService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
