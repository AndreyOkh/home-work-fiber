package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func Init(confPath ...string) {
	if len(confPath) == 0 {
		confPath = append(confPath, ".env")
	}
	for _, path := range confPath {
		if err := godotenv.Load(path); err != nil {
			log.Printf("Error loading %s file", path)
			continue
		}
		log.Printf("Loaded %s config file", path)
	}
}

type DatabaseConfig struct {
	Host string
	Port int
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host: getStrEnv("DB_HOST", "localhost"),
		Port: getIntEnv("DB_PORT", 5432),
	}
}

func getStrEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		val, err := strconv.Atoi(value)
		if err == nil {
			return val
		}
	}
	return fallback
}

//func getBoolEnv(key string, fallback bool) bool {
//	if value := os.Getenv(key); value != "" {
//		val, err := strconv.ParseBool(value)
//		if err == nil {
//			return val
//		}
//	}
//	return fallback
//}

type LogConfig struct {
	Level  string
	Format string
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  strings.ToLower(getStrEnv("LOG_LEVEL", "info")),
		Format: strings.ToLower(getStrEnv("LOG_FORMAT", "json")),
	}
}
