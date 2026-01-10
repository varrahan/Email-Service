package main

import (
	"email-service/internal/config"
	"email-service/internal/handler"
	"email-service/internal/logger"
	"email-service/internal/sender"
	"email-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	corsGin "github.com/rs/cors/wrapper/gin"
	"go.uber.org/zap"
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

	router.Use(corsGin.New(cors.Options{
		AllowedOrigins:   []string{config.CORSOrigin},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))
	router.LoadHTMLFiles("templates/index.html")

	router.GET("/", homeHander.HandlePage)
	router.POST("/send", emailHandler.SendEmail)

	if err := router.Run(config.AppPort); err != nil {
		logger.Log.Fatal("Email service stopped with error", zap.Error(err))
	}
	logger.Log.Info("Email service stopped gracefully")
}