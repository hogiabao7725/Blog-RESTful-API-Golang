package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type Config struct {
	ServerPort string
	DB         DBConfig
	JWTSecret  string
	LogLevel   string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("no found .env file, default using environment variables")
	}

	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWTSecret: getEnv("JWT_SECRET", ""),
		LogLevel:  getEnv("LOG_LEVEL", "info"),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	log.Printf("configuration loaded successfully: %s\n", cfg.String())

	return cfg, nil
}

func (c *Config) validate() error {
	if c.DB.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if c.DB.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if c.DB.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if c.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	return nil
}

/*
 * DSN returns the Data Source Name for connecting to the database
 * format in Postgres: "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
 */
func (c *Config) DBDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.Name, c.DB.SSLMode,
	)
}

// String in config
func (c *Config) String() string {
	return fmt.Sprintf(
		"ServerPort=%s, DBHost=%s, DBPort=%s, DBUser=%s, DBName=%s, DBSSLMode=%s, LogLevel=%s, JWTSecret=[hidden]",
		c.ServerPort, c.DB.Host, c.DB.Port, c.DB.User, c.DB.Name, c.DB.SSLMode, c.LogLevel,
	)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
