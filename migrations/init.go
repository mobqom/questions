package migrations

import (
	"fmt"

	"github.com/mobqom/questions/internal/repository"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	models := []any{&repository.QuestionModel{}, &repository.OptionModel{}}
	err := db.AutoMigrate(models...)
	if err != nil {
		fmt.Printf("error with gorm automigrate: %v \n", err)
		return
	}
}
