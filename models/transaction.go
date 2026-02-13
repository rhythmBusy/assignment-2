package models

type Transaction struct {
	Base

	AccountID uint `gorm:"not null;index"`
	Account   Account

	Amount float64 `gorm:"not null"`
	Type   string  `gorm:"size:20;not null"` // DEPOSIT / WITHDRAW
}
