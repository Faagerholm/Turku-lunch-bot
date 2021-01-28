package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey string
}

func Get(files ...string) (conf *Config, err error) {
	if err := godotenv.Load(files...); err != nil {
		return nil, err
	}
	conf = &Config{
		ApiKey: getEnv("TELEGRAM_API_KEY", ""),
	}
	return
}

func getEnv(key string, defValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defValue
}
