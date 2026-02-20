package config

type AppConfig struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

func Init() *AppConfig {
	return &AppConfig{
		DbHost: "localhost",
		DbPort: "5432",
		DbUser: "user",
		DbPass: "password",
		DbName: "questions",
	}
}
