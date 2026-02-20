package domain

type Options struct {
	Id         string `gorm:"primaryKey"`
	Content    string
	QuestionId string `gorm:"not null"`
}
