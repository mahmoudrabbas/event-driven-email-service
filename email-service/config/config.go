package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
}

var SMTP SMTPConfig

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	SMTP = SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Email:    os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}
}
