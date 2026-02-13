package services

import (
	"errors"

	"assignment2/models"
	"assignment2/repositories"

	"gorm.io/gorm"
)

type AccountService struct {
	DB              *gorm.DB
	AccountRepo     repositories.AccountRepo
	TransactionRepo repositories.TransactionRepo
}

func NewAccountService(db *gorm.DB) AccountService {
	return AccountService{
		DB:              db,
		AccountRepo:     repositories.AccountRepo{DB: db},
		TransactionRepo: repositories.TransactionRepo{DB: db},
	}
}

func (s AccountService) CreateAccount(a *models.Account) error {
	return s.AccountRepo.Create(a)
}

func (s AccountService) GetAccount(id uint) (models.Account, error) {
	return s.AccountRepo.GetByID(id)
}

func (s AccountService) Deposit(accountID uint, amount float64) error {

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {

		accRepo := repositories.AccountRepo{DB: tx}
		txRepo := repositories.TransactionRepo{DB: tx}

		acc, err := accRepo.GetByID(accountID)
		if err != nil {
			return err
		}

		if acc.Status != models.StatusActive {
			return errors.New("account inactive")
		}

		acc.Balance += amount

		if err := accRepo.Update(&acc); err != nil {
			return err
		}

		txn := models.Transaction{
			AccountID: accountID,
			Amount:    amount,
			Type:      "DEPOSIT",
		}

		return txRepo.Create(&txn)
	})
}
func (s AccountService) Delete(id uint) error {
	return s.AccountRepo.Delete(id)
}

func (s AccountService) Withdraw(accountID uint, amount float64) error {

	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	return s.DB.Transaction(func(tx *gorm.DB) error {

		accRepo := repositories.AccountRepo{DB: tx}
		txRepo := repositories.TransactionRepo{DB: tx}

		acc, err := accRepo.GetByID(accountID)
		if err != nil {
			return err
		}

		if acc.Status != models.StatusActive {
			return errors.New("account inactive")
		}

		if acc.Balance < amount {
			return errors.New("insufficient balance")
		}

		acc.Balance -= amount

		if err := accRepo.Update(&acc); err != nil {
			return err
		}

		txn := models.Transaction{
			AccountID: accountID,
			Amount:    amount,
			Type:      "WITHDRAW",
		}

		return txRepo.Create(&txn)
	})
}
