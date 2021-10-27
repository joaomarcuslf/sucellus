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
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Url:      os.Getenv("DB_URL"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
			Driver:   os.Getenv("DB_DRIVER"),
		},
	}
}
