package main

import (
	"email-service/internal/config"
	"email-service/internal/handler"
	"email-service/internal/logger"
	"email-service/internal/sender"
	"email-service/internal/service"


	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	defer logger.Sync()
	zap.RedirectStdLog(logger.Log)

	cfg := config.GetConfig()

	smtpSender := sender.NewSmtpSender(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass, cfg.ToAddr, cfg.FromAddr, logger.Log.Named("Sender"))
	emailService := service.NewEmailService(smtpSender, logger.Log.Named("Service"))
	emailHandler := handler.NewEmailHandler(emailService, logger.Log.Named("Handler"))

	router := gin.Default()
	router.POST("/send", emailHandler.SendEmail)

	if err := router.Run(":8080"); err != nil {
		logger.Log.Fatal("Email service stopped with error", zap.Error(err))
	}
	logger.Log.Info("Email service stopped gracefully")
}