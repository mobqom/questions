package usecase

import (
	"fmt"

	"github.com/mobqom/questions/internal/domain"
	"github.com/mobqom/questions/internal/dto"
	"gorm.io/gorm"
)

func GetQuestions(db *gorm.DB) []domain.Question {
	var questions []domain.Question
	db.Model(&domain.Question{}).Find(&questions)

	return questions
}

func AddQuestion(db *gorm.DB, question dto.AddQuestionDto) domain.Question {
	//ctx := context.Background()
	fmt.Println(question)
	q := domain.Question{Content: question.Content, Game: question.Game}
	err := db.Create(&q)
	if err != nil {
		fmt.Println(err)
		return domain.Question{}
	}
	return q
}
