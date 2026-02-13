package models

import "time"

type Base struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status `gorm:"type:varchar(20);default:'ACTIVE';index"`
}
