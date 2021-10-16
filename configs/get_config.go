package configs

import "os"

type AppConfig struct {
	Port     string
	Database DatabaseConfig
}

func GetConfig() *AppConfig {
	return &AppConfig{
		Port: os.Getenv("PORT"),
		Database: DatabaseConfig{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
			Url:      os.Getenv("MONGODB_URL"),
			Port:     os.Getenv("MONGODB_PORT"),
			Database: os.Getenv("MONGODB_DATABASE"),
		},
	}
}
