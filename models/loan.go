package models

type Loan struct {
	Base

	UserID uint `gorm:"not null"`
	User   User

	BranchID uint `gorm:"not null"`
	Branch   Branch

	Principal          float64 `gorm:"not null"`
	InterestRate       float64 `gorm:"not null;default:12"`
	RemainingPrincipal float64 `gorm:"not null"`
	InterestDue        float64 `gorm:"not null"`
	InterestThisYear   float64 `gorm:"not null"`

	Payments []LoanPayment
}
