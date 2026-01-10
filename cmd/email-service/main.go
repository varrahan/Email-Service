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
	
	config, err := config.LoadConfig()
	if err != nil {
		logger.Log.Fatal("Environment variables have not been set up correctly", zap.Error(err))
	}

	smtpSender := sender.NewSmtpSender(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass, config.ToAddr, config.FromAddr, logger.Log.Named("Sender"))
	emailService := service.NewEmailService(smtpSender, logger.Log.Named("Service"))
	emailHandler := handler.NewEmailHandler(emailService, logger.Log.Named("Handler"))
	homeHander := handler.NewHomeHandler(logger.Log.Named("Handler"))

	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")

	router.GET("/", homeHander.HandlePage)
	router.POST("/send", emailHandler.SendEmail)

	if err := router.Run(config.AppPort); err != nil {
		logger.Log.Fatal("Email service stopped with error", zap.Error(err))
	}
	logger.Log.Info("Email service stopped gracefully")
}