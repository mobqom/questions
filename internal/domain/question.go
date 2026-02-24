package domain

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content string    `gorm:"not null"`
	Game    string    `gorm:"type:varchar(255);not null"`
	Options []Options `gorm:"foreignKey:QuestionID"`
}
