package main

import (
	"email-service/internal/config"
	"email-service/internal/handler"
	"email-service/internal/sender"
	"email-service/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	smtpSender := sender.NewSmtpSender(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass, cfg.ToAddr, cfg.FromAddr)
	emailService := service.NewEmailService(smtpSender)
	emailHandler := handler.New(emailService)

	router := gin.Default()
	router.POST("/send", emailHandler.SendEmail)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}