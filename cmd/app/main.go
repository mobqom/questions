package main

// @title Questions API
// @version 1.0
// @description API для работы с вопросами
// @BasePath /api/questions
// @accept json
// @produce json

import (
	"github.com/mobqom/questions/config"
	"github.com/mobqom/questions/internal/server"
)

func main() {
	cfg := config.Init()
	server.Run(cfg)
}
