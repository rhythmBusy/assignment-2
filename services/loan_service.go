package services

import (
	"errors"
	"time"

	"assignment2/models"
	"assignment2/repositories"

	"gorm.io/gorm"
)

type LoanService struct {
	DB          *gorm.DB
	LoanRepo    repositories.LoanRepo
	PaymentRepo repositories.LoanPaymentRepo
}

func NewLoanService(db *gorm.DB) LoanService {
	return LoanService{
		DB:          db,
		LoanRepo:    repositories.LoanRepo{DB: db},
		PaymentRepo: repositories.LoanPaymentRepo{DB: db},
	}
}

// CREATE LOAN
func (s LoanService) CreateLoan(l *models.Loan) error {
	if l.Principal <= 0 {
		return errors.New("principal must be positive")
	}

	// fixed 12% interest
	l.InterestRate = 12

	l.RemainingPrincipal = l.Principal
	l.InterestDue = l.Principal * (l.InterestRate / 100)
	l.InterestThisYear = l.InterestDue
	return s.LoanRepo.Create(l)
}

// GET LOAN
func (s LoanService) GetLoan(id uint) (models.Loan, error) {
	return s.LoanRepo.GetByID(id)
}

// DELETE LOAN (soft delete if Base has gorm.DeletedAt — hard delete otherwise)
func (s LoanService) Delete(id uint) error {
	return s.LoanRepo.Delete(id)
}

// REPAY LOAN — interest first, then principal
func (s LoanService) RepayLoan(loanID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	paidAmount := amount

	return s.DB.Transaction(func(tx *gorm.DB) error {

		loanRepo := repositories.LoanRepo{DB: tx}
		payRepo := repositories.LoanPaymentRepo{DB: tx}

		loan, err := loanRepo.GetByID(loanID)
		if err != nil {
			return err
		}

		if loan.Status != models.StatusActive {
			return errors.New("loan inactive")
		}

		if loan.RemainingPrincipal == 0 && loan.InterestDue == 0 {
			return errors.New("loan already closed")
		}

		//
		// ✅ pay interest first
		//
		if loan.InterestDue > 0 {
			if amount >= loan.InterestDue {
				amount -= loan.InterestDue
				loan.InterestDue = 0
			} else {
				loan.InterestDue -= amount
				amount = 0
			}
		}

		if amount > 0 {
			if amount > loan.RemainingPrincipal {
				amount = loan.RemainingPrincipal
			}
			loan.RemainingPrincipal -= amount
		}
		if loan.RemainingPrincipal == 0 && loan.InterestDue == 0 {
			loan.Status = models.StatusInactive
		}

		if err := loanRepo.Update(&loan); err != nil {
			return err
		}
		payment := models.LoanPayment{
			LoanID:      loanID,
			AmountPaid:  paidAmount,
			PaymentDate: time.Now(),
		}

		return payRepo.Create(&payment)
	})
}
