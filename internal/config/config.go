package config

import (
	"os"
	"strconv"
    "errors"

    "github.com/joho/godotenv"
)

type Config struct {
    SMTPHost    string
    SMTPPort    int    // Port provided by third-party SMTP service
    SMTPUser    string
    SMTPPass    string
    ToAddr      string // Personal or service email
    FromAddr    string // Email connected to SMTP provider
    AppPort     string // Port running the application
    CORSOrigin  string // A Single origin, most likely the frontend for the service
}

func LoadConfig() (*Config, error) {
    godotenv.Load()

    smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		smtpPort = 587
	}

    appPort := os.Getenv("PORT")
    if appPort == "" {
        appPort = "4000"
    }

    smtpHost := os.Getenv("SMTP_HOST")
    if smtpHost == "" {
        return nil, errors.New("SMTP_HOST environment variable not configured")
    }
    smtpUser := os.Getenv("SMTP_USER")
    if smtpUser == "" {
        return nil, errors.New("SMTP_USER environment variable not configured")
    }
    smtpPass := os.Getenv("SMTP_PASS")
        if smtpPass == "" {
        return nil, errors.New("SMTP_PASS environment variable not configured")
    }
    toAddr := os.Getenv("TO_ADDRESS")
    if toAddr == "" {
        return nil, errors.New("TO_ADDRESS environment variable not configured")
    }
    fromAddr := os.Getenv("FROM_ADDRESS")
    if fromAddr == "" {
        return nil, errors.New("FROM_ADDRESS environment variable not configured")
    }
    corsOrigin := os.Getenv("ALLOWED_ORIGIN")

    return &Config{
		SMTPHost: smtpHost,
        SMTPPort: smtpPort,
        SMTPUser: smtpUser,
        SMTPPass: smtpPass,
        ToAddr: toAddr,
        FromAddr: fromAddr,
        AppPort: ":" + appPort,
        CORSOrigin: corsOrigin,
    }, nil
}

func GetConfig() *Config {

    godotenv.Load()

    portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 587
	}

    return &Config{
		SMTPHost: os.Getenv("SMTP_HOST"),
        SMTPPort: port,
        SMTPUser: os.Getenv("SMTP_USER"),
        SMTPPass: os.Getenv("SMTP_PASS"),
        ToAddr: os.Getenv("TO_ADDRESS"),
        FromAddr: os.Getenv("FROM_ADDRESS"),
    }
}