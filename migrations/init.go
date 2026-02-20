package migrations

import (
	"fmt"

	"github.com/mobqom/questions/internal/domain"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	models := []any{&domain.Options{}, &domain.Question{}}
	err := db.AutoMigrate(models...)
	if err != nil {
		fmt.Printf("error with gorm automigrate: %v \n", err)
		return
	}
}
