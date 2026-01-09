package config

import (
	"os"
	"strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    SMTPHost    string
    SMTPPort    int    // Port provided by third-party SMTP service
    SMTPUser    string
    SMTPPass    string
    ToAddr      string // Personal or service email
    FromAddr    string // Email connected to SMTP provider
    AppPort     string    // Port running the application
}

func LoadConfig() *Config {
    godotenv.Load()

    smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		smtpPort = 587
	}

    appPort := os.Getenv("APP_PORT")
    if appPort == "" {
        appPort = "4000"
    }

    return &Config{
		SMTPHost: os.Getenv("SMTP_HOST"),
        SMTPPort: smtpPort,
        SMTPUser: os.Getenv("SMTP_USER"),
        SMTPPass: os.Getenv("SMTP_PASS"),
        ToAddr: os.Getenv("TO_ADDRESS"),
        FromAddr: os.Getenv("FROM_ADDRESS"),
        AppPort: ":" + appPort,
    } 
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