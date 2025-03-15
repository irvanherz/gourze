package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Bunny    BunnyConfig
	Auth     AuthConfig
}

type DatabaseConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

type BunnyConfig struct {
	StorageZone     string
	AccessKey       string
	Region          string
	DownloadBaseURL string
}

type AuthConfig struct {
	JWTSecret string
}

func ProvideConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		panic("⚠️ No .env file found!")
	}
	return &Config{
		Database: DatabaseConfig{
			Host: getEnv("DB_HOST", ""),
			Port: getEnv("DB_PORT", ""),
			User: getEnv("DB_USER", ""),
			Pass: getEnv("DB_PASS", ""),
			Name: getEnv("DB_NAME", ""),
		},
		Bunny: BunnyConfig{
			StorageZone:     getEnv("BUNNY_STORAGE_ZONE", ""),
			AccessKey:       getEnv("BUNNY_ACCESS_KEY", ""),
			Region:          getEnv("BUNNY_REGION", ""),
			DownloadBaseURL: getEnv("BUNNY_DOWNLOAD_BASE_URL", ""),
		},
		Auth: AuthConfig{
			JWTSecret: getEnv("JWT_SECRET", ""),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
