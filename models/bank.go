package models

type Bank struct {
	Base

	Name string `gorm:"size:120;not null"`
	Code string `gorm:"size:20;uniqueIndex;not null"`

	Branches []Branch
}
