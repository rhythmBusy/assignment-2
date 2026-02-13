package models

type User struct {
	Base

	Name  string `gorm:"size:120;not null"`
	Email string `gorm:"size:150;uniqueIndex;not null"`

	
	Accounts []Account `gorm:"many2many:account_users;"`

	Loans []Loan
}
