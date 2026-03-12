package main

// @title Questions API
// @version 1.0
// @description API для работы с вопросами
// @host localhost:8081
// @BasePath /api/v1
// @schemes http
// @accept json
// @produce json

import (
	"fmt"

	"github.com/mobqom/questions/config"
	"github.com/mobqom/questions/internal/server"
)

func main() {
	fmt.Println("hello")
	cfg := config.Init()
	server.Run(cfg)
}
