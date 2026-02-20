package db

import (
	"fmt"

	"github.com/mobqom/questions/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(cfg *config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: fmt.Sprintf(
					"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
					cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort,
				),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			},
		),
		&gorm.Config{},
	)
	return db, err
}
