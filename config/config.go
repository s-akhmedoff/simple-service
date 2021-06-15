package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"os"
	"time"
)

const ServiceName = "service:simple_service"

// Config ...
type Config struct {
	Environment      string
	HTTPHost         string
	HTTPPort         string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

// Load ...
func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Environment:      cast.ToString(env("ENVIRONMENT", "dev")),
		HTTPHost:         cast.ToString(env("HTTP_HOST", "localhost")),
		HTTPPort:         cast.ToString(env("HTTP_PORT", ":7070")),
		PostgresHost:     cast.ToString(env("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToString(env("POSTGRES_PORT", "5432")),
		PostgresUser:     cast.ToString(env("POSTGRES_USER", "admin")),
		PostgresPassword: cast.ToString(env("POSTGRES_PASSWORD", "root")),
		PostgresDatabase: cast.ToString(env("POSTGRES_DATABASE", "postgres")),
	}
}
func (c *Config) PostgresURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDatabase)
}

func (c *Config) LoggerSetup() *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		ForceQuote:                true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           time.RFC850,
		PadLevelText:              true,
		QuoteEmptyFields:          true,
	})

	return log
}

func env(key string, defaultValue interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
