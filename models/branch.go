package models

type Branch struct {
	Base

	Name    string `gorm:"size:120;not null"`
	IFSC    string `gorm:"size:20;uniqueIndex"`
	Address string `gorm:"size:200"`

	BankID uint `gorm:"not null"`
	Bank   Bank

	Accounts []Account
	Loans    []Loan
}
