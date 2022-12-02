// Package config handles environment variables.
package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config contains environment variables.
type Config struct {
	Port                       string  `envconfig:"PORT" default:"8000"`
	SendGridAPIEndpoint        string  `envconfig:"SEND_GRID_API_ENDPOINT" required:"true"`
	SendGridAPIKey             string  `envconfig:"SEND_GRID_API_KEY" required:"true"`
}

// LoadConfig reads environment variables, populates and returns Config.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	var c Config

	err := envconfig.Process("", &c)

	return &c, err
}