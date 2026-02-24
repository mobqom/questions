package domain

import "gorm.io/gorm"

type Options struct {
	gorm.Model
	Content    string   `gorm:"not null"`
	QuestionID uint     `gorm:"not null"`
	Question   Question `gorm:"foreignKey:QuestionID"`
}
