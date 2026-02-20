package domain

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Id      string `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	Game    string `gorm:"type:varchar(255);not null"`
}
