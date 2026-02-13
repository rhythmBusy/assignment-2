package models

import "time"

type LoanPayment struct {
	Base

	LoanID uint `gorm:"not null;index"`
	Loan   Loan

	AmountPaid  float64   `gorm:"not null"`
	PaymentDate time.Time `gorm:"not null"`
}
