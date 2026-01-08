package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"email-service/internal/model"
	"email-service/internal/service"
)

type Handler struct {
	emailService *service.EmailService
}

func NewHandler(emailService *service.EmailService) *Handler {
	return &Handler{emailService: emailService}
}

func (h *Handler) SendEmail(c *gin.Context) {
	var req model.Email
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.emailService.SendContactEmail(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email:",})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
}