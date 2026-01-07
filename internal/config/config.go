package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
    SMTPHost     string
    SMTPPort     int
    SMTPUser     string
    SMTPPass     string
    FromAddr	 string
}

func GetConfig() *Config {
    portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error converting port:", err)
		port = 8000
	}

    return &Config{
		SMTPHost: os.Getenv("SMTP_HOST"),
        SMTPPort: port,
        SMTPUser: os.Getenv("SMTP_USER"),
        SMTPPass: os.Getenv("SMTP_PASS"),
        FromAddr: os.Getenv("FROM_ADDRESS"),
    }
}