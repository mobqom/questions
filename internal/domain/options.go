package domain

import "gorm.io/gorm"

type Options struct {
	gorm.Model
	Id         string `gorm:"primaryKey"`
	Content    string `gorm:"not null"`
	QuestionId string `gorm:"not null"`
}
