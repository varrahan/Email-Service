package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"email-service/internal/model"
	"email-service/internal/service"

	"go.uber.org/zap"
)

type EmailHandler struct {
	service *service.EmailService
	logger *zap.Logger
}

func NewEmailHandler(emailService *service.EmailService, logger *zap.Logger) *EmailHandler {
	return &EmailHandler{
		service: emailService,
		logger: logger,
	}
}

func (h *EmailHandler) SendEmail(c *gin.Context) {
	var req model.Email
	h.logger.Debug("handling send email request")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(
			"Error with binding request data to variable", 
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.SendContactEmail(c.Request.Context(), req); err != nil {
		h.logger.Error(
			"Error with sending email", 
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email:",})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}