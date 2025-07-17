package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "3000"),
		DBUser:     getEnv("DB_USER", "eugenio"),
		DBPassword: getEnv("DB_PASSWORD", "aA@12345"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "mysql"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "gontabilizador"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}