package models

type Account struct {
	Base

	AccountNumber string  `gorm:"size:30;uniqueIndex;not null"`
	Balance       float64 `gorm:"not null;default:0"`

	
	Users []User `gorm:"many2many:account_users;"`

	BranchID uint   `gorm:"not null"`
	Branch   Branch

	Transactions []Transaction
}

